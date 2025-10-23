package dagster_pipes

import (
	"encoding/json"
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wingyplus/dagster-pipes-go/metadata"
	"github.com/wingyplus/dagster-pipes-go/types"
)

func TestWritePipesMetadata(t *testing.T) {
	t.Parallel()
	file, context := singleAssetFileAndContext(t)

	assetMetadata := map[string]*types.PipesMetadataValue{
		"int":       metadata.FromInt(100),
		"float":     metadata.FromFloat(100.0),
		"bool":      metadata.FromBool(true),
		"none":      metadata.Null(),
		"timestamp": metadata.FromTimestamp(1000.0),
		"text":      metadata.FromText("hello"),
		"url":       metadata.FromURLString("http://someurl.com"),
		"path":      metadata.FromPath("file://some/path"),
		"notebook":  metadata.FromNotebook("notebook"),
		"json_object": metadata.FromJSON(map[string]any{
			"key": "value",
		}),
		"json_array": metadata.FromJSONArray([]any{
			map[string]any{"key": "value"},
		}),
		"md":          metadata.FromMd("## markdown"),
		"dagster_run": metadata.FromDagsterRun("1234"),
		"asset":       metadata.FromAsset("some_asset"),
		"job":         metadata.FromJob("some_job"),
	}

	err := context.ReportAssetMaterialization("asset1", assetMetadata, "v1")
	require.NoError(t, err)

	content, err := os.ReadFile(file.Path)
	require.NoError(t, err)

	var message types.PipesMessage
	err = json.Unmarshal(content, &message)
	require.NoError(t, err)

	require.Equal(t, &types.PipesMessage{
		Method: types.ReportAssetMaterialization,
		Params: map[string]any{
			"asset_key": "asset1",
			"metadata": map[string]any{
				"int":         map[string]any{"raw_value": float64(100), "type": "int"},
				"float":       map[string]any{"raw_value": float64(100), "type": "float"},
				"bool":        map[string]any{"raw_value": true, "type": "bool"},
				"none":        map[string]any{"raw_value": nil, "type": "null"},
				"timestamp":   map[string]any{"raw_value": float64(1000), "type": "timestamp"},
				"text":        map[string]any{"raw_value": "hello", "type": "text"},
				"url":         map[string]any{"raw_value": "http://someurl.com", "type": "url"},
				"path":        map[string]any{"raw_value": "file://some/path", "type": "path"},
				"notebook":    map[string]any{"raw_value": "notebook", "type": "notebook"},
				"json_object": map[string]any{"raw_value": map[string]any{"key": "value"}, "type": "json"},
				"json_array":  map[string]any{"raw_value": []any{map[string]any{"key": "value"}}, "type": "json"},
				"md":          map[string]any{"raw_value": "## markdown", "type": "md"},
				"dagster_run": map[string]any{"raw_value": "1234", "type": "dagster_run"},
				"asset":       map[string]any{"raw_value": "some_asset", "type": "asset"},
				"job":         map[string]any{"raw_value": "some_job", "type": "job"},
			},
			"data_version": "v1",
		},
	}, &message)
}

func TestClosePipesContext(t *testing.T) {
	t.Parallel()
	file, context := singleAssetFileAndContext(t)

	err := context.Close(PipesExceptionError(errors.New("some error")))
	require.NoError(t, err)

	content, err := os.ReadFile(file.Path)
	require.NoError(t, err)

	var message types.PipesMessage
	err = json.Unmarshal(content, &message)
	require.NoError(t, err)

	require.Equal(t, &types.PipesMessage{
		DagsterPipesVersion: "0.1",
		Method:              types.Closed,
		Params: map[string]any{
			"cause":   nil,
			"context": nil,
			"message": "some error",
			"name":    "some error",
			"stack":   nil,
		},
	}, &message)
}

func singleAssetFileAndContext(t *testing.T) (*FileChannel, *PipesContext) {
	t.Helper()
	return fileAndContext(t, []string{"asset1"})
}

func multiAssetFileAndContext(t *testing.T) (*FileChannel, *PipesContext) {
	t.Helper()
	return fileAndContext(t, []string{"asset1", "asset2"})
}

func fileAndContext(t *testing.T, assets []string) (*FileChannel, *PipesContext) {
	t.Helper()

	f, err := os.CreateTemp("", "")
	require.NoError(t, err)
	defer f.Close()

	channel := &FileChannel{Path: f.Name()}
	context := &PipesContext{
		Channel: channel,
		Data: &types.PipesContextData{
			AssetKeys: assets,
			RunID:     "012345",
		},
	}
	return channel, context
}
