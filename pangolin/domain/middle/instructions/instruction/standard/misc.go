package standard

type misc struct {
	isConcat bool
}

func createMiscWithConcatenation() Misc {
	return createMiscInternally(true)
}

func createMiscInternally(isConcat bool) Misc {
	out := misc{
		isConcat: isConcat,
	}

	return &out
}

// IsConcatenation returns true if concatenation, false otherwise
func (obj *misc) IsConcatenation() bool {
	return obj.isConcat
}
