package mines

import (
	"time"

	"github.com/deepvalue-network/software/blockchain/domain/chains"
	"github.com/deepvalue-network/software/libs/hash"
	"github.com/deepvalue-network/software/libs/hashtree"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// Builder represents a mine builder
type Builder interface {
	Create() Builder
	WithChain(chain chains.Chain) Builder
	WithDiamonds(diamonds hashtree.HashTree) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Mine, error)
}

// Mine represents the diamond mine
type Mine interface {
	Hash() hash.Hash
	Chain() chains.Chain
	Diamonds() hashtree.HashTree
	CreatedOn() time.Time
}

// Repository represents a mine repository
type Repository interface {
	Retrieve() (Mine, error)
}

// Service represents a mine service
type Service interface {
	Insert(mine Mine) error
	Delete(mine Mine) error
}
