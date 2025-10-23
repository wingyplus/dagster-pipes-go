package dagster_pipes

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wingyplus/dagster-pipes-go/types"
)

func TestDefaultLoader_LoadContext(t *testing.T) {
	t.Parallel()
	params := `
{
	"asset_keys": [
		"asset1",
		"asset2"
	],
	"extras": {
		"key": "value"
	},
	"retry_number": 0,
	"run_id": "012345"
}
`
	expected := &types.PipesContextData{
		AssetKeys: []string{"asset1", "asset2"},
		Extras: map[string]interface{}{
			"key": "value",
		},
		RetryNumber: 0,
		RunID:       "012345",
	}
	loader := NewDefaultContextLoader()

	t.Run("from file", func(t *testing.T) {
		t.Parallel()
		f, err := os.CreateTemp("", "*")
		require.NoError(t, err)
		defer f.Close()
		f.Write([]byte(params))

		param := map[string]json.RawMessage{
			"path": json.RawMessage([]byte(`"` + f.Name() + `"`)),
		}
		contextData, err := loader.LoadContext(param)
		require.Equal(t, expected, contextData)
	})

	t.Run("from data", func(t *testing.T) {
		t.Parallel()
		param := map[string]json.RawMessage{
			"data": json.RawMessage([]byte(params)),
		}
		contextData, err := loader.LoadContext(param)
		require.NoError(t, err)
		require.Equal(t, expected, contextData)
	})

	t.Run("prioritizes file path when both are present", func(t *testing.T) {
		t.Parallel()
		// Create temp file with different data
		f, err := os.CreateTemp("", "*")
		require.NoError(t, err)
		defer f.Close()

		fileData := `
{
	"asset_keys": ["asset_from_path"],
	"extras": {"key_from_path": "value_from_path"},
	"retry_number": 0,
	"run_id": "id_from_path"
}`
		f.Write([]byte(fileData))

		// Provide both path and data
		param := map[string]json.RawMessage{
			"path": json.RawMessage([]byte(`"` + f.Name() + `"`)),
			"data": json.RawMessage([]byte(`{
				"asset_keys": ["asset_from_data"],
				"extras": {"key_from_data": "value_from_data"},
				"retry_number": 0,
				"run_id": "id_from_data"
			}`)),
		}

		contextData, err := loader.LoadContext(param)
		require.NoError(t, err)

		// Should use data from file, not from data param
		expectedFromPath := &types.PipesContextData{
			AssetKeys: []string{"asset_from_path"},
			Extras: map[string]interface{}{
				"key_from_path": "value_from_path",
			},
			RetryNumber: 0,
			RunID:       "id_from_path",
		}
		require.Equal(t, expectedFromPath, contextData)
	})
}
