package dagster_pipes

import (
	"encoding/json"

	"github.com/wingyplus/dagster-pipes-go/types"
)

type PipesContext struct {
	Data    *types.PipesContextData
	Channel MessageWriterChannel
}

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

func (context *PipesContext) ReportAssetMaterialization(
	assetKey string,
	metadata map[string]*types.PipesMetadataValue,
	dataVersion string,
) error {
	metadataMap := make(map[string]any)
	for k, v := range metadata {
		metadataMap[k] = v
	}

	var params = map[string]any{
		"asset_key":    assetKey,
		"metadata":     metadataMap,
		"data_version": dataVersion,
	}
	return context.Channel.Write(&types.PipesMessage{
		Method: types.ReportAssetMaterialization,
		Params: params,
	})
}

func (context *PipesContext) ReportAssetCheck(
	checkName string,
	passed bool,
	assetKey string,
	severity *types.AssetCheckSeverity,
	metadata map[string]*types.PipesMetadataValue,
) error {
	panic("implements me")
}

func (context *PipesContext) ReportCustomMessage(payload any) error {
	panic("implements me")
}

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
