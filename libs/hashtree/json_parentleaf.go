package hashtree

// JSONParentLeaf represents a json parent leaf
type JSONParentLeaf struct {
	Left  *JSONLeaf `json:"left"`
	Right *JSONLeaf `json:"right"`
}

func createJSONParentLeafFromParentLeaf(parent ParentLeaf) *JSONParentLeaf {
	left := createJSONLeafFromLeaf(parent.Left())
	right := createJSONLeafFromLeaf(parent.Right())
	return createJSONParentLeaf(left, right)
}

func createJSONParentLeaf(left *JSONLeaf, right *JSONLeaf) *JSONParentLeaf {
	out := JSONParentLeaf{
		Left:  left,
		Right: right,
	}

	return &out
}
