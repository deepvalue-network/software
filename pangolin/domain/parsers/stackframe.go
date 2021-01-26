package parsers

type stackFrame struct {
	assignment FrameAssignment
	push       Push
	pop        Pop
}

func createStackFrameWithAssignment(assignment FrameAssignment) StackFrame {
	return createStackFrameInternally(assignment, nil, nil)
}

func createStackFrameWithPush(push Push) StackFrame {
	return createStackFrameInternally(nil, push, nil)
}

func createStackFrameWithPop(pop Pop) StackFrame {
	return createStackFrameInternally(nil, nil, pop)
}

func createStackFrameInternally(
	assignment FrameAssignment,
	push Push,
	pop Pop,
) StackFrame {
	out := stackFrame{
		assignment: assignment,
		push:       push,
		pop:        pop,
	}

	return &out
}

// IsAssignment returns true if there is an assignment, false otherwise
func (obj *stackFrame) IsAssignment() bool {
	return obj.assignment != nil
}

// Assignment returns the assignment, if any
func (obj *stackFrame) Assignment() FrameAssignment {
	return obj.assignment
}

// IsPush returns true if there is a push, false otherwise
func (obj *stackFrame) IsPush() bool {
	return obj.push != nil
}

// Push returns the push, if any
func (obj *stackFrame) Push() Push {
	return obj.push
}

// IsPop returns true if there is a pop, false otherwise
func (obj *stackFrame) IsPop() bool {
	return obj.pop != nil
}

// Pop returns the pop, if any
func (obj *stackFrame) Pop() Pop {
	return obj.pop
}
