package dagster_pipes

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
)

type LoadParams interface {
	IsDagsterPipesProcess() bool
	LoadContextParams() (map[string]json.RawMessage, error)
	LoadMessageParams() (map[string]json.RawMessage, error)
}

type ParamsError struct {
	Param  string
	Origin ParamOrigin
	Source ParamsErrorKind
}

func (e *ParamsError) Error() string {
	return fmt.Sprintf("param: %s, origin: %s, source: %s", e.Param, e.Origin, e.Source)
}

type ParamOrigin string

var (
	ParamOrigin_CLI    ParamOrigin = "cli"
	ParamOrigin_EnvVar ParamOrigin = "env var"
)

type ParamsErrorKind string

var (
	ParamsErrorKind_NotPresent ParamsErrorKind = "not present"
	ParamsErrorKind_Invalid    ParamsErrorKind = "invalid"
)

var (
	DAGSTER_PIPES_CONTEXT_ENV_VAR  = "DAGSTER_PIPES_CONTEXT"
	DAGSTER_PIPES_MESSAGES_ENV_VAR = "DAGSTER_PIPES_MESSAGES"
)

type EnvVarLoader struct {
}

func NewEnvVarLoader() *EnvVarLoader {
	return &EnvVarLoader{}
}

func (loader *EnvVarLoader) IsDagsterPipesProcess() bool {
	_, ok := os.LookupEnv(DAGSTER_PIPES_CONTEXT_ENV_VAR)
	return ok
}

func (loader *EnvVarLoader) LoadContextParams() (map[string]json.RawMessage, error) {
	param, ok := os.LookupEnv(DAGSTER_PIPES_CONTEXT_ENV_VAR)
	if !ok {
		return nil, &ParamsError{
			Source: ParamsErrorKind_NotPresent,
			Param:  DAGSTER_PIPES_CONTEXT_ENV_VAR,
			Origin: ParamOrigin_CLI,
		}
	}
	result, err := DecodeEnvVar(param)
	if err != nil {
		// TODO: convert error to ParamsError.
		return nil, err
	}
	return result, nil
}

func (loader *EnvVarLoader) LoadMessageParams() (map[string]json.RawMessage, error) {
	param, ok := os.LookupEnv(DAGSTER_PIPES_MESSAGES_ENV_VAR)
	if !ok {
		return nil, &ParamsError{
			Source: ParamsErrorKind_NotPresent,
			Param:  DAGSTER_PIPES_MESSAGES_ENV_VAR,
			Origin: ParamOrigin_CLI,
		}
	}
	result, err := DecodeEnvVar(param)
	if err != nil {
		// TODO: convert error to ParamsError.
		return nil, err
	}
	return result, nil
}

func DecodeEnvVar(param string) (map[string]json.RawMessage, error) {
	zlibCompressedBytes, err := base64.StdEncoding.DecodeString(param)
	if err != nil {
		return nil, err
	}

	r, err := zlib.NewReader(bytes.NewBuffer(zlibCompressedBytes))
	if err != nil {
		return nil, err
	}

	var result map[string]json.RawMessage
	if err := json.NewDecoder(r).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}
