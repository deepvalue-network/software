package standard

type misc struct {
	isConcat          bool
	isFrameAssignment bool
}

func createMiscWithConcatenation() Misc {
	return createMiscInternally(true, false)
}

func createMiscWithFrameAssignment() Misc {
	return createMiscInternally(false, true)
}

func createMiscInternally(isConcat bool, isFrameAssignment bool) Misc {
	out := misc{
		isConcat:          isConcat,
		isFrameAssignment: isFrameAssignment,
	}

	return &out
}

// IsConcatenation returns true if concatenation, false otherwise
func (obj *misc) IsConcatenation() bool {
	return obj.isConcat
}

// IsFrameAssignment returns true if frameAssignment, false otherwise
func (obj *misc) IsFrameAssignment() bool {
	return obj.isFrameAssignment
}
