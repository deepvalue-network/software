package parsers

type pop struct {
	stackFrame TransformOperation
}

func createPop() Pop {
	return createPopInternally(nil)
}

func createPopWithStackframe(stackFrame TransformOperation) Pop {
	return createPopInternally(stackFrame)
}

func createPopInternally(stackFrame TransformOperation) Pop {
	out := pop{
		stackFrame: stackFrame,
	}

	return &out
}

// HasStackFrame returns true if there is a stackFrame, false otherwise
func (obj *pop) HasStackFrame() bool {
	return obj.stackFrame != nil
}

// StackFrame returns the stackframe, if any
func (obj *pop) StackFrame() TransformOperation {
	return obj.stackFrame
}
