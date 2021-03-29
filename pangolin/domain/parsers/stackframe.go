package parsers

type stackFrame struct {
	isPush bool
	isPop  bool
}

func createStackFrameWithPush() StackFrame {
	return createStackFrameInternally(true, false)
}

func createStackFrameWithPop() StackFrame {
	return createStackFrameInternally(false, true)
}

func createStackFrameInternally(
	isPush bool,
	isPop bool,
) StackFrame {
	out := stackFrame{
		isPush: isPush,
		isPop:  isPop,
	}

	return &out
}

// IsPush returns true if there is a push, false otherwise
func (obj *stackFrame) IsPush() bool {
	return obj.isPush
}

// IsPop returns true if there is a pop, false otherwise
func (obj *stackFrame) IsPop() bool {
	return obj.isPop
}
