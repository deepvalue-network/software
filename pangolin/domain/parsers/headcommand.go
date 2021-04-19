package parsers

type headCommand struct {
	variable string
	values   []HeadValue
}

func createHeadCommand(
	variable string,
	values []HeadValue,
) HeadCommand {
	out := headCommand{
		variable: variable,
		values:   values,
	}

	return &out
}

// Variable return the variable
func (obj *headCommand) Variable() string {
	return obj.variable
}

// Values return the values
func (obj *headCommand) Values() []HeadValue {
	return obj.values
}
