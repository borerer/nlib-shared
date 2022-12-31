package nlibshared

type PayloadRegisterFunctionRequest struct {
	Name   string `json:"name"`
	UseHAR bool   `json:"use_har"`
}
