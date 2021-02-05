package sets

import (
	"github.com/deepvalue-network/software/bobby/domain/resources"
	"github.com/deepvalue-network/software/libs/hash"
)

type unrankedElements struct {
	hash hash.Hash
	list []resources.Immutable
}

func createUnrankedElements(
	hash hash.Hash,
	list []resources.Immutable,
) UnrankedElements {
	out := unrankedElements{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *unrankedElements) Hash() hash.Hash {
	return obj.hash
}

// All returns the unranked elements
func (obj *unrankedElements) All() []resources.Immutable {
	return obj.list
}

// IsEmpty returns true if empty, false otherwise
func (obj *unrankedElements) IsEmpty() bool {
	return len(obj.list) <= 0
}

// IsUnique returns true if the elements are unique, false otherwise
func (obj *unrankedElements) IsUnique() bool {
	unique := map[string]string{}
	for _, oneElement := range obj.list {
		keyname := oneElement.Hash().String()
		unique[keyname] = ""
	}

	return len(unique) == len(obj.list)
}
