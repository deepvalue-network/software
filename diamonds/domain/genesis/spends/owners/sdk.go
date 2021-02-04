package owners

import (
	"github.com/steve-care-software/products/diamonds/domain/genesis/spends/views"
	"github.com/steve-care-software/products/diamonds/domain/owners"
	"github.com/steve-care-software/products/libs/hash"
)

// Spent represents an owned genesis spent
type Spent interface {
	Hash() hash.Hash
	Owner() owners.Owner
	Genesis() views.Genesis
}
