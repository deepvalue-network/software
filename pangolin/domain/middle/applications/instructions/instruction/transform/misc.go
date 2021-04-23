package transform

type misc struct {
	isPop bool
}

func createMiscWithPop() Misc {
	return createMiscInternally(true)
}

func createMiscInternally(isPop bool) Misc {
	out := misc{
		isPop: isPop,
	}

	return &out
}

// IsPop returns true if pop, false otherwise
func (obj *misc) IsPop() bool {
	return obj.isPop
}
