package nlibshared

type PayloadRegisterFunctionRequest struct {
	FuncName string `json:"func_name" mapstructure:"func_name"`
	UseHAR   bool   `json:"use_har" mapstructure:"use_har"`
}
