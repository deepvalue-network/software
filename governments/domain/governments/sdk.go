package governments

import (
	"time"

	"github.com/deepvalue-network/software/blockchain/domain/chains"
	"github.com/deepvalue-network/software/libs/hash"
	uuid "github.com/satori/go.uuid"
)

// Builder represents a government builder
type Builder interface {
	Create() Builder
	WithCurrent(current Content) Builder
	WithPrevious(prev Government) Builder
	Now() (Government, error)
}

// Government represents a government
type Government interface {
	Hash() hash.Hash
	Current() Content
	HasPrevious() bool
	Previous() Government
}

// ContentBuilder represents a content builder
type ContentBuilder interface {
	Create() ContentBuilder
	WithID(id *uuid.UUID) ContentBuilder
	WithChain(chain chains.Chain) ContentBuilder
	WithMinPowerToPassResolution(minPowerToPassRes uint) ContentBuilder
	WithMinPowerToPropose(minPowerToPropose uint) ContentBuilder
	WithSharesVelocity(sharesVelocity uint) ContentBuilder
	WithSharesCap(sharesCap uint) ContentBuilder
	CanCancelVote() ContentBuilder
	BurnOnReceive() ContentBuilder
	CreatedOn(createdOn time.Time) ContentBuilder
	Now() (Content, error)
}

// Content represents a government content
type Content interface {
	Hash() hash.Hash
	ID() *uuid.UUID
	Chain() chains.Chain
	MinPowerToPassResolution() uint
	MinPowerToPropose() uint
	CanCancelVote() bool
	BurnOnReceive() bool
	SharesVelocity() uint
	SharesCap() uint
	CreatedOn() time.Time
}
