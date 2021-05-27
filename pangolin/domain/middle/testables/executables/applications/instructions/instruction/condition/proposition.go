package condition

type proposition struct {
	name      string
	condition string
}

func createProposition(name string) Proposition {
	return createPropositionInternally(name, "")
}

func createPropositionWithCondition(name string, condition string) Proposition {
	return createPropositionInternally(name, condition)
}

func createPropositionInternally(name string, condition string) Proposition {
	out := proposition{
		name:      name,
		condition: condition,
	}

	return &out
}

// Name returns the name
func (obj *proposition) Name() string {
	return obj.name
}

// HasCondition returns true if there is a condition, false otherwise
func (obj *proposition) HasCondition() bool {
	return obj.condition != ""
}

// Condition returns the condition, if any
func (obj *proposition) Condition() string {
	return obj.condition
}
