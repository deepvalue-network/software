package remaining

type misc struct {
	isMatch bool
}

func createMiscWithMatch() Misc {
	return createMiscInternally(true)
}

func createMiscInternally(isMatch bool) Misc {
	out := misc{
		isMatch: isMatch,
	}

	return &out
}

// IsMatch returns true if match, false otherwise
func (obj *misc) IsMatch() bool {
	return obj.isMatch
}
