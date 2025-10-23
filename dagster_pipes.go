package dagster_pipes

import (
	"encoding/json"
)

type PipesContext struct {
	Data    *PipesContextData
	Channel MessageWriterChannel
}

func (context *PipesContext) Close(exception *PipesException) error {
	if exception != nil {
		// TODO: convert to params
	}
	panic("implements me")
}

func (context *PipesContext) ReportAssetMaterialization(
	assetKey string,
	metadata map[string]*PipesMetadataValue,
	dataVersion string,
) error {
	panic("implements me")
}

func (context *PipesContext) ReportAssetCheck(
	checkName string,
	passed bool,
	assetKey string,
	severity *AssetCheckSeverity,
	metadata map[string]*PipesMetadataValue,
) error {
	panic("implements me")
}

func (context *PipesContext) ReportCustomMessage(payload any) error {
	panic("implements me")
}

func NewPipesContext(contextData *PipesContextData, messageParams map[string]json.RawMessage, messageWriter MessageWriter) (*PipesContext, error) {
	channel := messageWriter.Open(messageParams)
	// TODO: initialize PipesLogger

	openedPayload := messageWriter.GetOpenedPayload()
	openedMessage := &PipesMessage{Method: Opened, Params: openedPayload}

	if err := channel.Write(openedMessage); err != nil {
		// TODO: wrap error
		return nil, err
	}

	return &PipesContext{
		Data:    contextData,
		Channel: channel,
	}, nil
}
