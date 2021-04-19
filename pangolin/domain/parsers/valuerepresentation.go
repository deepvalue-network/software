package parsers

type valueRepresentation struct {
	value    Value
	variable string
}

func createValueRepresentationWithValue(
	value Value,
) ValueRepresentation {
	return createValueRepresentationInternally(value, "")
}

func createValueRepresentationWithVariable(
	variable string,
) ValueRepresentation {
	return createValueRepresentationInternally(nil, variable)
}

func createValueRepresentationInternally(
	value Value,
	variable string,
) ValueRepresentation {
	out := valueRepresentation{
		value:    value,
		variable: variable,
	}

	return &out
}

// IsValue returns true if there is a value, false otherwise
func (obj *valueRepresentation) IsValue() bool {
	return obj.value != nil
}

// Value returns the value, if any
func (obj *valueRepresentation) Value() Value {
	return obj.value
}

// IsVariable returns true if there is a variable, false otherwise
func (obj *valueRepresentation) IsVariable() bool {
	return obj.variable != ""
}

// Variable returns the variable, if any
func (obj *valueRepresentation) Variable() string {
	return obj.variable
}
