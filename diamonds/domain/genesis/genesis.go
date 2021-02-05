package genesis

import (
	"time"

	"github.com/deepvalue-network/software/blockchain/domain/chains"
	"github.com/deepvalue-network/software/libs/hash"
)

type genesis struct {
	hash      hash.Hash
	amount    uint64
	chain     chains.Chain
	owner     []hash.Hash
	createdOn time.Time
	activeOn  time.Time
}

func createGenesis(
	hash hash.Hash,
	amount uint64,
	chain chains.Chain,
	owner []hash.Hash,
	createdOn time.Time,
	activeOn time.Time,
) Genesis {
	out := genesis{
		hash:      hash,
		amount:    amount,
		chain:     chain,
		owner:     owner,
		createdOn: createdOn,
		activeOn:  activeOn,
	}

	return &out
}

// Hash returns the hash
func (obj *genesis) Hash() hash.Hash {
	return obj.hash
}

// Amount returns the amount
func (obj *genesis) Amount() uint64 {
	return obj.amount
}

// Chain returns the chain
func (obj *genesis) Chain() chains.Chain {
	return obj.chain
}

// Owner returns the owner possible pubkey hashes
func (obj *genesis) Owner() []hash.Hash {
	return obj.owner
}

// CreatedOn returns the creation time
func (obj *genesis) CreatedOn() time.Time {
	return obj.createdOn
}

// ActiveOn returns the activeOn time
func (obj *genesis) ActiveOn() time.Time {
	return obj.activeOn
}
