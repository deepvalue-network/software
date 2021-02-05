package specifiers

import "github.com/deepvalue-network/software/libs/hash"

type identifier struct {
	element  Element
	comparer Comparer
}

func createIdentifierWithElement(
	element Element,
) Identifier {
	return createIdentifierInternally(element, nil)
}

func createIdentifierWithComparer(
	comparer Comparer,
) Identifier {
	return createIdentifierInternally(nil, comparer)
}

func createIdentifierInternally(
	element Element,
	comparer Comparer,
) Identifier {
	out := identifier{
		element:  element,
		comparer: comparer,
	}

	return &out
}

// Hash returns the hash
func (obj *identifier) Hash() hash.Hash {
	if obj.IsElement() {
		return obj.Element().Hash()
	}

	return obj.Comparer().Hash()
}

// IsElement returns true if there is an element, false otherwise
func (obj *identifier) IsElement() bool {
	return obj.element != nil
}

// Element returns the element, if any
func (obj *identifier) Element() Element {
	return obj.element
}

// IsComparer returns true if there is a comparer, false otherwise
func (obj *identifier) IsComparer() bool {
	return obj.comparer != nil
}

// Comparer returns the comparer, if any
func (obj *identifier) Comparer() Comparer {
	return obj.comparer
}
