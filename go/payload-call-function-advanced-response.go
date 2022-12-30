package nlibshared

type PayloadCallFunctionAdvancedResponse struct {
	FuncName string   `json:"func_name" mapstructure:"func_name"`
	Response Response `json:"response" mapstructure:"response"`
}
