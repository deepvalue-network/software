package tokens

type line struct {
	elements []Element
}

func createLine(
	elements []Element,
) Line {
	out := line{
		elements: elements,
	}

	return &out
}

// Elements returns the elements
func (obj *line) Elements() []Element {
	return obj.elements
}
