package nlibshared

type PayloadCallFunctionRequest struct {
	FuncName string      `json:"func_name"`
	UseHAR   bool        `json:"use_har"`
	Request  interface{} `json:"request"`
}
