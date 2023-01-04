package nlibshared

type SimpleFunctionIn = map[string]interface{}
type SimpleFunctionOut = interface{}
type SimpleFunction = func(*SimpleFunctionIn) (*SimpleFunctionOut, error)

type HARFunctionIn = Request
type HARFunctionOut = Response
type HARFunction = func(*HARFunctionIn) (*HARFunctionOut, error)
