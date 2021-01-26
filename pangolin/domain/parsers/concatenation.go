package parsers

type concatenation struct {
	operation StandardOperation
}

func createConcatenation(operation StandardOperation) Concatenation {
	out := concatenation{
		operation: operation,
	}

	return &out
}

// Operation returns the standard operation
func (obj *concatenation) Operation() StandardOperation {
	return obj.operation
}
