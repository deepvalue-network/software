package links

import (
	"github.com/deepvalue-network/software/blockchain/domain/blocks"
	"github.com/deepvalue-network/software/libs/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// NewPointer creates a new pointer instance
func NewPointer() *link {
	return new(link)
}

// Builder represents a link builder
type Builder interface {
	Create() Builder
	WithIndex(index uint) Builder
	WithPreviousMinedLink(prevMinedLink hash.Hash) Builder
	WithNextBlock(block blocks.Block) Builder
	Now() (Link, error)
}

// Link represents a linked block
type Link interface {
	Hash() hash.Hash
	Index() uint
	PrevMinedLink() hash.Hash
	NextBlock() blocks.Block
}

// Repository represents a link repository
type Repository interface {
	List() ([]hash.Hash, error)
	Retrieve(linkHash hash.Hash) (Link, error)
}

// Service represents a link service
type Service interface {
	Insert(link Link) error
	Delete(link Link) error
}
