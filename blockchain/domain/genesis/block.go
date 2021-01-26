package genesis

type block struct {
	base        uint
	incrPerHash float64
}

func createBlock(
	base uint,
	incrPerHash float64,
) Block {
	out := block{
		base:        base,
		incrPerHash: incrPerHash,
	}

	return &out
}

// Base returns the base block difficulty
func (obj *block) Base() uint {
	return obj.base
}

// IncreasePerHash returns the increase per hash
func (obj *block) IncreasePerHash() float64 {
	return obj.incrPerHash
}
