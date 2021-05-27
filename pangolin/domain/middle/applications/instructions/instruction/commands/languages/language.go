package languages

type language struct {
	variable string
	values   []Value
}

func createLanguage(
	variable string,
	values []Value,
) Language {
	out := language{
		variable: variable,
		values:   values,
	}

	return &out
}

// Variable returns the variable
func (obj *language) Variable() string {
	return obj.variable
}

// Values returns the values
func (obj *language) Values() []Value {
	return obj.values
}
