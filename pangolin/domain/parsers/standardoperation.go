package parsers

type standardOperation struct {
	first  Identifier
	second Identifier
	result VariableName
}

func createStandardOperation(first Identifier, second Identifier, result VariableName) StandardOperation {
	out := standardOperation{
		first:  first,
		second: second,
		result: result,
	}

	return &out
}

// First returns the first identifier of the operation
func (obj *standardOperation) First() Identifier {
	return obj.first
}

// Second returns the second identifier of the operation
func (obj *standardOperation) Second() Identifier {
	return obj.second
}

// Result returns the result variable of the operation
func (obj *standardOperation) Result() VariableName {
	return obj.result
}
