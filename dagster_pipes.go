package dagster_pipes

import (
	"encoding/json"

	"github.com/wingyplus/dagster_pipes_go/types"
)

type PipesContext struct {
	Data    *types.PipesContextData
	Channel MessageWriterChannel
}

func (context *PipesContext) Close(exception *types.PipesException) error {
	if exception != nil {
		// TODO: convert to params
	}
	panic("implements me")
}

func (context *PipesContext) ReportAssetMaterialization(
	assetKey string,
	metadata map[string]*types.PipesMetadataValue,
	dataVersion string,
) error {
	panic("implements me")
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
