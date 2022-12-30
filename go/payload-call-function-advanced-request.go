package nlibshared

type PayloadCallFunctionAdvancedRequest struct {
	FuncName string  `json:"func_name" mapstructure:"func_name"`
	Request  Request `json:"request" mapstructure:"request"`
}
