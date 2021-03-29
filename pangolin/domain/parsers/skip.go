package parsers

type skip struct {
	pointer IntPointer
}

func createSkip(
	pointer IntPointer,
) Skip {
	out := skip{
		pointer: pointer,
	}

	return &out
}

// Pointer returns the pointer
func (obj *skip) Pointer() IntPointer {
	return obj.pointer
}
