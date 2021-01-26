package hashtree

// JSONHashTree represents a json hashtree
type JSONHashTree struct {
	Head string          `json:"head"`
	Node *JSONParentLeaf `json:"node"`
}

func createJSONHashTreeFromHashTree(hashTree HashTree) *JSONHashTree {
	head := hashTree.Head().String()
	node := createJSONParentLeafFromParentLeaf(hashTree.Parent())
	return createJSONHashTree(head, node)
}

func createJSONHashTree(head string, node *JSONParentLeaf) *JSONHashTree {
	out := JSONHashTree{
		Head: head,
		Node: node,
	}

	return &out
}
