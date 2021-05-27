package condition

type condition struct {
	proposition Proposition
	operation   Operation
}

func createCondition(proposition Proposition, operation Operation) Condition {
	out := condition{
		proposition: proposition,
		operation:   operation,
	}

	return &out
}

// Proposition returns the proposition
func (obj *condition) Proposition() Proposition {
	return obj.proposition
}

// Operation returns the operation
func (obj *condition) Operation() Operation {
	return obj.operation
}
