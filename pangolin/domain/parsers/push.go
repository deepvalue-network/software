package parsers

type push struct {
	stackFrame VariableName
}

func createPush() Push {
	return createPushInternally(nil)
}

func createPushWithStackframe(stackFrame VariableName) Push {
	return createPushInternally(stackFrame)
}

func createPushInternally(stackFrame VariableName) Push {
	out := push{
		stackFrame: stackFrame,
	}

	return &out
}

// HasStackFrame returns true if there is a stackFrame, false otherwise
func (obj *push) HasStackFrame() bool {
	return obj.stackFrame != nil
}

// StackFrame returns the stackframe, if any
func (obj *push) StackFrame() VariableName {
	return obj.stackFrame
}
