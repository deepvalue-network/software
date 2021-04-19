package parsers

type scriptCommand struct {
	variable string
	values   []ScriptValue
}

func createScriptCommand(
	variable string,
	values []ScriptValue,
) ScriptCommand {
	out := scriptCommand{
		variable: variable,
		values:   values,
	}

	return &out
}

// Variable returns the variable
func (obj *scriptCommand) Variable() string {
	return obj.variable
}

// Values returns the values
func (obj *scriptCommand) Values() []ScriptValue {
	return obj.values
}
