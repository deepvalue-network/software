package parsers

type relational struct {
	lessThan StandardOperation
	equal    StandardOperation
	notEqual StandardOperation
}

func createRelationalWithLessThan(lessThan StandardOperation) Relational {
	return createRelationalInternally(lessThan, nil, nil)
}

func createRelationalWithEqual(equal StandardOperation) Relational {
	return createRelationalInternally(nil, equal, nil)
}

func createRelationalWithNotEqual(notEqual StandardOperation) Relational {
	return createRelationalInternally(nil, nil, notEqual)
}

func createRelationalInternally(lessThan StandardOperation, equal StandardOperation, notEqual StandardOperation) Relational {
	out := relational{
		lessThan: lessThan,
		equal:    equal,
		notEqual: notEqual,
	}

	return &out
}

// IsLessThan returns true if the relational operator is lessThan, false otherwise
func (obj *relational) IsLessThan() bool {
	return obj.lessThan != nil
}

// LessThan returns lessThan, if any
func (obj *relational) LessThan() StandardOperation {
	return obj.lessThan
}

// IsEqual returns true if the relational operator is equal, false otherwise
func (obj *relational) IsEqual() bool {
	return obj.equal != nil
}

// Equal returns equal, if any
func (obj *relational) Equal() StandardOperation {
	return obj.equal
}

// IsNotEqual returns true if the relational operator is not equal, false otherwise
func (obj *relational) IsNotEqual() bool {
	return obj.notEqual != nil
}

// NotEqual returns notEqual, if any
func (obj *relational) NotEqual() StandardOperation {
	return obj.notEqual
}
