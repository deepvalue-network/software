package owners

import (
	"github.com/steve-care-software/products/diamonds/domain/bills"
	"github.com/steve-care-software/products/diamonds/domain/owners"
	"github.com/steve-care-software/products/libs/hash"
)

type bill struct {
	hash  hash.Hash
	owner owners.Owner
	bill  bills.ViewBill
}

func createBill(
	hash hash.Hash,
	owner owners.Owner,
	viewBill bills.ViewBill,
) Bill {
	out := bill{
		hash:  hash,
		owner: owner,
		bill:  viewBill,
	}

	return &out
}

// Hash returns the hash
func (obj *bill) Hash() hash.Hash {
	return obj.hash
}

// Owner returns the owner
func (obj *bill) Owner() owners.Owner {
	return obj.owner
}

// Bill returns the bill
func (obj *bill) Bill() bills.ViewBill {
	return obj.bill
}
