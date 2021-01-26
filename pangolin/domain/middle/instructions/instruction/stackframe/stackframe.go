package stackframe

type stackframe struct {
	isPush bool
	isPop  bool
}

func createStackframeWithPush() Stackframe {
	return createStackframeInternally(true, false)
}

func createStackframeWithPop() Stackframe {
	return createStackframeInternally(false, true)
}

func createStackframeInternally(isPush bool, isPop bool) Stackframe {
	out := stackframe{
		isPush: isPush,
		isPop:  isPop,
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
