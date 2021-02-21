package governments

import (
	"time"

	"github.com/deepvalue-network/software/blockchain/domain/chains"
	"github.com/deepvalue-network/software/governments/domain/governments/shareholders"
	"github.com/deepvalue-network/software/libs/hash"
	uuid "github.com/satori/go.uuid"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// NewContentBuilder creates a new content builder instance
func NewContentBuilder() ContentBuilder {
	hashAdapter := hash.NewAdapter()
	return createContentBuilder(hashAdapter)
}

// Builder represents a government builder
type Builder interface {
	Create() Builder
	WithID(id *uuid.UUID) Builder
	WithCurrent(current Content) Builder
	WithShareHolders(shareholders shareholders.ShareHolders) Builder
	WithPrevious(prev Government) Builder
	Now() (Government, error)
}

// Government represents a government
type Government interface {
	Hash() hash.Hash
	ID() *uuid.UUID
	Current() Content
	ShareHolders() shareholders.ShareHolders
	HasPrevious() bool
	Previous() Government
}

// ContentBuilder represents a content builder
type ContentBuilder interface {
	Create() ContentBuilder
	WithChain(chain chains.Chain) ContentBuilder
	WithMinPowerToPassResolution(minPowerToPassRes uint) ContentBuilder
	WithMinPowerToPropose(minPowerToPropose uint) ContentBuilder
	WithSharesVelocity(sharesVelocity uint) ContentBuilder
	WithSharesCap(sharesCap uint) ContentBuilder
	CanCancelVote() ContentBuilder
	CreatedOn(createdOn time.Time) ContentBuilder
	Now() (Content, error)
}

// Content represents a government content
type Content interface {
	Hash() hash.Hash
	Chain() chains.Chain
	MinPowerToPassResolution() uint
	MinPowerToPropose() uint
	CanCancelVote() bool
	SharesVelocity() uint
	SharesCap() uint
	CreatedOn() time.Time
}

// Repository represents a government repository
type Repository interface {
	Retrieve(id *uuid.UUID) (Government, error)
}
