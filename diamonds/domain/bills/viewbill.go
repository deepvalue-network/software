package bills

import "github.com/steve-care-software/products/libs/hash"

type viewBill struct {
	hash   hash.Hash
	bill   Bill
	seed   string
	amount uint64
}

func createViewBill(
	hash hash.Hash,
	bill Bill,
	seed string,
	amount uint64,
) ViewBill {
	out := viewBill{
		hash:   hash,
		bill:   bill,
		seed:   seed,
		amount: amount,
	}

	return &out
}

// Hash returns the hash
func (obj *viewBill) Hash() hash.Hash {
	return obj.hash
}

// Bill returns the bill
func (obj *viewBill) Bill() Bill {
	return obj.bill
}

// Seed returns the seed
func (obj *viewBill) Seed() string {
	return obj.seed
}

// Amount returns the amount
func (obj *viewBill) Amount() uint64 {
	return obj.amount
}
