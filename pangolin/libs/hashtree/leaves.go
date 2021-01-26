package hashtree

type leaves struct {
	Lst []*leaf `json:"leaves"`
}

func createLeaves(list []*leaf) *leaves {
	out := leaves{
		Lst: list,
	}

	return &out
}

// Leaves returns the leaves
func (obj *leaves) Leaves() []Leaf {
	out := []Leaf{}
	for _, oneLeaf := range obj.Lst {
		out = append(out, oneLeaf)
	}

	return out
}

// Merge merge Leaves instances
func (obj *leaves) Merge(lves Leaves) Leaves {
	for _, oneLeaf := range lves.Leaves() {
		obj.Lst = append(obj.Lst, oneLeaf.(*leaf))
	}

	return obj
}

// HashTree returns the hashtree
func (obj *leaves) HashTree() (HashTree, error) {
	length := len(obj.Lst)
	if length == 2 {
		left := obj.Lst[0]
		right := obj.Lst[1]
		parent := createParentLeaf(left, right)
		return parent.HashTree()
	}

	childrenLeaves, err := obj.createChildrenLeaves()
	if err != nil {
		return nil, err
	}

	return childrenLeaves.HashTree()
}

func (obj *leaves) createChildrenLeaves() (Leaves, error) {
	childrenLeaves := []*leaf{}
	for index, oneLeaf := range obj.Lst {

		if index%2 != 0 {
			continue
		}

		left := oneLeaf
		right := obj.Lst[index+1]
		child, err := createChildLeaf(left, right)
		if err != nil {
			return nil, err
		}

		parent := createParentLeaf(left, right)
		childWithParent := createLeafWithParent(child.Head(), parent)
		childrenLeaves = append(childrenLeaves, childWithParent)
	}

	return createLeaves(childrenLeaves), nil
}
