package closes

import (
	"time"

	"github.com/deepvalue-network/software/governments/domain/governments/shareholders/swaps/trades"
	"github.com/deepvalue-network/software/libs/hash"
)

type content struct {
	hash      hash.Hash
	trade     trades.Trade
	createdOn time.Time
}

func createContent(
	hash hash.Hash,
	trade trades.Trade,
	createdOn time.Time,
) Content {
	out := content{
		hash:      hash,
		trade:     trade,
		createdOn: createdOn,
	}

	return &out
}

// Hash returns the hash
func (obj *content) Hash() hash.Hash {
	return obj.hash
}

// Trade returns the trade
func (obj *content) Trade() trades.Trade {
	return obj.trade
}

// CreatedOn returns the creation time
func (obj *content) CreatedOn() time.Time {
	return obj.createdOn
}
