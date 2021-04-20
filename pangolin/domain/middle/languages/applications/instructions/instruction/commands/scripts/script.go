package scripts

type script struct {
	variable string
	values   []Value
}

func createScript(
	variable string,
	values []Value,
) Script {
	out := script{
		variable: variable,
		values:   values,
	}

	return &out
}

// Variable returns the variable
func (obj *script) Variable() string {
	return obj.variable
}

// Values returns the values
func (obj *script) Values() []Value {
	return obj.values
}
