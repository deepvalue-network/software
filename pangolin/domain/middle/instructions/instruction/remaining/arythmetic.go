package remaining

type arythmetic struct {
	isDiv bool
}

func createArythmeticWithDiv() Arythmetic {
	return createArythmeticInternally(true)
}

func createArythmeticInternally(isDiv bool) Arythmetic {
	out := arythmetic{
		isDiv: isDiv,
	}

	return &out
}

// IsDiv returns true if a division
func (obj *arythmetic) IsDiv() bool {
	return obj.isDiv
}
