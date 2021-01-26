package bodies

import (
	"time"

	"github.com/steve-care-software/products/bobby/domain/resources"
	"github.com/steve-care-software/products/bobby/domain/transactions/bodies/access"
	"github.com/steve-care-software/products/bobby/domain/transactions/bodies/containers"
	"github.com/steve-care-software/products/bobby/domain/transactions/bodies/contents"
	"github.com/steve-care-software/products/libs/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	immutableBuilder := resources.NewImmutableBuilder()
	return createBuilder(hashAdapter, immutableBuilder)
}

// Builder represents the transaction body
type Builder interface {
	Create() Builder
	WithContainer(container containers.Transaction) Builder
	WithContent(content contents.Transaction) Builder
	WithAccess(access access.Transaction) Builder
	ExecutesOn(executesOn time.Time) Builder
	Now() (Body, error)
}

// Body represents an unsigned body of a transaction
type Body interface {
	Resource() resources.Immutable
	Content() Content
	HasExecutesOn() bool
	ExecutesOn() *time.Time
}

// Content represents the content of a body transaction
type Content interface {
	Hash() hash.Hash
	IsContainer() bool
	Container() containers.Transaction
	IsContent() bool
	Content() contents.Transaction
	IsAccess() bool
	Access() access.Transaction
}
