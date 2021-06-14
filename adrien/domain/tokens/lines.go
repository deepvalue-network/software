package tokens

type lines struct {
	list []Line
}

func createLines(
	list []Line,
) Lines {
	out := lines{
		list: list,
	}

	return &out
}

// All returns all lines
func (obj *lines) All() []Line {
	return obj.list
}
