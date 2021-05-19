package remaining

type operation struct {
	ary Arythmetic
}

func createOperationWithArythmetic(ary Arythmetic) Operation {
	return createOperationInternally(ary)
}

func createOperationInternally(ary Arythmetic) Operation {
	out := operation{
		ary: ary,
	}

	return &out
}

// IsArythmetic returns true if the operation is an arythmetic, false otherwise
func (obj *operation) IsArythmetic() bool {
	return obj.ary != nil
}

// Arythmetic returns the arythmetic operation, if any
func (obj *operation) Arythmetic() Arythmetic {
	return obj.ary
}
