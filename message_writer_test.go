package dagster_pipes

import (
	"encoding/json"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wingyplus/dagster-pipes-go/types"
)

func TestFileChannel(t *testing.T) {
	t.Parallel()
	f, err := os.CreateTemp("", "file_channel*")
	require.NoError(t, err)
	defer f.Close()

	channel := &FileChannel{Path: f.Name()}
	err = channel.Write(&types.PipesMessage{Method: types.Opened, Params: nil})
	require.NoError(t, err)

	content, err := io.ReadAll(f)
	require.NoError(t, err)

	var message types.PipesMessage
	err = json.Unmarshal(content, &message)
	require.NoError(t, err)
	require.Equal(t, types.Opened, message.Method)
	require.Nil(t, message.Params)
}

func TestDefaultMessageWriter(t *testing.T) {
	t.Parallel()
	t.Run("open with file path key", func(t *testing.T) {
		t.Parallel()
		writer := NewDefaultMessageWriter()
		channel := writer.Open(map[string]json.RawMessage{
			"path": json.RawMessage([]byte(`"tmp/my-file-path"`)),
		})
		require.Equal(t, &FileChannel{Path: "tmp/my-file-path"}, channel)
	})
}
