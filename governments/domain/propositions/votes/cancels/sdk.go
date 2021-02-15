package cancels

import (
	"time"

	"github.com/deepvalue-network/software/governments/domain/propositions/votes"
	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
	"github.com/deepvalue-network/software/libs/hash"
)

// Builder represents a cancel builder
type Builder interface {
	Create() Builder
	WithContent(content Content) Builder
	WithSignature(sig signature.RingSignature) Builder
	Now() (Cancel, error)
}

// Cancel represents a vote cancel
type Cancel interface {
	Hash() hash.Hash
	Content() Content
	Signature() signature.RingSignature
}

// ContentBuilder represents a content builder
type ContentBuilder interface {
	Create() ContentBuilder
	WithVote(vote votes.Vote) ContentBuilder
	CreatedOn(createdOn time.Time) ContentBuilder
	Now() (Content, error)
}

// Content represents a cancel content
type Content interface {
	Hash() hash.Hash
	Vote() votes.Vote
	CreatedOn() time.Time
}
