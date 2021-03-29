package stackframe

type stackframe struct {
	isPush bool
	isPop  bool
	index  string
	skip   Skip
}

func createStackframeWithPush() Stackframe {
	return createStackframeInternally(true, false, "", nil)
}

func createStackframeWithPop() Stackframe {
	return createStackframeInternally(false, true, "", nil)
}

func createStackframeWithIndex(index string) Stackframe {
	return createStackframeInternally(false, false, index, nil)
}

func createStackframeWithSkip(skip Skip) Stackframe {
	return createStackframeInternally(false, false, "", skip)
}

func createStackframeInternally(
	isPush bool,
	isPop bool,
	index string,
	skip Skip,
) Stackframe {
	out := stackframe{
		isPush: isPush,
		isPop:  isPop,
		index:  index,
		skip:   skip,
	}

	return &out
}

// IsPush returns true if the instruction is push
func (obj *stackframe) IsPush() bool {
	return obj.isPush
}

// IsPop returns true if the instruction is pop
func (obj *stackframe) IsPop() bool {
	return obj.isPop
}

// IsIndex returns true if there is an index, false otherwise
func (obj *stackframe) IsIndex() bool {
	return obj.index != ""
}

// Index returns the index, if any
func (obj *stackframe) Index() string {
	return obj.index
}

// IsSkip returns true if there is a skip, false otherwise
func (obj *stackframe) IsSkip() bool {
	return obj.skip != nil
}

// Skip returns the skip, if any
func (obj *stackframe) Skip() Skip {
	return obj.skip
}
