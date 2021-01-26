package variablename

type misc struct {
	isPush bool
}

func createMiscWithPop() Misc {
	return createMiscInternally(true)
}

func createMiscInternally(isPush bool) Misc {
	out := misc{
		isPush: isPush,
	}

	return &out
}

// IsPush returns true if pop, false otherwise
func (obj *misc) IsPush() bool {
	return obj.isPush
}
