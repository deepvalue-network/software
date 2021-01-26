package remaining

type remaining struct {
	operation Operation
	result    string
	remaining string
	first     string
	second    string
}

func createRemaining(
	operation Operation,
	result string,
	rem string,
	first string,
	second string,
) Remaining {
	out := remaining{
		operation: operation,
		result:    result,
		remaining: rem,
		first:     first,
		second:    second,
	}

	return &out
}

// Operation returns the operation
func (obj *remaining) Operation() Operation {
	return obj.operation
}

// Result returns the result
func (obj *remaining) Result() string {
	return obj.result
}

// Remaining returns the remaining
func (obj *remaining) Remaining() string {
	return obj.remaining
}

// First returns the first
func (obj *remaining) First() string {
	return obj.first
}

// Second returns the second
func (obj *remaining) Second() string {
	return obj.second
}
