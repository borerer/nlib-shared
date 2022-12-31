package nlibshared

type SimpleFunctionIn = map[string]interface{}
type SimpleFunctionOut = interface{}
type SimpleFunction = func(SimpleFunctionIn) SimpleFunctionOut

type HARFunctionIn = Request
type HARFunctionOut = Response
type HARFunction = func(HARFunctionIn) HARFunctionOut
