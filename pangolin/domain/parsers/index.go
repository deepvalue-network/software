package parsers

type index struct {
	variable string
}

func createIndex(
	variable string,
) Index {
	out := index{
		variable: variable,
	}

	return &out
}

// Variable returns the variable
func (obj *index) Variable() string {
	return obj.variable
}
