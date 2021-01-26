package genesis

import "github.com/steve-care-software/products/libs/hash"

type genesis struct {
	hash        hash.Hash
	miningValue uint8
	diff        Difficulty
}

func createGenesis(
	hash hash.Hash,
	miningValue uint8,
	diff Difficulty,
) Genesis {
	out := genesis{
		hash:        hash,
		miningValue: miningValue,
		diff:        diff,
	}

	return &out
}

// Hash returns the hash
func (obj *genesis) Hash() hash.Hash {
	return obj.hash
}

// MiningValue returns the mining value
func (obj *genesis) MiningValue() uint8 {
	return obj.miningValue
}

// Difficulty returns the difficulty
func (obj *genesis) Difficulty() Difficulty {
	return obj.diff
}
