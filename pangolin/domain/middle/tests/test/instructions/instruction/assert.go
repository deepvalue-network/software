package instruction

type assert struct {
	condition string
}

func createAssert() Assert {
	return createAssertInternally("")
}

func createAssertWithCondition(condition string) Assert {
	return createAssertInternally(condition)
}

func createAssertInternally(
	condition string,
) Assert {
	out := assert{
		condition: condition,
	}

	return &out
}

// HasCondition returns true if there is a condition, false otherwise
func (obj *assert) HasCondition() bool {
	return obj.condition != ""
}

// Condition returns the condition, if any
func (obj *assert) Condition() string {
	return obj.condition
}
