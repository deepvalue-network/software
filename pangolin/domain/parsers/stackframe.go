package parsers

type stackFrame struct {
	push Push
	pop  Pop
}

func createStackFrameWithPush(push Push) StackFrame {
	return createStackFrameInternally(push, nil)
}

func createStackFrameWithPop(pop Pop) StackFrame {
	return createStackFrameInternally(nil, pop)
}

func createStackFrameInternally(
	push Push,
	pop Pop,
) StackFrame {
	out := stackFrame{
		push: push,
		pop:  pop,
	}

	return &out
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
