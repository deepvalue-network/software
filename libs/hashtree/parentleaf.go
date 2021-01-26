package hashtree

import (
	"bytes"

	"github.com/steve-care-software/products/libs/hash"
)

type parentLeaf struct {
	Lft *leaf `json:"left"`
	Rgt *leaf `json:"right"`
}

func createParentLeafFromJSON(ins *JSONParentLeaf) (*parentLeaf, error) {
	left, err := createLeafFromJSON(ins.Left)
	if err != nil {
		return nil, err
	}

	right, err := createLeafFromJSON(ins.Right)
	if err != nil {
		return nil, err
	}

	return createParentLeaf(left, right), nil
}

func createParentLeaf(left *leaf, right *leaf) *parentLeaf {
	out := parentLeaf{
		Lft: left,
		Rgt: right,
	}

	return &out
}

// HashTree returns the hashtree
func (obj *parentLeaf) HashTree() (HashTree, error) {
	data := bytes.Join([][]byte{
		obj.Left().Head().Bytes(),
		obj.Right().Head().Bytes(),
	}, []byte{})

	hsh, err := hash.NewAdapter().FromBytes(data)
	if err != nil {
		return nil, err
	}

	out := createHashTree(*hsh, obj)
	return out, nil
}

// BlockLeaves returns the block leaves
func (obj *parentLeaf) BlockLeaves() Leaves {
	left := obj.Left()
	right := obj.Right()
	leftLeaves := left.Leaves()
	rightLeaves := right.Leaves()
	return leftLeaves.Merge(rightLeaves)
}

// Left returns the left leaf
func (obj *parentLeaf) Left() Leaf {
	return obj.Lft
}

// Right returns the right leaf
func (obj *parentLeaf) Right() Leaf {
	return obj.Rgt
}
