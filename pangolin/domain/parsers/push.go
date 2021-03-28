package parsers

type push struct {
	stackFrame string
}

func createPush() Push {
	return createPushInternally("")
}

func createPushWithStackframe(stackFrame string) Push {
	return createPushInternally(stackFrame)
}

func createPushInternally(stackFrame string) Push {
	out := push{
		stackFrame: stackFrame,
	}

	return &out
}

// HasStackFrame returns true if there is a stackFrame, false otherwise
func (obj *push) HasStackFrame() bool {
	return obj.stackFrame != ""
}

// StackFrame returns the stackframe, if any
func (obj *push) StackFrame() string {
	return obj.stackFrame
}
