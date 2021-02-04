package chains

import (
	"time"

	"github.com/steve-care-software/products/blockchain/domain/chains"
	"github.com/steve-care-software/products/libs/hash"
	"github.com/steve-care-software/products/libs/hashtree"
)

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
