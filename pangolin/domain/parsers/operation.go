package parsers

type operation struct {
	arythmetic Arythmetic
	relational Relational
	logical    Logical
}

func createOperationWithArythmetic(arythmetic Arythmetic) Operation {
	return createOperationInternally(arythmetic, nil, nil)
}

func createOperationWithRelational(relational Relational) Operation {
	return createOperationInternally(nil, relational, nil)
}

func createOperationWithLogical(logical Logical) Operation {
	return createOperationInternally(nil, nil, logical)
}

func createOperationInternally(arythmetic Arythmetic, relational Relational, logical Logical) Operation {
	out := operation{
		arythmetic: arythmetic,
		relational: relational,
		logical:    logical,
	}

	return &out
}

// IsArythmetic returns true if the operation is arythmetic, false otherwise
func (obj *operation) IsArythmetic() bool {
	return obj.arythmetic != nil
}

// Arythmetic returns the arythmetic operator, if any
func (obj *operation) Arythmetic() Arythmetic {
	return obj.arythmetic
}

// IsRelational returns true if the operation is relational, false otherwise
func (obj *operation) IsRelational() bool {
	return obj.relational != nil
}

// Relational returns the relational operator, if any
func (obj *operation) Relational() Relational {
	return obj.relational
}

// IsLogical returns true if the operation is logical, false otherwise
func (obj *operation) IsLogical() bool {
	return obj.logical != nil
}

// Logical returns the logical operator, if any
func (obj *operation) Logical() Logical {
	return obj.logical
}
