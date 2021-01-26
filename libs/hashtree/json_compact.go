package hashtree

// JSONCompact represents a json compact
type JSONCompact struct {
	Head   string      `json:"head"`
	Leaves []*JSONLeaf `json:"leaves"`
}

func createJSONCompactFromCompact(hashTree Compact) *JSONCompact {
	head := hashTree.Head().String()

	list := []*JSONLeaf{}
	leaves := hashTree.Leaves().Leaves()
	for _, oneLeaf := range leaves {
		list = append(list, createJSONLeafFromLeaf(oneLeaf))
	}

	return createJSONCompact(head, list)
}

func createJSONCompact(head string, leaves []*JSONLeaf) *JSONCompact {
	out := JSONCompact{
		Head:   head,
		Leaves: leaves,
	}

	return &out
}
