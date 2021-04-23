package instruction

type assert struct {
	index     int
	condition string
}

func createAssert(index int) Assert {
	return createAssertInternally(index, "")
}

func createAssertWithCondition(index int, condition string) Assert {
	return createAssertInternally(index, condition)
}

func createAssertInternally(
	index int,
	condition string,
) Assert {
	out := assert{
		index:     index,
		condition: condition,
	}

	return &out
}

// Index returns the index
func (obj *assert) Index() int {
	return obj.index
}

// HasCondition returns true if there is a condition, false otherwise
func (obj *assert) HasCondition() bool {
	return obj.condition != ""
}

// Condition returns the condition, if any
func (obj *assert) Condition() string {
	return obj.condition
}
