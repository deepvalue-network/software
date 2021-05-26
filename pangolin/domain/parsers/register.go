package parsers

type register struct {
	variable string
	index    IntPointer
}

func createRegister(
	variable string,
) Register {
	return createRegisterInternally(variable, nil)
}

func createRegisterWithIndex(
	variable string,
	index IntPointer,
) Register {
	return createRegisterInternally(variable, index)
}

func createRegisterInternally(
	variable string,
	index IntPointer,
) Register {
	out := register{
		variable: variable,
		index:    index,
	}

	return &out
}

// Variable adds a variable to the builder
func (obj *register) Variable() string {
	return obj.variable
}

// HasIndex returns true if there is an index, false otherwise
func (obj *register) HasIndex() bool {
	return obj.index != nil
}

// Index returns the index, if any
func (obj *register) Index() IntPointer {
	return obj.index
}
