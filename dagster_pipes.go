package dagster_pipes

import (
	"encoding/json"
	"errors"
	"slices"

	"github.com/wingyplus/dagster-pipes-go/types"
)

var ErrMissingAssetKey = errors.New("asset key is missing")

type Metadata map[string]*types.PipesMetadataValue

// PipesContext represents the connection to Dagster and provides methods
// to interact with the Dagster orchestration process.
//
// The context contains the serializable data passed from the orchestration
// process to the external process and a message channel for communication.
type PipesContext struct {
	// Data contains context information from Dagster including asset keys,
	// run ID, partition information, and other metadata.
	Data *types.PipesContextData
	// Channel is the communication channel for sending messages back to Dagster.
	Channel MessageWriterChannel
}

// Close sends a close message to Dagster and terminates the pipes connection.
//
// This should be called when your application is done communicating with Dagster.
// It's recommended to use defer to ensure Close is called even if an error occurs:
//
//	context, err := dagster_pipes.OpenDasterPipes()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer context.Close(nil)
//
// If your application encountered an exception that you want to report to Dagster,
// you can pass a PipesException:
//
//	defer func() {
//	    if r := recover(); r != nil {
//	        exc := &types.PipesException{
//	            Message: helper.Ptr(fmt.Sprintf("%v", r)),
//	            Name:    helper.Ptr("panic"),
//	        }
//	        context.Close(exc)
//	    }
//	}()
func (context *PipesContext) Close(exception *types.PipesException) error {
	var params map[string]any = nil
	if exception != nil {
		params = map[string]any{
			"cause":   exception.Cause,
			"context": exception.Context,
			"message": exception.Message,
			"name":    exception.Name,
			"stack":   exception.Stack,
		}
	}
	closedMessage := types.NewMessage(types.Closed, params)
	return context.Channel.Write(closedMessage)
}

// ReportAssetMaterialization reports an asset materialization to Dagster.
//
// This method informs Dagster that an asset has been produced or updated.
// You can include arbitrary metadata and a data version.
//
// Parameters:
//   - assetKey: The key identifying the asset that was materialized
//   - metadata: A map of metadata key-value pairs providing additional information
//     about the materialization. Use the metadata package helpers to create values.
//   - dataVersion: A version string to track changes to the asset data
//
// Example:
//
//	err := context.ReportAssetMaterialization(
//	    "my_dataset",
//	    map[string]*types.PipesMetadataValue{
//	        "row_count": metadata.FromInt(1000),
//	        "schema_version": metadata.FromText("v2.1"),
//	        "output_path": metadata.FromPath("/data/output.parquet"),
//	    },
//	    "2024-01-15",
//	)
func (context *PipesContext) ReportAssetMaterialization(
	assetKey string,
	metadata Metadata,
	dataVersion string,
) error {
	ak, err := context.resolveOptionallyPassedAssetKey(assetKey)
	if err != nil {
		return err
	}

	var params = map[string]any{
		"asset_key":    ak,
		"metadata":     metadata,
		"data_version": stringOrNil(dataVersion),
	}
	return context.Channel.Write(&types.PipesMessage{
		Method: types.ReportAssetMaterialization,
		Params: params,
	})
}

func (context *PipesContext) resolveOptionallyPassedAssetKey(assetKey string) (string, error) {
	if assetKeys := context.Data.AssetKeys; len(assetKeys) > 0 {
		if slices.Contains(context.Data.AssetKeys, assetKey) {
			return resolveNoEmpty(assetKey)
		}
		return assetKeys[0], nil
	}
	return resolveNoEmpty(assetKey)
}

