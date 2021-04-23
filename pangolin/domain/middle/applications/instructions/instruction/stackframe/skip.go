package stackframe

type skip struct {
	intVal   int64
	variable string
}

func createSkipWithInt(intVal int64) Skip {
	return createSkipInternally(intVal, "")
}

func createSkipWithVariable(variable string) Skip {
	return createSkipInternally(-1, variable)
}

func createSkipInternally(
	intVal int64,
	variable string,
) Skip {
	out := skip{
		intVal:   intVal,
		variable: variable,
	}

	return &out
}

// IsInt returns true if there is an int, false otherwise
func (obj *skip) IsInt() bool {
	return obj.intVal >= 0
}

// Int returns the int, if any
func (obj *skip) Int() int64 {
	return obj.intVal
}

// IsVariable returns true if there is a variable, false otherwise
func (obj *skip) IsVariable() bool {
	return obj.variable != ""
}

// Variable returns the variable, if any
func (obj *skip) Variable() string {
	return obj.variable
}
