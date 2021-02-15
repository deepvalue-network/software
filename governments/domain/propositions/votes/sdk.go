package votes

import (
	"time"

	"github.com/deepvalue-network/software/governments/domain/propositions"
	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
	"github.com/deepvalue-network/software/libs/hash"
)

// Builder represents a vote builder
type Builder interface {
	Create() Builder
	WithContent(content Content) Builder
	WithSignature(sig signature.RingSignature) Builder
	Now() (Vote, error)
}

// Vote represents a proposition vote
type Vote interface {
	Hash() hash.Hash
	Content() Content
	Signature() signature.RingSignature
}

// ContentBuilder represents a content builder
type ContentBuilder interface {
	Create() ContentBuilder
	WithProposition(proposition propositions.Proposition) ContentBuilder
	IsApproved() ContentBuilder
	IsCancel() ContentBuilder
	IsDisapproved() ContentBuilder
	CreatedOn(createdOn time.Time) ContentBuilder
	Now() (Content, error)
}

// Content represents a vote content
type Content interface {
	Hash() hash.Hash
	Proposition() propositions.Proposition
	IsApproved() bool
	IsCancel() bool
	IsDisapproved() bool
	CreatedOn() time.Time
}