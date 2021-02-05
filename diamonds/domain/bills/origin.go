package bills

import (
	"github.com/deepvalue-network/software/diamonds/domain/genesis/spends/views"
	"github.com/deepvalue-network/software/libs/hash"
)

type origin struct {
	gen  views.Genesis
	bill ViewBill
}

func createOriginWithGenesis(
	gen views.Genesis,
) Origin {
	return createOriginInternally(gen, nil)
}

func createOriginWithBill(
	bill ViewBill,
) Origin {
	return createOriginInternally(nil, bill)
}

func createOriginInternally(
	gen views.Genesis,
	bill ViewBill,
) Origin {
	out := origin{
		gen:  gen,
		bill: bill,
	}

	return &out
}

// Hash returns the hash
func (obj *origin) Hash() hash.Hash {
	if obj.IsGenesis() {
		return obj.Genesis().Hash()
	}

	return obj.Bill().Hash()
}

// IsGenesis returns true if there is a genesis, false otherwise
func (obj *origin) IsGenesis() bool {
	return obj.gen != nil
}

// Genesis returns the genesis, if any
func (obj *origin) Genesis() views.Genesis {
	return obj.gen
}

// IsBill returns true if there is a bill, false otherwise
func (obj *origin) IsBill() bool {
	return obj.bill != nil
}

// Bill returns the bill, if any
func (obj *origin) Bill() ViewBill {
	return obj.bill
}
