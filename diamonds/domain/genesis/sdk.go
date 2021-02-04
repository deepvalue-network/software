package genesis

import (
	"time"

	"github.com/steve-care-software/products/libs/hash"
)

// Genesis represents a genesis diamond
type Genesis interface {
	Hash() hash.Hash
	Owner() []hash.Hash
	CreatedOn() time.Time
	ActiveOn() time.Time
}
