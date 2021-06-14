package tokens

type block struct {
	must Lines
	not  Lines
}

func createBlock(
	must Lines,
) Block {
	return createBlockInternally(must, nil)
}

func createBlockWithNot(
	must Lines,
	not Lines,
) Block {
	return createBlockInternally(must, not)
}

func createBlockInternally(
	must Lines,
	not Lines,
) Block {
	out := block{
		must: must,
		not:  not,
	}

	return &out
}

// Must returns the must lines
func (obj *block) Must() Lines {
	return obj.must
}

// HasNot returns true if there is not lines, false otherwise
func (obj *block) HasNot() bool {
	return obj.not != nil
}

// Not returns the not lines, if any
func (obj *block) Not() Lines {
	return obj.not
}
