package transform

type transform struct {
	operation Operation
	result    string
	input     string
}

func createTransform(operation Operation, result string, input string) Transform {
	out := transform{
		operation: operation,
		result:    result,
		input:     input,
	}

	return &out
}

// Operation returns the operation
func (obj *transform) Operation() Operation {
	return obj.operation
}

// Result returns the result
func (obj *transform) Result() string {
	return obj.result
}

// Input returns the input
func (obj *transform) Input() string {
	return obj.input
}
