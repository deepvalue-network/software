package parsers

type exit struct {
	condition Identifier
}

func createExit() Exit {
	return createExitInternally(nil)
}

func createExitWithCondition(condition Identifier) Exit {
	return createExitInternally(condition)
}

func createExitInternally(
	condition Identifier,
) Exit {
	out := exit{
		condition: condition,
	}

	return &out
}

// HasCondition returns true if there is a condition, false otherwise
func (obj *exit) HasCondition() bool {
	return obj.condition != nil
}

// Condition returns the condition, if any
func (obj *exit) Condition() Identifier {
	return obj.condition
}
