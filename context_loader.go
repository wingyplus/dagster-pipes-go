package dagster_pipes

import (
	"encoding/json"
	"os"

	"github.com/wingyplus/dagster-pipes-go/types"
)

type LoadContext interface {
	LoadContext(params map[string]json.RawMessage) (*types.PipesContextData, error)
}

type PayloadErrorKind string

func (e PayloadErrorKind) Error() string {
	return string(e)
}

var (
	PayloadErrorKind_Missing PayloadErrorKind = "no payload found in params"
)

type DefaultContextLoader struct{}

func NewDefaultContextLoader() *DefaultContextLoader {
	return &DefaultContextLoader{}
}

func (loader *DefaultContextLoader) LoadContext(params map[string]json.RawMessage) (*types.PipesContextData, error) {
	if pathJsonString, ok := params["path"]; ok {
		var path string
		if err := json.Unmarshal(pathJsonString, &path); err != nil {
			// TODO: convert to PayloadErrorKind
			return nil, err
		}

		f, err := os.Open(path)
		if err != nil {
			// TODO: convert to PayloadErrorKind
			return nil, err
		}
		defer f.Close()

		var contextData types.PipesContextData
		if err := json.NewDecoder(f).Decode(&contextData); err != nil {
			// TODO: convert to PayloadErrorKind
			return nil, err
		}
		return &contextData, nil
	}
	if data, ok := params["data"]; ok {
		var contextData types.PipesContextData
		if err := json.Unmarshal(data, &contextData); err != nil {
			// TODO: convert to PayloadErrorKind
			return nil, err
		}
		return &contextData, nil
	}
	return nil, PayloadErrorKind_Missing
}
