package standard

type operation struct {
	ary  Arythmetic
	rel  Relational
	log  Logical
	misc Misc
}

func createOperationWithArythmetic(ary Arythmetic) Operation {
	return createOperationInternally(ary, nil, nil, nil)
}

func createOperationWithRelational(rel Relational) Operation {
	return createOperationInternally(nil, rel, nil, nil)
}

func createOperationWithLogical(log Logical) Operation {
	return createOperationInternally(nil, nil, log, nil)
}

func createOperationWithMisc(misc Misc) Operation {
	return createOperationInternally(nil, nil, nil, misc)
}

func createOperationInternally(
	ary Arythmetic,
	rel Relational,
	log Logical,
	misc Misc,
) Operation {
	out := operation{
		ary:  ary,
		rel:  rel,
		log:  log,
		misc: misc,
	}

	return &out
}

// IsArythmetic returns true if the operation is arythmetic, false otherwise
func (obj *operation) IsArythmetic() bool {
	return obj.ary != nil
}

// Arythmetic returns the arythmetic, if any
func (obj *operation) Arythmetic() Arythmetic {
	return obj.ary
}

// IsRelational returns true if the operation is relational, false otherwise
func (obj *operation) IsRelational() bool {
	return obj.rel != nil
}

// Relational returns the relational, if any
func (obj *operation) Relational() Relational {
	return obj.rel
}

// IsLogical returns true if the operation is logical, false otherwise
func (obj *operation) IsLogical() bool {
	return obj.log != nil
}

// Logical returns the logical, if any
func (obj *operation) Logical() Logical {
	return obj.log
}

// IsMisc returns true if the operation is misc, false otherwise
func (obj *operation) IsMisc() bool {
	return obj.misc != nil
}

// Misc returns the misc, if any
func (obj *operation) Misc() Misc {
	return obj.misc
}
