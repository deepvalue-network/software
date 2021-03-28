package parsers

type standardOperation struct {
	first  string
	second string
	result string
}

func createStandardOperation(first string, second string, result string) StandardOperation {
	out := standardOperation{
		first:  first,
		second: second,
		result: result,
	}

	return &out
}

// First returns the first identifier of the operation
func (obj *standardOperation) First() string {
	return obj.first
}

// Second returns the second identifier of the operation
func (obj *standardOperation) Second() string {
	return obj.second
}

// Result returns the result variable of the operation
func (obj *standardOperation) Result() string {
	return obj.result
}
