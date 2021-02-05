package mined

import (
	"time"

	"github.com/deepvalue-network/software/blockchain/domain/links"
	"github.com/deepvalue-network/software/libs/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
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
	Retrieve(linkHash hash.Hash) (Link, error)
}

// Service represents a link service
type Service interface {
	Insert(link Link) error
	Delete(link Link) error
}
