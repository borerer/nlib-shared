package nlibshared

type FunctionIn = Request
type FunctionOut = Response
type Function = func(*FunctionIn) (*FunctionOut, error)
