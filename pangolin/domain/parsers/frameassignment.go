package parsers

type frameAssignment struct {
	standard StandardOperation
}

func createFrameAssignment(standard StandardOperation) FrameAssignment {
	out := frameAssignment{
		standard: standard,
	}

	return &out
}

// Standard returns the standardOperation
func (obj *frameAssignment) Standard() StandardOperation {
	return obj.standard
}
