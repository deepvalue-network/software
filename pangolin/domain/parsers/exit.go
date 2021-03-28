package parsers

type exit struct {
	condition string
}

func createExit() Exit {
	return createExitInternally("")
}

func createExitWithCondition(condition string) Exit {
	return createExitInternally(condition)
}

func createExitInternally(
	condition string,
) Exit {
	out := exit{
		condition: condition,
	}

	return &out
}

// HasCondition returns true if there is a condition, false otherwise
func (obj *exit) HasCondition() bool {
	return obj.condition != ""
}

// Condition returns the condition, if any
func (obj *exit) Condition() string {
	return obj.condition
}
