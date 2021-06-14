package tokens

type cardinality struct {
	isNonZeroMultiple bool
	isZeroMultiple    bool
	specific          SpecificCardinality
}

func createCardinalityWithNonZeroMultiple() Cardinality {
	return createCardinalityInternally(true, false, nil)
}

func createCardinalityWithZeroMultiple() Cardinality {
	return createCardinalityInternally(false, true, nil)
}

func createCardinalityWithSpecific(specific SpecificCardinality) Cardinality {
	return createCardinalityInternally(false, false, specific)
}

func createCardinalityInternally(
	isNonZeroMultiple bool,
	isZeroMultiple bool,
	specific SpecificCardinality,
) Cardinality {
	out := cardinality{
		isNonZeroMultiple: isNonZeroMultiple,
		isZeroMultiple:    isZeroMultiple,
		specific:          specific,
	}

	return &out
}

// IsNonZeroMultiple returns true if there is a non-zero multiple, false otherwise
func (obj *cardinality) IsNonZeroMultiple() bool {
	return obj.isNonZeroMultiple
}

// IsNonZeroMultiple returns true if there is a zero multiple, false otherwise
func (obj *cardinality) IsZeroMultiple() bool {
	return obj.isZeroMultiple
}

// IsSpecific returns true if there is a specific cardinality, false otherwise
func (obj *cardinality) IsSpecific() bool {
	return obj.specific != nil
}

// Specific returns the specific cardinality, if any
func (obj *cardinality) Specific() SpecificCardinality {
	return obj.specific
}
