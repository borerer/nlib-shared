package nlibshared

type PayloadCallFunctionResponse struct {
	FuncName string      `json:"func_name" mapstructure:"func_name"`
	Response interface{} `json:"response" mapstructure:"response"`
}
