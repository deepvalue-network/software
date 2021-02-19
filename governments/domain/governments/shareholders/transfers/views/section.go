package views

import (
	"github.com/deepvalue-network/software/governments/domain/governments/shareholders/transfers"
	"github.com/deepvalue-network/software/libs/hash"
)

type section struct {
	hash     hash.Hash
	transfer transfers.Transfer
	seed     string
	amount   uint
}

func createSection(
	hash hash.Hash,
	transfer transfers.Transfer,
	seed string,
	amount uint,
) Section {
	out := section{
		hash:     hash,
		transfer: transfer,
		seed:     seed,
		amount:   amount,
	}

	return &out
}

// Hash returns the hash
func (obj *section) Hash() hash.Hash {
	return obj.hash
}

// Transfer returns the transfer
func (obj *section) Transfer() transfers.Transfer {
	return obj.transfer
}

// Seed returns the seed
func (obj *section) Seed() string {
	return obj.seed
}

// Amount returns the amount
func (obj *section) Amount() uint {
	return obj.amount
}
