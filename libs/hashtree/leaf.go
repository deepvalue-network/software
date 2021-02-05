package hashtree

import (
	"bytes"

	"github.com/deepvalue-network/software/libs/hash"
)

type leaf struct {
	Hd hash.Hash   `json:"head"`
	Pt *parentLeaf `json:"parent"`
}

func createLeafFromJSON(ins *JSONLeaf) (*leaf, error) {
	head, err := hash.NewAdapter().FromString(ins.Head)
	if err != nil {
		return nil, err
	}

	if ins.Node == nil {
		return createLeaf(*head), nil
	}

	pt, err := createParentLeafFromJSON(ins.Node)
	if err != nil {
		return nil, err
	}

	return createLeafWithParent(*head, pt), nil
}

func createLeaf(head hash.Hash) *leaf {
	out := leaf{
		Hd: head,
		Pt: nil,
	}

	return &out
}

func createLeafWithParent(head hash.Hash, parent *parentLeaf) *leaf {
	out := leaf{
		Hd: head,
		Pt: parent,
	}

	return &out
}

func createChildLeaf(left *leaf, right *leaf) (*leaf, error) {
	data := bytes.Join([][]byte{
		left.Head().Bytes(),
		right.Head().Bytes(),
	}, []byte{})

	h, err := hash.NewAdapter().FromBytes(data)
	if err != nil {
		return nil, err
	}

	out := createLeaf(*h)
	return out, nil
}

// Head returns the head hash
func (obj *leaf) Head() hash.Hash {
	return obj.Hd
}

// HasParent returns true if there is a parent, false otherwise
func (obj *leaf) HasParent() bool {
	return obj.Pt != nil
}

// Parent returns the parent, if any
func (obj *leaf) Parent() ParentLeaf {
	return obj.Pt
}

// Leaves returns the leaves
func (obj *leaf) Leaves() Leaves {
	if obj.HasParent() {
		return obj.Parent().BlockLeaves()
	}

	leaves := []*leaf{
		obj,
	}

	output := createLeaves(leaves)
	return output
}

// Height returns the leaf height
func (obj *leaf) Height() int {
	cpt := 0
	var oneLeaf Leaf
	for {

		if oneLeaf == nil {
			oneLeaf = obj
		}

		if !oneLeaf.HasParent() {
			return cpt
		}

		cpt++
		oneLeaf = oneLeaf.Parent().Left()
	}
}
