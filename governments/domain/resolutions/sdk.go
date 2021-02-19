package resolutions

import (
	"time"

	"github.com/deepvalue-network/software/governments/domain/propositions"
	"github.com/deepvalue-network/software/governments/domain/propositions/votes"
	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
	"github.com/deepvalue-network/software/libs/hash"
)

// NewContentBuilder creates a new content builder instance
func NewContentBuilder() ContentBuilder {
	hashAdapter := hash.NewAdapter()
	return createContentBuilder(hashAdapter)
}

// Builder represents a resolution builder
type Builder interface {
	Create() Builder
	WithContent(content Content) Builder
	WithSignatures(sigs []signature.RingSignature) Builder
	Now() (Resolution, error)
}

// Resolution represents a resolution
type Resolution interface {
	Hash() hash.Hash
	Content() Content
	Signatures() []signature.RingSignature
}

// ContentBuilder represents a content builder
type ContentBuilder interface {
	Create() ContentBuilder
	WithProposition(propositon propositions.Proposition) ContentBuilder
	WithVotes(votes []votes.Vote) ContentBuilder
	CreatedOn(createdOn time.Time) ContentBuilder
	Now() (Content, error)
}

// Content represents a resolution content
type Content interface {
	Hash() hash.Hash
	Proposition() propositions.Proposition
	Votes() []votes.Vote
	CreatedOn() time.Time
}
