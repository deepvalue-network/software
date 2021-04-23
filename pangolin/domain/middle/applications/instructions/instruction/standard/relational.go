package standard

type relational struct {
	isLessThan bool
	isEqual    bool
	isNotEqual bool
}

func createRelationalWithLessThan() Relational {
	return createRelationalInternally(true, false, false)
}

func createRelationalWithEqual() Relational {
	return createRelationalInternally(false, true, false)
}

func createRelationalWithNotEqual() Relational {
	return createRelationalInternally(false, false, true)
}

func createRelationalInternally(isLessThan bool, isEqual bool, isNotEqual bool) Relational {
	out := relational{
		isLessThan: isLessThan,
		isEqual:    isEqual,
		isNotEqual: isNotEqual,
	}

	return &out
}

// IsLessThan returns true if less than, false otherwise
func (obj *relational) IsLessThan() bool {
	return obj.isLessThan
}

// IsEqual returns true if equal, false otherwise
func (obj *relational) IsEqual() bool {
	return obj.isEqual
}

// IsNotEqual returns true if not equal, false otherwise
func (obj *relational) IsNotEqual() bool {
	return obj.isNotEqual
}
