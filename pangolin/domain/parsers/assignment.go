package parsers

type assignment struct {
	variable string
	value    Value
}

func createAssignment(variable string, value Value) Assignment {
	out := assignment{
		variable: variable,
		value:    value,
	}

	return &out
}

// Variable returns the variable
func (obj *assignment) Variable() string {
	return obj.variable
}

// Value returns the value
func (obj *assignment) Value() Value {
	return obj.value
}
