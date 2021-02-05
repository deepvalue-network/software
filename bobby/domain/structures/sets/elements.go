package sets

import "github.com/deepvalue-network/software/libs/hash"

type elements struct {
	ranked   RankedElements
	unranked UnrankedElements
}

func createElementsWithRanked(
	ranked RankedElements,
) Elements {
	return createElementsInternally(ranked, nil)
}

func createElementsWithUnranked(
	unranked UnrankedElements,
) Elements {
	return createElementsInternally(nil, unranked)
}

func createElementsInternally(
	ranked RankedElements,
	unranked UnrankedElements,
) Elements {
	out := elements{
		ranked:   ranked,
		unranked: unranked,
	}

	return &out
}

// Hash returns the hash
func (obj *elements) Hash() hash.Hash {
	if obj.IsUnranked() {
		return obj.UnRanked().Hash()
	}

	return obj.Ranked().Hash()
}

// IsUnique returns true if the elements are unique, false otherwise
func (obj *elements) IsUnique() bool {
	if obj.IsUnranked() {
		return obj.UnRanked().IsUnique()
	}

	return obj.Ranked().IsUnique()
}

// IsRanked returns true if the elements are ranked, false otherwise
func (obj *elements) IsRanked() bool {
	return obj.ranked != nil
}

// Ranked returns the ranked elements, if any
func (obj *elements) Ranked() RankedElements {
	return obj.ranked
}

// IsUnranked returns true if the elements are unranked, false otherwise
func (obj *elements) IsUnranked() bool {
	return obj.unranked != nil
}

// UnRanked returns the unranked elements, if any
func (obj *elements) UnRanked() UnrankedElements {
	return obj.unranked
}
