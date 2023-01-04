package nlibshared

type PayloadCallFunctionRequest struct {
	Name    string      `json:"name"`
	Request interface{} `json:"request"`
}