// ReportAssetCheck reports the result of an asset check to Dagster.
//
// Asset checks are data quality tests that validate assumptions about your assets.
// Use this method to report whether a check passed or failed.
//
// Parameters:
//   - checkName: The name of the check being reported
//   - passed: Whether the check passed (true) or failed (false)
//   - assetKey: The key identifying the asset being checked
//   - severity: The severity level (types.Warn or types.AssetCheckSeverityERROR)
//   - metadata: Additional metadata about the check result
//
// Example:
//
//	severity := types.AssetCheckSeverityERROR
//	err := context.ReportAssetCheck(
//	    "row_count_check",
//	    true, // check passed
//	    "my_dataset",
//	    &severity,
//	    map[string]*types.PipesMetadataValue{
//	        "actual_count": metadata.FromInt(1000),
//	        "expected_min": metadata.FromInt(100),
//	    },
//	)
func (context *PipesContext) ReportAssetCheck(
	checkName string,
	passed bool,
	assetKey string,
	severity *types.AssetCheckSeverity,
	metadata Metadata,
) error {
	ak, err := context.resolveOptionallyPassedAssetKey(assetKey)
	if err != nil {
		return err
	}
	var params = map[string]any{
		"asset_key":  ak,
		"check_name": checkName,
		"passed":     passed,
		"severity":   severity,
		"metadata":   metadata,
	}
	return context.Channel.Write(types.NewMessage(types.ReportAssetCheck, params))
}

// ReportCustomMessage sends a custom message payload to Dagster.
//
// This method allows you to send arbitrary structured data to Dagster
// that doesn't fit into the standard asset materialization or check patterns.
// The payload can be any JSON-serializable value.
//
// Example:
//
//	err := context.ReportCustomMessage(map[string]any{
//	    "event_type": "processing_complete",
//	    "duration_ms": 1234,
//	    "items_processed": 5000,
//	})
func (context *PipesContext) ReportCustomMessage(payload any) error {
	var params = map[string]any{
		"payload": payload,
	}
	return context.Channel.Write(types.NewMessage(types.ReportCustomMessage, params))
}

// NewPipesContext creates a new PipesContext with the given parameters.
//
// This is a lower-level constructor that allows you to provide custom
// context data, message parameters, and message writer. For most use cases,
// use OpenDasterPipes instead, which handles all the initialization automatically.
//
// Parameters:
//   - contextData: The context data received from Dagster
//   - messageParams: Parameters for configuring the message channel
//   - messageWriter: The message writer implementation to use
//
// The function opens a message channel, sends an "opened" message to Dagster
// to signal that the pipes connection is established, and returns the context.
func NewPipesContext(contextData *types.PipesContextData, messageParams map[string]json.RawMessage, messageWriter MessageWriter) (*PipesContext, error) {
	channel := messageWriter.Open(messageParams)
	// TODO: initialize PipesLogger

	openedPayload := messageWriter.GetOpenedPayload()
	openedMessage := &types.PipesMessage{Method: types.Opened, Params: openedPayload}

	if err := channel.Write(openedMessage); err != nil {
		// TODO: wrap error
		return nil, err
	}

	return &PipesContext{
		Data:    contextData,
		Channel: channel,
	}, nil
}

// OpenDasterPipes opens a connection to Dagster and returns a PipesContext.
//
// This is the main entry point for using Dagster Pipes in your Go application.
// It automatically loads configuration from environment variables, establishes
// a communication channel with Dagster, and returns a context that you can use
// to report materializations, checks, and custom messages.
//
// The function uses default implementations for:
//   - Parameter loading (from environment variables)
//   - Context loading (deserializing Dagster context data)
//   - Message writing (file-based or stdout-based communication)
//
// Example:
//
//	context, err := dagster_pipes.OpenDasterPipes()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer context.Close(nil)
//
//	// Use the context to interact with Dagster
//	err = context.ReportAssetMaterialization(...)
//
// Returns an error if the environment is not properly configured for Dagster Pipes
// or if communication with Dagster cannot be established.
func OpenDasterPipes() (*PipesContext, error) {
	paramsLoader := NewEnvVarLoader()
	contextLoader := NewDefaultContextLoader()
	messageWriter := NewDefaultMessageWriter()

	contextParams, err := paramsLoader.LoadContextParams()
	if err != nil {
		return nil, err
	}

	messageParams, err := paramsLoader.LoadMessageParams()
	if err != nil {
		return nil, err
	}

	contextData, err := contextLoader.LoadContext(contextParams)
	if err != nil {
		return nil, err
	}

	return NewPipesContext(contextData, messageParams, messageWriter)
}

func resolveNoEmpty(assetKey string) (string, error) {
	if len(assetKey) == 0 {
		return "", ErrMissingAssetKey
	}
	return assetKey, nil
}

func stringOrNil(s string) *string {
	if len(s) == 0 {
		return nil
	}
	return &s
}
