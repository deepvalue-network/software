package parsers

type languageCommand struct {
	variable string
	values   []LanguageValue
}

func createLanguageCommand(
	variable string,
	values []LanguageValue,
) LanguageCommand {
	out := languageCommand{
		variable: variable,
		values:   values,
	}

	return &out
}

// Variable returns the variable
func (obj *languageCommand) Variable() string {
	return obj.variable
}

// Values returns the values
func (obj *languageCommand) Values() []LanguageValue {
	return obj.values
}
