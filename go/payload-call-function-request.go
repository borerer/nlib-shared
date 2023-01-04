package nlibshared

type PayloadCallFunctionRequest struct {
	Name    string  `json:"name"`
	Request Request `json:"request"`
}
