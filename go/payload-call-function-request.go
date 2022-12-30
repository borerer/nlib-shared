package nlibshared

type PayloadCallFunctionRequest struct {
	FuncName string                 `json:"func_name" mapstructure:"func_name"`
	Params   map[string]interface{} `json:"params" mapstructure:"params"`
}
