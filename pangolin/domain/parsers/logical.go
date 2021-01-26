package parsers

type logical struct {
	and StandardOperation
	or  StandardOperation
}

func createLogicalWithAnd(and StandardOperation) Logical {
	return createLogicalInternally(and, nil)
}

func createLogicalWithOr(or StandardOperation) Logical {
	return createLogicalInternally(nil, or)
}

func createLogicalInternally(and StandardOperation, or StandardOperation) Logical {
	out := logical{
		and: and,
		or:  or,
	}

	return &out
}

// IsAnd returns true if the logical operator is and, false otherwise
func (obj *logical) IsAnd() bool {
	return obj.and != nil
}

// And returns the and logical operator, if any
func (obj *logical) And() StandardOperation {
	return obj.and
}

// IsOr returns true if the logical operator is or, false otherwise
func (obj *logical) IsOr() bool {
	return obj.or != nil
}

// Or returns the or logical operator, if any
func (obj *logical) Or() StandardOperation {
	return obj.or
}
