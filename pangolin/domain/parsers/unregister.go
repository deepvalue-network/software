package parsers

type unregister struct {
	variable string
}

func createUnregister(
	variable string,
) Unregister {
	out := unregister{
		variable: variable,
	}

	return &out
}

// Variable returns the variable
func (obj *unregister) Variable() string {
	return obj.variable
}
