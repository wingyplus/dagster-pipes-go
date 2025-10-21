package dagster_pipes

import (
	"encoding/json"
	"os"
)

type LoadContext interface {
	LoadContext(params map[string]json.RawMessage) (*PipesContextData, error)
}

type PayloadErrorKind string

func (e PayloadErrorKind) Error() string {
	return string(e)
}

var (
	PayloadErrorKind_Missing PayloadErrorKind = "no payload found in params"
)

type DefaultLoader struct{}

func NewDefaultLoader() *DefaultLoader {
	return &DefaultLoader{}
}

func (loader *DefaultLoader) LoadContext(params map[string]json.RawMessage) (*PipesContextData, error) {
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

		var contextData PipesContextData
		if err := json.NewDecoder(f).Decode(&contextData); err != nil {
			// TODO: convert to PayloadErrorKind
			return nil, err
		}
		return &contextData, nil
	}
	if data, ok := params["data"]; ok {
		var contextData PipesContextData
		if err := json.Unmarshal(data, &contextData); err != nil {
			// TODO: convert to PayloadErrorKind
			return nil, err
		}
		return &contextData, nil
	}
	return nil, PayloadErrorKind_Missing
}
