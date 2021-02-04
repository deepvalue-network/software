package views

import (
	"github.com/steve-care-software/products/diamonds/domain/genesis/spends"
	"github.com/steve-care-software/products/libs/hash"
)

// Genesis represents a genesis view spent
type Genesis interface {
	Hash() hash.Hash
	Genesis() spends.Genesis
	Seed() string
	Amount() uint64
}
