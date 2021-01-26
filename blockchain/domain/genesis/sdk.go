package genesis

import (
	"github.com/steve-care-software/products/libs/hash"
)

// DefaultMiningValue represents the default mining value
const DefaultMiningValue = 0x0

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// Builder represenst a genesis builder
type Builder interface {
	Create() Builder
	WithMiningValue(miningValue uint8) Builder
	WithBlockBaseDifficulty(blockBaseDiff uint) Builder
	WithBlockIncreasePerHashDifficulty(incrPerHashDiff float64) Builder
	WithLinkDifficulty(linkDiff uint) Builder
	Now() (Genesis, error)
}

// Genesis represents a genesis
type Genesis interface {
	Hash() hash.Hash
	MiningValue() uint8
	Difficulty() Difficulty
}

// Difficulty represents the genesis difficulty
type Difficulty interface {
	Block() Block
	Link() uint
}

// Block represents the block difficulty related data
type Block interface {
	Base() uint
	IncreasePerHash() float64
}
