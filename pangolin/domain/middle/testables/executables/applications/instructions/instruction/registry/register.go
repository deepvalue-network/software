package registry

type register struct {
	variable string
	index    Index
}

func createRegister(
	variable string,
) Register {
	return createRegisterInternally(variable, nil)
}

func createRegisterWithIndex(
	variable string,
	index Index,
) Register {
	return createRegisterInternally(variable, index)
}

func createRegisterInternally(
	variable string,
	index Index,
) Register {
	out := register{
		variable: variable,
		index:    index,
	}

	return &out
}

// Variable returns the variable
func (obj *register) Variable() string {
	return obj.variable
}

// HasIndex returns true if there is an index, false otherwise
func (obj *register) HasIndex() bool {
	return obj.index != nil
}

// Index returns the index, if any
func (obj *register) Index() Index {
	return obj.index
}
