package parsers

type remainingOperation struct {
	first     string
	second    string
	result    string
	remaining string
}

func createRemainingOperation(first string, second string, result string, remaining string) RemainingOperation {
	out := remainingOperation{
		first:     first,
		second:    second,
		result:    result,
		remaining: remaining,
	}

	return &out
}

// First returns the first identifier of the operation
func (obj *remainingOperation) First() string {
	return obj.first
}

// Second returns the second identifier of the operation
func (obj *remainingOperation) Second() string {
	return obj.second
}

// Result returns the result variable of the operation
func (obj *remainingOperation) Result() string {
	return obj.result
}

// Remaining returns the remaining variable of the operation
func (obj *remainingOperation) Remaining() string {
	return obj.remaining
}
