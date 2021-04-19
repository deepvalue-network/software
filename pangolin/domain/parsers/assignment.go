package parsers

type assignment struct {
	variable string
	value    ValueRepresentation
}

func createAssignment(variable string, value ValueRepresentation) Assignment {
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
func (obj *assignment) Value() ValueRepresentation {
	return obj.value
}
