package parsers

type intPointer struct {
	intVal   int64
	variable string
}

func createIntPointerWithInt(intVal int64) IntPointer {
	return createIntPointerInternally(intVal, "")
}

func createIntPointerWithVariable(variable string) IntPointer {
	return createIntPointerInternally(-1, variable)
}

func createIntPointerInternally(
	intVal int64,
	variable string,
) IntPointer {
	out := intPointer{
		intVal:   intVal,
		variable: variable,
	}

	return &out
}

// IsInt returns true if there is an int, false otherwise
func (obj *intPointer) IsInt() bool {
	return obj.intVal >= 0
}

// Int returns the int, if any
func (obj *intPointer) Int() int64 {
	return obj.intVal
}

// IsVariable returns true if there is a variakble, false otherwise
func (obj *intPointer) IsVariable() bool {
	return obj.variable != ""
}

// Variable returns the variable, if any
func (obj *intPointer) Variable() string {
	return obj.variable
}
