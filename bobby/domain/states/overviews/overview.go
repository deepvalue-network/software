package overviews

type overview struct {
	valid      []ValidTransaction
	invalid    []InvalidTransaction
	canBeSaved bool
}

func createOverview(
	valid []ValidTransaction,
	invalid []InvalidTransaction,
	canBeSaved bool,
) Overview {
	out := overview{
		valid:      valid,
		invalid:    invalid,
		canBeSaved: canBeSaved,
	}

	return &out
}

// Valid returns the valid transactions
func (obj *overview) Valid() []ValidTransaction {
	return obj.valid
}

// Invalid returns the invalid transactions
func (obj *overview) Invalid() []InvalidTransaction {
	return obj.invalid
}

// CanBeSaved returns true if the overview can be saved, false otherwise
func (obj *overview) CanBeSaved() bool {
	return obj.canBeSaved
}
