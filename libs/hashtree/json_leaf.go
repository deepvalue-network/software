package hashtree

// JSONLeaf represents a json leaf
type JSONLeaf struct {
	Head string          `json:"head"`
	Node *JSONParentLeaf `json:"node"`
}

func createJSONLeafFromLeaf(leaf Leaf) *JSONLeaf {
	head := leaf.Head().String()
	var node *JSONParentLeaf
	if leaf.HasParent() {
		node = createJSONParentLeafFromParentLeaf(leaf.Parent())
	}

	return createJSONLeaf(head, node)
}

func createJSONLeaf(head string, node *JSONParentLeaf) *JSONLeaf {
	out := JSONLeaf{
		Head: head,
		Node: node,
	}

	return &out
}
