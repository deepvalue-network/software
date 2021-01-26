package parsers

type specificTokenCode struct {
	content         VariableName
	tokenVariable   string
	patternVariable string
	amount          VariableName
}

func createSpecificTokenCode(
	content VariableName,
	tokenVariable string,
	patternVariable string,
) SpecificTokenCode {
	return createSpecificTokenCodeInternally(content, tokenVariable, patternVariable, nil)
}

func createSpecificTokenCodeWithAmount(
	content VariableName,
	tokenVariable string,
	patternVariable string,
	amount VariableName,
) SpecificTokenCode {
	return createSpecificTokenCodeInternally(content, tokenVariable, patternVariable, amount)
}

func createSpecificTokenCodeInternally(
	content VariableName,
	tokenVariable string,
	patternVariable string,
	amount VariableName,
) SpecificTokenCode {
	out := specificTokenCode{
		content:         content,
		tokenVariable:   tokenVariable,
		patternVariable: patternVariable,
		amount:          amount,
	}

	return &out
}

// Content returns the content
func (obj *specificTokenCode) Content() VariableName {
	return obj.content
}

// TokenVariable returns the tokenVariable
func (obj *specificTokenCode) TokenVariable() string {
	return obj.tokenVariable
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
