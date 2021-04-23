package remaining

type operation struct {
	ary  Arythmetic
	misc Misc
}

func createOperationWithArythmetic(ary Arythmetic) Operation {
	return createOperationInternally(ary, nil)
}

func createOperationWithMisc(misc Misc) Operation {
	return createOperationInternally(nil, misc)
}

func createOperationInternally(ary Arythmetic, misc Misc) Operation {
	out := operation{
		ary:  ary,
		misc: misc,
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

// IsMisc returns true if the operation is a misc, false otherwise
func (obj *operation) IsMisc() bool {
	return obj.misc != nil
}

// Misc returns the misc operation, if any
func (obj *operation) Misc() Misc {
	return obj.misc
}
