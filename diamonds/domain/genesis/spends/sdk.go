package spends

import (
	"time"

	"github.com/steve-care-software/products/diamonds/domain/genesis"
	"github.com/steve-care-software/products/libs/hash"
)

// Genesis represents a spent genesis
type Genesis interface {
	Hash() hash.Hash
	Amount() []byte
	Seed() []byte
	Genesis() genesis.Genesis
	CreatedOn() time.Time
}
