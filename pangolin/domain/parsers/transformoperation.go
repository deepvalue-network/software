package parsers

type transformOperation struct {
	input  string
	result string
}

func createTransformOperation(input string, result string) TransformOperation {
	out := transformOperation{
		input:  input,
		result: result,
	}

	return &out
}

// Input returns the input identifier
func (obj *transformOperation) Input() string {
	return obj.input
}

// Result returns the result variableName
func (obj *transformOperation) Result() string {
	return obj.result
}
