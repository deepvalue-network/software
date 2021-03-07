package parsers

type assert struct {
	condition Identifier
}

func createAssert() Assert {
	return createAssertInternally(nil)
}

func createAssertWithCondition(
	condition Identifier,
) Assert {
	return createAssertInternally(condition)
}

func createAssertInternally(
	condition Identifier,
) Assert {
	out := assert{
		condition: condition,
	}

	return &out
}

// HasCondition returns true if there is a condition, false otherwise
func (obj *assert) HasCondition() bool {
	return obj.condition != nil
}

// Condition returns the condition, if any
func (obj *assert) Condition() Identifier {
	return obj.condition
}
