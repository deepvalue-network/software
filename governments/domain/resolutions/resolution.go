package resolutions

import (
	"time"

	"github.com/deepvalue-network/software/governments/domain/propositions"
	"github.com/deepvalue-network/software/governments/domain/propositions/votes"
	"github.com/deepvalue-network/software/libs/hash"
)

type resolution struct {
	hash      hash.Hash
	prop      propositions.Proposition
	votes     []votes.Vote
	createdOn time.Time
}

func createResolution(
	hash hash.Hash,
	prop propositions.Proposition,
	votes []votes.Vote,
	createdOn time.Time,
) Resolution {
	out := resolution{
		hash:      hash,
		prop:      prop,
		votes:     votes,
		createdOn: createdOn,
	}

	return &out
}

// Hash returns the hash
func (obj *resolution) Hash() hash.Hash {
	return obj.hash
}

// Proposition returns the proposition
func (obj *resolution) Proposition() propositions.Proposition {
	return obj.prop
}

// Votes returns the votes
func (obj *resolution) Votes() []votes.Vote {
	return obj.votes
}

// CreatedOn returns the creation time
func (obj *resolution) CreatedOn() time.Time {
	return obj.createdOn
}
