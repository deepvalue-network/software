package specifiers

import "github.com/deepvalue-network/software/libs/hash"

type comparer struct {
	hash   hash.Hash
	first  Identifier
	second Identifier
	isAnd  bool
}

func createComparerWithAnd(
	hash hash.Hash,
	first Identifier,
	second Identifier,
) Comparer {
	return createComparerInternally(hash, first, second, true)
}

func createComparerWithOr(
	hash hash.Hash,
	first Identifier,
	second Identifier,
) Comparer {
	return createComparerInternally(hash, first, second, false)
}

func createComparerInternally(
	hash hash.Hash,
	first Identifier,
	second Identifier,
	isAnd bool,
) Comparer {
	out := comparer{
		hash:   hash,
		first:  first,
		second: second,
		isAnd:  isAnd,
	}

	return &out
}

// Hash returns the hash
func (obj *comparer) Hash() hash.Hash {
	return obj.hash
}

// First returns the first identifier
func (obj *comparer) First() Identifier {
	return obj.first
}

// Second returns the second identifier
func (obj *comparer) Second() Identifier {
	return obj.second
}

// IsAnd returns true if and, false otherwise
func (obj *comparer) IsAnd() bool {
	return obj.isAnd
}
