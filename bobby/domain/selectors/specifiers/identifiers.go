package specifiers

import "github.com/deepvalue-network/software/libs/hash"

type identifiers struct {
	hash hash.Hash
	list []Identifier
}

func createIdentifiers(hash hash.Hash, list []Identifier) Identifiers {
	out := identifiers{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *identifiers) Hash() hash.Hash {
	return obj.hash
}

// All return the identifiers
func (obj *identifiers) All() []Identifier {
	return obj.list
}
