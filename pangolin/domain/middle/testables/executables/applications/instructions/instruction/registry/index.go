package registry

type index struct {
	intVal   int64
	variable string
}

func createIndexWithInt(intVal int64) Index {
	return createIndexInternally(intVal, "")
}

func createIndexWithVariable(variable string) Index {
	return createIndexInternally(-1, variable)
}

func createIndexInternally(
	intVal int64,
	variable string,
) Index {
	out := index{
		intVal:   intVal,
		variable: variable,
	}

	return &out
}

// IsInt returns true if there is an int, false otherwise
func (obj *index) IsInt() bool {
	return obj.intVal >= 0
}

// Int returns the int, if any
func (obj *index) Int() int64 {
	return obj.intVal
}

// IsVariable returns true if there is a variable, false otherwise
func (obj *index) IsVariable() bool {
	return obj.variable != ""
}

// Variable returns the variable, if any
func (obj *index) Variable() string {
	return obj.variable
}
