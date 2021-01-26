package parsers

type transformOperation struct {
	input  Identifier
	result VariableName
}

func createTransformOperation(input Identifier, result VariableName) TransformOperation {
	out := transformOperation{
		input:  input,
		result: result,
	}

	return &out
}

// Input returns the input identifier
func (obj *transformOperation) Input() Identifier {
	return obj.input
}

// Result returns the result variableName
func (obj *transformOperation) Result() VariableName {
	return obj.result
}
