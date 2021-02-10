package genesis

import "github.com/deepvalue-network/software/libs/hash"

type genesis struct {
	hash                           hash.Hash
	miningValue                    uint8   `hydro:"MiningValue, MiningValue"`
	blockBaseDifficulty            uint    `hydro:"BlockBaseDifficulty, BlockBaseDifficulty"`
	blockIncreasePerHashDifficulty float64 `hydro:"BlockIncreasePerHashDifficulty, BlockIncreasePerHashDifficulty"`
	linkDifficulty                 uint    `hydro:"LinkDifficulty, LinkDifficulty"`
}

func createGenesis(
	hash hash.Hash,
	miningValue uint8,
	blockBaseDifficulty uint,
	blockIncreasePerHashDifficulty float64,
	linkDifficulty uint,
) Genesis {
	out := genesis{
		hash:                           hash,
		miningValue:                    miningValue,
		blockBaseDifficulty:            blockBaseDifficulty,
		blockIncreasePerHashDifficulty: blockIncreasePerHashDifficulty,
		linkDifficulty:                 linkDifficulty,
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

// BlockBaseDifficulty returns the block base difficulty
func (obj *genesis) BlockBaseDifficulty() uint {
	return obj.blockBaseDifficulty
}

// BlockIncreasePerHashDifficulty returns the block increase per hash difficulty
func (obj *genesis) BlockIncreasePerHashDifficulty() float64 {
	return obj.blockIncreasePerHashDifficulty
}

// LinkDifficulty returns the link difficulty
func (obj *genesis) LinkDifficulty() uint {
	return obj.linkDifficulty
}
