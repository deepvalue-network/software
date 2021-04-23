package variablename

type operation struct {
	misc Misc
}

func createOperationWithMisc(misc Misc) Operation {
	return createOperationInternally(misc)
}

func createOperationInternally(misc Misc) Operation {
	out := operation{
		misc: misc,
	}

	return &out
}

// IsMisc returns true if there is a misc, false otherwise
func (obj *operation) IsMisc() bool {
	return obj.misc != nil
}

// Misc returns the misc, if any
func (obj *operation) Misc() Misc {
	return obj.misc
}
