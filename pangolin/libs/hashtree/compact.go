package hashtree

import (
	"encoding/json"

	"github.com/steve-care-software/products/pangolin/libs/hash"
)

type compact struct {
	hash   hash.Hash
	leaves Leaves
}

func createCompactFromJSON(ins *JSONCompact) (Compact, error) {
	head, err := hash.NewAdapter().FromString(ins.Head)
	if err != nil {
		return nil, err
	}

	list := []*leaf{}
	for _, oneleafJSON := range ins.Leaves {
		leaf, err := createLeafFromJSON(oneleafJSON)
		if err != nil {
			return nil, err
		}

		list = append(list, leaf)
	}

	leaves := createLeaves(list)
	return createCompact(*head, leaves), nil
}

func createCompact(head hash.Hash, leaves Leaves) Compact {
	out := compact{
		hash:   head,
		leaves: leaves,
	}

	return &out
}

// Head returns the head hash
func (obj *compact) Head() hash.Hash {
	return obj.hash
}

// Leaves returns the leaves
func (obj *compact) Leaves() Leaves {
	return obj.leaves
}

// Length returns the length of the compact hashtree
func (obj *compact) Length() int {
	return len(obj.leaves.Leaves())
}

// MarshalJSON converts the instance to JSON
func (obj *compact) MarshalJSON() ([]byte, error) {
	ins := createJSONCompactFromCompact(obj)
	return json.Marshal(ins)
}

// UnmarshalJSON converts the JSON to an instance
func (obj *compact) UnmarshalJSON(data []byte) error {
	ins := new(JSONCompact)
	err := json.Unmarshal(data, ins)
	if err != nil {
		return err
	}

	ht, err := createCompactFromJSON(ins)
	if err != nil {
		return err
	}

	obj.hash = ht.Head()
	obj.leaves = ht.Leaves()
	return nil
}
