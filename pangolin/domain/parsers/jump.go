package parsers

type jump struct {
	label     string
	condition string
}

func createJump(label string) Jump {
	return createJumpInternally(label, "")
}

func createJumpWithCondition(label string, condition string) Jump {
	return createJumpInternally(label, condition)
}

func createJumpInternally(label string, condition string) Jump {
	out := jump{
		label:     label,
		condition: condition,
	}

	return &out
}

// Label returns the label
func (obj *jump) Label() string {
	return obj.label
}

// HasCondition returns true if there is a condition, false otherwise
func (obj *jump) HasCondition() bool {
	return obj.condition != ""
}

// Condition returns the condition, if any
func (obj *jump) Condition() string {
	return obj.condition
}
