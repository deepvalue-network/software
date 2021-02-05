package sets

import (
	"github.com/deepvalue-network/software/bobby/domain/resources"
	"github.com/deepvalue-network/software/libs/hash"
)

type rankedElements struct {
	hash hash.Hash
	mp   map[uint]resources.Immutable
}

func createRankedElements(
	hash hash.Hash,
	mp map[uint]resources.Immutable,
) RankedElements {
	out := rankedElements{
		hash: hash,
		mp:   mp,
	}

	return &out
}

// Hash returns the hash
func (obj *rankedElements) Hash() hash.Hash {
	return obj.hash
}

// All returns the ranked elements
func (obj *rankedElements) All() map[uint]resources.Immutable {
	return obj.mp
}

// IsEmpty returns true if empty, false otherwise
func (obj *rankedElements) IsEmpty() bool {
	return len(obj.mp) <= 0
}

// IsUnique returns true if the elements are unique, false otherwise
func (obj *rankedElements) IsUnique() bool {
	unique := map[string]string{}
	for _, oneElement := range obj.mp {
		keyname := oneElement.Hash().String()
		unique[keyname] = ""
	}

	return len(unique) == len(obj.mp)
}
