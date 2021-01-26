package parsers

type codeMatch struct {
	content          VariableName
	section          VariableName
	tokenVariable    string
	patternVariables []string
}

func createCodeMatch(
	content VariableName,
	section VariableName,
	tokenVariable string,
	patternVariables []string,
) CodeMatch {
	out := codeMatch{
		content:          content,
		section:          section,
		tokenVariable:    tokenVariable,
		patternVariables: patternVariables,
	}

	return &out
}

// Content returns the content
func (obj *codeMatch) Content() VariableName {
	return obj.content
}

// Section returns the section
func (obj *codeMatch) Section() VariableName {
	return obj.section
}

// TokenVariable returns the tokenVariable
func (obj *codeMatch) TokenVariable() string {
	return obj.tokenVariable
}

// PatternVariables returns the patternVariables
func (obj *codeMatch) PatternVariables() []string {
	return obj.patternVariables
}
