package parsers

type tokenCode struct {
	content       VariableName
	tokenVariable string
}

func createTokenCode(
	content VariableName,
	tokenVariable string,
) TokenCode {
	out := tokenCode{
		content:       content,
		tokenVariable: tokenVariable,
	}

	return &out
}

// Content returns the content
func (obj *tokenCode) Content() VariableName {
	return obj.content
}

// TokenVariable returns the tokenVariable
func (obj *tokenCode) TokenVariable() string {
	return obj.tokenVariable
}
