package parsers

type stackFrame struct {
	isPush bool
	isPop  bool
	index  Index
	skip   Skip
}

func createStackFrameWithPush() StackFrame {
	return createStackFrameInternally(true, false, nil, nil)
}

func createStackFrameWithPop() StackFrame {
	return createStackFrameInternally(false, true, nil, nil)
}

func createStackFrameWithIndex(index Index) StackFrame {
	return createStackFrameInternally(false, false, index, nil)
}

func createStackFrameWithSkip(skip Skip) StackFrame {
	return createStackFrameInternally(false, false, nil, skip)
}

func createStackFrameInternally(
	isPush bool,
	isPop bool,
	index Index,
	skip Skip,
) StackFrame {
	out := stackFrame{
		isPush: isPush,
		isPop:  isPop,
		index:  index,
		skip:   skip,
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

// IsIndex returns true if there is an index, false otherwise
func (obj *stackFrame) IsIndex() bool {
	return obj.index != nil
}

// Index returns the index, if any
func (obj *stackFrame) Index() Index {
	return obj.index
}

// IsSkip returns true if there is a skip, false otherwise
func (obj *stackFrame) IsSkip() bool {
	return obj.skip != nil
}

// Skip returns the skip, if any
func (obj *stackFrame) Skip() Skip {
	return obj.skip
}
