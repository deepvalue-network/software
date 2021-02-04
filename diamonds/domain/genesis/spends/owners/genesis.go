package owners

import (
	"github.com/steve-care-software/products/diamonds/domain/genesis/spends/views"
	"github.com/steve-care-software/products/diamonds/domain/owners"
	"github.com/steve-care-software/products/libs/hash"
)

type genesis struct {
	hash    hash.Hash
	owner   owners.Owner
	genesis views.Genesis
}

func createGenesis(
	hash hash.Hash,
	owner owners.Owner,
	gen views.Genesis,
) Genesis {
	out := genesis{
		hash:    hash,
		owner:   owner,
		genesis: gen,
	}

	return &out
}

// Hash returns the hash
func (obj *genesis) Hash() hash.Hash {
	return obj.hash
}

// Owner returns the owner
func (obj *genesis) Owner() owners.Owner {
	return obj.owner
}

// Genesis returns the genesis
func (obj *genesis) Genesis() views.Genesis {
	return obj.genesis
}
