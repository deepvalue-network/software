package hashtree

import (
	"encoding/json"
	"math"

	"github.com/deepvalue-network/software/libs/hash"
)

type block struct {
	List []hash.Hash `json:"hashes"`
}

func createBlockFromData(data [][]byte) (Block, error) {

	if len(data) <= 1 {
		data = append(data, []byte(""))
	}

	hashes := []hash.Hash{}
	for _, oneData := range data {
		oneHash, err := hash.NewAdapter().FromBytes(oneData)
		if err != nil {
			return nil, err
		}

		hashes = append(hashes, *oneHash)
	}

	blk := block{
		List: hashes,
	}

	return blk.resize(), nil
}

func (obj *block) resize() Block {
	//need to make sure the elements are always a power of 2:
	isPowerOfTwo := obj.isLengthPowerForTwo()
	if !isPowerOfTwo {
		obj.resizeToNextPowerOfTwo()
	}

	return obj
}

func (obj *block) isLengthPowerForTwo() bool {
	length := len(obj.List)
	return (length != 0) && ((length & (length - 1)) == 0)
}

func (obj *block) resizeToNextPowerOfTwo() (Block, error) {
	lengthAsFloat := float64(len(obj.List))
	next := uint(math.Pow(2, math.Ceil(math.Log(lengthAsFloat)/math.Log(2))))
	remaining := int(next) - int(lengthAsFloat)
	for i := 0; i < remaining; i++ {
		single, err := hash.NewAdapter().FromBytes(nil)
		if err != nil {
			return nil, err
		}

		obj.List = append(obj.List, *single)
	}

	return obj, nil
}

// Leaves returns the leaves of the block
func (obj *block) Leaves() Leaves {
	leaves := []*leaf{}
	for _, oneBlockHash := range obj.List {
		oneLeaf := createLeaf(oneBlockHash)
		leaves = append(leaves, oneLeaf)
	}

	return createLeaves(leaves)
}

// HashTree returns the HashTree
func (obj *block) HashTree() (HashTree, error) {
	leaves := obj.Leaves()
	return leaves.HashTree()
}

// MarshalJSON converts the instance to JSON
func (obj *block) MarshalJSON() ([]byte, error) {
	list := []string{}
	for _, oneHash := range obj.List {
		list = append(list, oneHash.String())
	}

	return json.Marshal(list)
}

// UnmarshalJSON converts the JSON to an instance
func (obj *block) UnmarshalJSON(data []byte) error {
	list := new([]string)
	err := json.Unmarshal(data, list)
	if err != nil {
		return err
	}

	hashes := []hash.Hash{}
	for _, oneHashStr := range *list {
		hsh, err := hash.NewAdapter().FromString(oneHashStr)
		if err != nil {
			return err
		}

		hashes = append(hashes, *hsh)
	}

	obj.List = hashes
	return nil
}
