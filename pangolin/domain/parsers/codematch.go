package parsers

type codeMatch struct {
	content          VariableName
	section          VariableName
	patternVariables []string
}

func createCodeMatch(
	content VariableName,
	section VariableName,
	patternVariables []string,
) CodeMatch {
	out := codeMatch{
		content:          content,
		section:          section,
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

// PatternVariables returns the patternVariables
func (obj *codeMatch) PatternVariables() []string {
	return obj.patternVariables
}
