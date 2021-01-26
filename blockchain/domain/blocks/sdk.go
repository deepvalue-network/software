package blocks

import (
	"github.com/steve-care-software/products/libs/hash"
	"github.com/steve-care-software/products/libs/hashtree"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashTreeBuilder := hashtree.NewBuilder()
	return createBuilder(hashTreeBuilder)
}

// Builder represents a block builder
type Builder interface {
	Create() Builder
	WithHashes(hashes []hash.Hash) Builder
	Now() (Block, error)
}

// Block represents a block
type Block interface {
	Tree() hashtree.HashTree
	Hashes() []hash.Hash
}

// Repository represents a block repository
type Repository interface {
	List() ([]hash.Hash, error)
	Retrieve(blockHash hash.Hash) (Block, error)
}

// Service represents a block service
type Service interface {
	Insert(block Block) error
	Delete(block Block) error
}
