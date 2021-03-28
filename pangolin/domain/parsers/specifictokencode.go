package parsers

type specificTokenCode struct {
	variableName    string
	patternVariable string
	amount          string
}

func createSpecificTokenCode(
	variableName string,
	patternVariable string,
) SpecificTokenCode {
	return createSpecificTokenCodeInternally(variableName, patternVariable, "")
}

func createSpecificTokenCodeWithAmount(
	variableName string,
	patternVariable string,
	amount string,
) SpecificTokenCode {
	return createSpecificTokenCodeInternally(variableName, patternVariable, amount)
}

func createSpecificTokenCodeInternally(
	variableName string,
	patternVariable string,
	amount string,
) SpecificTokenCode {
	out := specificTokenCode{
		variableName:    variableName,
		patternVariable: patternVariable,
		amount:          amount,
	}

	return &out
}

// VariableName returns the variable name
func (obj *specificTokenCode) VariableName() string {
	return obj.variableName
}

// PatternVariable returns the pattern variable
func (obj *specificTokenCode) PatternVariable() string {
	return obj.patternVariable
}

// HasAmount returns true if there is an amount, false otherwise
func (obj *specificTokenCode) HasAmount() bool {
	return obj.amount != ""
}

// Amount returns the amount if any
func (obj *specificTokenCode) Amount() string {
	return obj.amount
}
