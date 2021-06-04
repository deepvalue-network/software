package parsers

type swtch struct {
	variable string
}

func createSwitch(
	variable string,
) Switch {
	out := swtch{
		variable: variable,
	}

	return &out
}

// Variable returns the variable
func (obj *swtch) Variable() string {
	return obj.variable
}
