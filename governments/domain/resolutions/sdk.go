package resolutions

import (
	"time"

	"github.com/deepvalue-network/software/governments/domain/propositions"
	"github.com/deepvalue-network/software/governments/domain/propositions/votes"
	"github.com/deepvalue-network/software/libs/hash"
)

// NewBuilder creates a new resolution builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// Builder represents a resolution builder
type Builder interface {
	Create() Builder
	WithProposition(propositon propositions.Proposition) Builder
	WithVotes(votes []votes.Vote) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Resolution, error)
}

// Resolution represents a resolution
type Resolution interface {
	Hash() hash.Hash
	Proposition() propositions.Proposition
	Votes() []votes.Vote
	CreatedOn() time.Time
}
