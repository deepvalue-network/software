package mined

import (
	"time"

	"github.com/deepvalue-network/software/blockchain/domain/blocks"
	"github.com/deepvalue-network/software/blockchain/domain/genesis"
	"github.com/deepvalue-network/software/blockchain/domain/links"
	"github.com/deepvalue-network/software/libs/hash"
)

// NewValidator creates a new validator instance
func NewValidator(minedLinkRepository Repository) Validator {
	hashAdapter := hash.NewAdapter()
	return createValidator(hashAdapter, minedLinkRepository)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// NewPointer creates a new pointer instance
func NewPointer() *link {
	return new(link)
}

// Validator represents a mined link validator
type Validator interface {
	Execute(gen genesis.Genesis, minedLink Link, root blocks.Block) (uint, uint, error)
}

// Builder represenst the link builder
type Builder interface {
	Create() Builder
	WithLink(link links.Link) Builder
	WithResults(results string) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Link, error)
}

// Link represents a mined link
type Link interface {
	Hash() hash.Hash
	Link() links.Link
	Results() string
	CreatedOn() time.Time
}

// Repository represents a link repository
type Repository interface {
	Head() (Link, error)
	List() ([]hash.Hash, error)
	Retrieve(minedLinkHash hash.Hash) (Link, error)
	RetrieveByLinkHash(linkHash hash.Hash) (Link, error)
}

// Service represents a link service
type Service interface {
	Insert(minedLink Link) error
	Delete(minedLink Link) error
	DeleteByLink(link links.Link) error
}
