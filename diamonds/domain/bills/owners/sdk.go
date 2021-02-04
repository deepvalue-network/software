package owners

import (
	"github.com/steve-care-software/products/diamonds/domain/bills"
	"github.com/steve-care-software/products/diamonds/domain/owners"
	"github.com/steve-care-software/products/libs/hash"
)

// Bill represents an owned bill
type Bill interface {
	Hash() hash.Hash
	Owner() owners.Owner
	Bill() bills.ViewBill
}
