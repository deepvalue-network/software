package parsers

type call struct {
	name      string
	condition Identifier
}

func createCall(name string) Call {
	return createCallInternally(name, nil)
}

func createCallWithCondition(name string, condition Identifier) Call {
	return createCallInternally(name, condition)
}

func createCallInternally(
	name string,
	condition Identifier,
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
	return obj.condition != nil
}

// Condition returns the condition, if any
func (obj *call) Condition() Identifier {
	return obj.condition
}
