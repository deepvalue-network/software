package parsers

type remainingOperation struct {
	first     Identifier
	second    Identifier
	result    VariableName
	remaining VariableName
}

func createRemainingOperation(first Identifier, second Identifier, result VariableName, remaining VariableName) RemainingOperation {
	out := remainingOperation{
		first:     first,
		second:    second,
		result:    result,
		remaining: remaining,
	}

	return &out
}

// First returns the first identifier of the operation
func (obj *remainingOperation) First() Identifier {
	return obj.first
}

// Second returns the second identifier of the operation
func (obj *remainingOperation) Second() Identifier {
	return obj.second
}

// Result returns the result variable of the operation
func (obj *remainingOperation) Result() VariableName {
	return obj.result
}

// Remaining returns the remaining variable of the operation
func (obj *remainingOperation) Remaining() VariableName {
	return obj.remaining
}
