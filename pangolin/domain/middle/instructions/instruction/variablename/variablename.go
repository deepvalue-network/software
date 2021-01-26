package variablename

type variableName struct {
	operation Operation
	variable  string
}

func createVariableName(operation Operation, variable string) VariableName {
	out := variableName{
		operation: operation,
		variable:  variable,
	}

	return &out
}

// Operation returns the operation
func (obj *variableName) Operation() Operation {
	return obj.operation
}

// Variable returns the variable
func (obj *variableName) Variable() string {
	return obj.variable
}
