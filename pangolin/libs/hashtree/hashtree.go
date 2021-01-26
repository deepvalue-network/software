package hashtree

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/steve-care-software/products/pangolin/libs/hash"
)

// HTree represents a concrete HashTree implementation
type hashtree struct {
	Hd hash.Hash   `json:"head"`
	Pt *parentLeaf `json:"parent"`
}

func createHashTreeFromJS(js []byte) (*hashtree, error) {
	ins := new(JSONHashTree)
	err := json.Unmarshal(js, ins)
	if err != nil {
		return nil, err
	}

	return createHashTreeFromJSON(ins)
}

func createHashTreeFromJSON(ins *JSONHashTree) (*hashtree, error) {
	head, err := hash.NewAdapter().FromString(ins.Head)
	if err != nil {
		return nil, err
	}

	pt, err := createParentLeafFromJSON(ins.Node)
	if err != nil {
		return nil, err
	}

	return createHashTree(*head, pt), nil
}

func createHashTree(head hash.Hash, parent *parentLeaf) *hashtree {
	out := hashtree{
		Hd: head,
		Pt: parent,
	}

	return &out
}

func createHashTreeFromBlocks(blocks [][]byte) (*hashtree, error) {
	blockHashes, blockHashesErr := createBlockFromData(blocks)
	if blockHashesErr != nil {
		return nil, blockHashesErr
	}

	tree, err := blockHashes.HashTree()
	return tree.(*hashtree), err
}

// Height returns the hashtree height
func (obj *hashtree) Height() int {
	left := obj.Pt.Left()
	return left.Height() + 2
}

// Length returns the hashtree length
func (obj *hashtree) Length() int {
	blockLeaves := obj.Pt.BlockLeaves()
	return len(blockLeaves.Leaves())
}

// Head returns the head hash
func (obj *hashtree) Head() hash.Hash {
	return obj.Hd
}

// Parent returns the parent leaf
func (obj *hashtree) Parent() ParentLeaf {
	return obj.Pt
}

// Compact returns the compact version of the hashtree
func (obj *hashtree) Compact() Compact {
	blockLeaves := obj.Pt.BlockLeaves()
	return createCompact(obj.Hd, blockLeaves)
}

// Order orders data that matches the leafs of the HashTree
func (obj *hashtree) Order(data [][]byte) ([][]byte, error) {
	hashAdapter := hash.NewAdapter()
	hashed := map[string][]byte{}
	for _, oneData := range data {
		hsh, err := hashAdapter.FromBytes(oneData)
		if err != nil {
			return nil, err
		}

		hashAsString := hsh.String()
		hashed[hashAsString] = oneData
	}

	out := [][]byte{}
	leaves := obj.Pt.BlockLeaves().Leaves()
	for _, oneLeaf := range leaves {
		leafHashAsString := oneLeaf.Head().String()
		if oneData, ok := hashed[leafHashAsString]; ok {
			out = append(out, oneData)
			continue
		}

		//must be a filling Leaf, so continue:
		continue
	}

	if len(out) != len(data) {
		str := fmt.Sprintf("the length of the input data (%d) does not match the length of the output (%d), therefore, some data blocks could not be found in the hash leaves", len(data), len(out))
		return nil, errors.New(str)
	}

	return out, nil
}

// MarshalJSON converts the instance to JSON
func (obj *hashtree) MarshalJSON() ([]byte, error) {
	ins := createJSONHashTreeFromHashTree(obj)
	return json.Marshal(ins)
}

// UnmarshalJSON converts the JSON to an instance
func (obj *hashtree) UnmarshalJSON(data []byte) error {
	ht, err := createHashTreeFromJS(data)
	if err != nil {
		return err
	}

	obj.Hd = ht.Hd
	obj.Pt = ht.Pt
	return nil
}
