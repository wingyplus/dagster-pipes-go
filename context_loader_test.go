package dagster_pipes

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDefaultLoader_LoadContext(t *testing.T) {
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
	expected := &PipesContextData{
		AssetKeys: []string{"asset1", "asset2"},
		Extras: map[string]interface{}{
			"key": "value",
		},
		RetryNumber: 0,
		RunID:       "012345",
	}
	loader := NewDefaultLoader()

	t.Run("from file", func(t *testing.T) {
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
		param := map[string]json.RawMessage{
			"data": json.RawMessage([]byte(params)),
		}
		contextData, err := loader.LoadContext(param)
		require.NoError(t, err)
		require.Equal(t, expected, contextData)
	})
}
