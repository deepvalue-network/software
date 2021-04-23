package call

type call struct {
	name      string
	condition string
}

func createCall(
	name string,
) Call {
	return createCallInternally(name, "")
}

func createCallWithCondition(
	name string,
	condition string,
) Call {
	return createCallInternally(name, condition)
}

func createCallInternally(
	name string,
	condition string,
) Call {
	out := call{
		name:      name,
		condition: condition,
	}

	return &out
}

// Name returns the name
func (obj *call) Name() string {
	return obj.name
}

// HasCondition returns true if there is a condition, false otherwise
func (obj *call) HasCondition() bool {
	return obj.condition != ""
}

// Condition returns the condition, if any
func (obj *call) Condition() string {
	return obj.condition
}
