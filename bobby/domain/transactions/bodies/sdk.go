package bodies

import (
	"time"

	"github.com/deepvalue-network/software/bobby/domain/resources"
	"github.com/deepvalue-network/software/bobby/domain/transactions/bodies/access"
	"github.com/deepvalue-network/software/bobby/domain/transactions/bodies/containers"
	"github.com/deepvalue-network/software/bobby/domain/transactions/bodies/contents"
	"github.com/deepvalue-network/software/libs/hash"
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
