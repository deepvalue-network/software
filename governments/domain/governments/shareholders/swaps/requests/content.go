package requests

import (
	"time"

	"github.com/deepvalue-network/software/governments/domain/governments"
	"github.com/deepvalue-network/software/governments/domain/governments/shareholders/transfers/views"
	"github.com/deepvalue-network/software/libs/hash"
)

type content struct {
	hash      hash.Hash
	from      governments.Government
	stake     views.Section
	forGov    governments.Government
	to        []hash.Hash
	expiresOn time.Time
	createdOn time.Time
}

func createContent(
	hash hash.Hash,
	from governments.Government,
	stake views.Section,
	forGov governments.Government,
	to []hash.Hash,
	expiresOn time.Time,
	createdOn time.Time,
) Content {
	out := content{
		hash:      hash,
		from:      from,
		stake:     stake,
		forGov:    forGov,
		to:        to,
		expiresOn: expiresOn,
		createdOn: createdOn,
	}

	return &out
}

// Hash returns the hash
func (obj *content) Hash() hash.Hash {
	return obj.hash
}

// From returns the from government
func (obj *content) From() governments.Government {
	return obj.from
}

// Stake returns the stake
func (obj *content) Stake() views.Section {
	return obj.stake
}

// For returns the for government
func (obj *content) For() governments.Government {
	return obj.forGov
}

// To returns the to hashes pubkeys
func (obj *content) To() []hash.Hash {
	return obj.to
}

// ExpiresOn returns the expiration time
func (obj *content) ExpiresOn() time.Time {
	return obj.expiresOn
}

// CreatedOn returns the creation time
func (obj *content) CreatedOn() time.Time {
	return obj.createdOn
}
