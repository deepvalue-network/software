package governments

import (
	"time"

	"github.com/deepvalue-network/software/blockchain/domain/chains"
	"github.com/deepvalue-network/software/libs/hash"
)

type content struct {
	hash                     hash.Hash
	chain                    chains.Chain
	minPowerToPassResolution uint
	minPowerToPropose        uint
	canCancelVote            bool
	sharesVelocity           uint
	sharesCap                uint
	createdOn                time.Time
}

func createContent(
	hash hash.Hash,
	chain chains.Chain,
	minPowerToPassResolution uint,
	minPowerToPropose uint,
	canCancelVote bool,
	sharesVelocity uint,
	sharesCap uint,
	createdOn time.Time,
) Content {
	out := content{
		hash:                     hash,
		chain:                    chain,
		minPowerToPassResolution: minPowerToPassResolution,
		minPowerToPropose:        minPowerToPropose,
		canCancelVote:            canCancelVote,
		sharesVelocity:           sharesVelocity,
		sharesCap:                sharesCap,
		createdOn:                createdOn,
	}

	return &out
}

// Hash returns the hash
func (obj *content) Hash() hash.Hash {
	return obj.hash
}

// Chain returns the chain
func (obj *content) Chain() chains.Chain {
	return obj.chain
}

// MinPowerToPassResolution returns the minimum power to pass a resolution
func (obj *content) MinPowerToPassResolution() uint {
	return obj.minPowerToPassResolution
}

// MinPowerToPropose returns the minimum power to propose a resolution
func (obj *content) MinPowerToPropose() uint {
	return obj.minPowerToPropose
}

// CanCancelVote returns true if we can cancel a vote, false otherwise
func (obj *content) CanCancelVote() bool {
	return obj.canCancelVote
}

// SharesVelocity returns the shares velocity
func (obj *content) SharesVelocity() uint {
	return obj.sharesVelocity
}

// SharesCap returns the shares cap
func (obj *content) SharesCap() uint {
	return obj.sharesCap
}

// CreatedOn returns the creation time
func (obj *content) CreatedOn() time.Time {
	return obj.createdOn
}
