package resolutions

import (
	"time"

	"github.com/deepvalue-network/software/governments/domain/propositions"
	"github.com/deepvalue-network/software/governments/domain/propositions/votes"
	"github.com/deepvalue-network/software/libs/hash"
)

type content struct {
	hash      hash.Hash
	prop      propositions.Proposition
	votes     []votes.Vote
	createdOn time.Time
}

func createContent(
	hash hash.Hash,
	prop propositions.Proposition,
	votes []votes.Vote,
	createdOn time.Time,
) Content {
	out := content{
		hash:      hash,
		prop:      prop,
		votes:     votes,
		createdOn: createdOn,
	}

	return &out
}

// Hash returns the hash
func (obj *content) Hash() hash.Hash {
	return obj.hash
}

// Proposition returns the proposition
func (obj *content) Proposition() propositions.Proposition {
	return obj.prop
}

// Votes returns the votes
func (obj *content) Votes() []votes.Vote {
	return obj.votes
}

// CreatedOn returns the creation time
func (obj *content) CreatedOn() time.Time {
	return obj.createdOn
}
