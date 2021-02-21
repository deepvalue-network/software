package trades

import (
	"time"

	"github.com/deepvalue-network/software/governments/domain/governments/shareholders/swaps/requests"
	"github.com/deepvalue-network/software/governments/domain/governments/shareholders/transfers/views"
	"github.com/deepvalue-network/software/libs/hash"
)

type content struct {
	hash      hash.Hash
	request   requests.Request
	transfer  views.Transfer
	to        []hash.Hash
	expiresOn  time.Time
	createdOn time.Time
}

func createContent(
	hash hash.Hash,
	request requests.Request,
	transfer views.Transfer,
	to []hash.Hash,
	expiresOn time.Time,
	createdOn time.Time,
) Content {
	out := content{
		hash:      hash,
		request:   request,
		transfer:  transfer,
		to:        to,
		expiresOn:  expiresOn,
		createdOn: createdOn,
	}

	return &out
}

// Hash returns the hash
func (obj *content) Hash() hash.Hash {
	return obj.hash
}

// Request returns the request
func (obj *content) Request() requests.Request {
	return obj.request
}

// Transfer returns the transfer
func (obj *content) Transfer() views.Transfer {
	return obj.transfer
}

// To returns the to pubKey hashes
func (obj *content) To() []hash.Hash {
	return obj.to
}

// ExpiresOn returns the expiration time
func (obj *content) ExpiresOn() time.Time {
	return obj.expiresOn
}

// CreatedOn returns the creation time time
func (obj *content) CreatedOn() time.Time {
	return obj.createdOn
}
