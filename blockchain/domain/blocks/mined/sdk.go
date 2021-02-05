package mined

import (
	"time"

	"github.com/deepvalue-network/software/blockchain/domain/blocks"
	"github.com/deepvalue-network/software/libs/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// NewPointer creates a new mined block pointer
func NewPointer() Block {
	return new(block)
}

// Builder represents a block builder
type Builder interface {
	Create() Builder
	WithBlock(block blocks.Block) Builder
	WithResults(results string) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Block, error)
}

// Block represents a mined block
type Block interface {
	Hash() hash.Hash
	Block() blocks.Block
	Results() string
	CreatedOn() time.Time
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
