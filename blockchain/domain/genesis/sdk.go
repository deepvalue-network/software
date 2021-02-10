package genesis

import (
	"github.com/deepvalue-network/software/libs/hash"
)

// DefaultMiningValue represents the default mining value
const DefaultMiningValue = 0x0

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// NewPointer returns a genesis pointer
func NewPointer() *genesis {
	return new(genesis)
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
	BlockBaseDifficulty() uint
	BlockIncreasePerHashDifficulty() float64
	LinkDifficulty() uint
}
