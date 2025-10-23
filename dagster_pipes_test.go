package dagster_pipes

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wingyplus/dagster_pipes_go/types"
)

func TestWritePipesMetadata(t *testing.T) {
	// file, context := singleAssetFileAndContext(t)
	//
	// assetMetadata := map[string]PipesMetadataValue{
	// 	"int": PipesMetadataValue{Int: 100},
	// 	"int": PipesMetadataValue{Int: 100},
	// }
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
