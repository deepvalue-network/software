package condition

type operation struct {
	isJump bool
}

func createOperationWithJump() Operation {
	return createOperationInternally(true)
}

func createOperationInternally(isJump bool) Operation {
	out := operation{
		isJump: isJump,
	}

	return &out
}

// IsJump returns true if the operation if a jump, false otherwise
func (obj *operation) IsJump() bool {
	return obj.isJump
}
