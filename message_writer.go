package dagster_pipes

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/wingyplus/dagster-pipes-go/types"
)

type MessageWriterChannel interface {
	Write(*types.PipesMessage) error
}

type FileChannel struct {
	Path string
}

func (channel *FileChannel) Write(message *types.PipesMessage) error {
	f, err := os.OpenFile(channel.Path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := json.NewEncoder(f).Encode(message); err != nil {
		// TODO: map error
		return err
	}
	return nil
}

type MessageWriter interface {
	Open(params map[string]json.RawMessage) MessageWriterChannel
	GetOpenedPayload() map[string]any
	GetOpenedExtras() map[string]any
}

type DefaultMessageWriter struct {
}

func NewDefaultMessageWriter() *DefaultMessageWriter {
	return &DefaultMessageWriter{}
}

func (writer *DefaultMessageWriter) Open(params map[string]json.RawMessage) MessageWriterChannel {
	var (
		filePathKey = "path"
		// stdioKey         = "stdio"
		// bufferedStdioKey = "buffered_stdio"
		// stderr = "stderr"
		// stdout = "stdout"
	)

	if value, ok := params[filePathKey]; ok {
		var path string
		if err := json.Unmarshal(value, &path); err != nil {
			panic(fmt.Errorf("cannot unmarshal path: %w", err))
		}
		return &FileChannel{Path: path}
	}

	// if stream, ok := params[stdioKey]; ok {
	// 	// TODO: Stream channel
	// 	panic("implements me")
	// }
	//
	// if stream, ok := params[bufferedStdioKey]; ok {
	// 	// TODO: BufferedStream channel
	// 	panic("implements me")
	// }

	return nil
}

func (writer *DefaultMessageWriter) GetOpenedPayload() map[string]any {
	return map[string]any{
		"extras": writer.GetOpenedExtras(),
	}
}

func (writer *DefaultMessageWriter) GetOpenedExtras() map[string]any {
	return make(map[string]any)
}
