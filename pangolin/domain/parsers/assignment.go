package parsers

type assignment struct {
	variable VariableName
	value    Value
}

func createAssignment(variable VariableName, value Value) Assignment {
	out := assignment{
		variable: variable,
		value:    value,
	}

	return &out
}

// Variable returns the variable
func (obj *assignment) Variable() VariableName {
	return obj.variable
}

// Value returns the value
func (obj *assignment) Value() Value {
	return obj.value
}
