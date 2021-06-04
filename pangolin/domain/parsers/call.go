package parsers

type call struct {
	name       string
	stackFrame string
	condition  string
}

func createCall(name string, stackFrame string) Call {
	return createCallInternally(name, stackFrame, "")
}

func createCallWithCondition(name string, stackFrame string, condition string) Call {
	return createCallInternally(name, stackFrame, condition)
}

func createCallInternally(
	name string,
	stackFrame string,
	condition string,
) Call {
	out := call{
		name:       name,
		stackFrame: stackFrame,
		condition:  condition,
	}

	return &out
}

// Name returns the name
func (obj *call) Name() string {
	return obj.name
}

// StackFrame returns the stackFrame
func (obj *call) StackFrame() string {
	return obj.stackFrame
}

// HasCondition returns true if there is a condition, false otherwise
func (obj *call) HasCondition() bool {
	return obj.condition != ""
}

// Condition returns the condition, if any
func (obj *call) Condition() string {
	return obj.condition
}
