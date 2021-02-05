package contents

import (
	"github.com/deepvalue-network/software/bobby/domain/selectors"
	"github.com/deepvalue-network/software/bobby/domain/structures/sets"
)

type set struct {
	selector selectors.Selector
	elements sets.Elements
}

func createSet(
	selector selectors.Selector,
	elements sets.Elements,
) Set {
	out := set{
		selector: selector,
		elements: elements,
	}

	return &out
}

// Set returns the set selector
func (obj *set) Set() selectors.Selector {
	return obj.selector
}

// Elements returns the set elements
func (obj *set) Elements() sets.Elements {
	return obj.elements
}
