package tokens

type specificCardinality struct {
	amount *uint
	rnge   Range
}

func createSpecificCardinalityWithAmount(
	amount *uint,
) SpecificCardinality {
	return createSpecificCardinalityInternally(amount, nil)
}

func createSpecificCardinalityWithRange(
	rnge Range,
) SpecificCardinality {
	return createSpecificCardinalityInternally(nil, rnge)
}

func createSpecificCardinalityInternally(
	amount *uint,
	rnge Range,
) SpecificCardinality {
	out := specificCardinality{
		amount: amount,
		rnge:   rnge,
	}

	return &out
}

// IsAmount returns true if there is an amount, false otherwise
func (obj *specificCardinality) IsAmount() bool {
	return obj.amount != nil
}

// Amount returns the amount, if any
func (obj *specificCardinality) Amount() *uint {
	return obj.amount
}

// IsRange returns true if there is a range, false otherwise
func (obj *specificCardinality) IsRange() bool {
	return obj.rnge != nil
}

// Range returns the range, if any
func (obj *specificCardinality) Range() Range {
	return obj.rnge
}
