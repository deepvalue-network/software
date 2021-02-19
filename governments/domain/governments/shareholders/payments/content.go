package payments

import (
	"time"

	"github.com/deepvalue-network/software/governments/domain/governments/shareholders"
	"github.com/deepvalue-network/software/libs/hash"
)

type content struct {
	hash      hash.Hash
	holder    shareholders.ShareHolder
	amount    uint
	createdOn time.Time
}

func createContent(
	hash hash.Hash,
	holder shareholders.ShareHolder,
	amount uint,
	createdOn time.Time,
) Content {
	out := content{
		hash:      hash,
		holder:    holder,
		amount:    amount,
		createdOn: createdOn,
	}

	return &out
}

// Hash returns the hash
func (obj *content) Hash() hash.Hash {
	return obj.hash
}

// ShareHolder returns the shareholder
func (obj *content) ShareHolder() shareholders.ShareHolder {
	return obj.holder
}

// Amount returns the amount
func (obj *content) Amount() uint {
	return obj.amount
}

// CreatedOn returns the creation time
func (obj *content) CreatedOn() time.Time {
	return obj.createdOn
}
