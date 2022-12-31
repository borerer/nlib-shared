package nlibshared

type PayloadCallFunctionRequest struct {
	Name    string      `json:"name"`
	UseHAR  bool        `json:"use_har"`
	Request interface{} `json:"request"`
}
