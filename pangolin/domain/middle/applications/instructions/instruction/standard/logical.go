package standard

type logical struct {
	isAnd bool
	isOr  bool
}

func createLogicalWithAnd() Logical {
	return createLogicalInternally(true, false)
}

func createLogicalWithOr() Logical {
	return createLogicalInternally(false, true)
}

func createLogicalInternally(isAnd bool, isOr bool) Logical {
	out := logical{
		isAnd: isAnd,
		isOr:  isOr,
	}

	return &out
}

// IsAnd returns true if and, false otherwise
func (obj *logical) IsAnd() bool {
	return obj.isAnd
}

// IsOr returns true if or, false otherwise
func (obj *logical) IsOr() bool {
	return obj.isOr
}
