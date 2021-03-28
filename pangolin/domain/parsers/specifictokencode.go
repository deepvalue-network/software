package parsers

type specificTokenCode struct {
	variableName    VariableName
	patternVariable string
	amount          VariableName
}

func createSpecificTokenCode(
	variableName VariableName,
	patternVariable string,
) SpecificTokenCode {
	return createSpecificTokenCodeInternally(variableName, patternVariable, nil)
}

func createSpecificTokenCodeWithAmount(
	variableName VariableName,
	patternVariable string,
	amount VariableName,
) SpecificTokenCode {
	return createSpecificTokenCodeInternally(variableName, patternVariable, amount)
}

func createSpecificTokenCodeInternally(
	variableName VariableName,
	patternVariable string,
	amount VariableName,
) SpecificTokenCode {
	out := specificTokenCode{
		variableName:    variableName,
		patternVariable: patternVariable,
		amount:          amount,
	}

	return &out
}

// VariableName returns the variable name
func (obj *specificTokenCode) VariableName() VariableName {
	return obj.variableName
}

// PatternVariable returns the pattern variable
func (obj *specificTokenCode) PatternVariable() string {
	return obj.patternVariable
}

// HasAmount returns true if there is an amount, false otherwise
func (obj *specificTokenCode) HasAmount() bool {
	return obj.amount != nil
}

// Amount returns the amount if any
func (obj *specificTokenCode) Amount() VariableName {
	return obj.amount
}
