package heads

type head struct {
	variable string
	values   []Value
}

func createHead(
	variable string,
	values []Value,
) Head {
	out := head{
		variable: variable,
		values:   values,
	}

	return &out
}

// Variable returns the variable
func (obj *head) Variable() string {
	return obj.variable
}

// Values returns the values
func (obj *head) Values() []Value {
	return obj.values
}
