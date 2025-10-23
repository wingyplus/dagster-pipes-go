package types

func NewMessage(method Method, params map[string]any) *PipesMessage {
	return &PipesMessage{
		DagsterPipesVersion: "0.1",
		Method:              method,
		Params:              params,
	}
}
