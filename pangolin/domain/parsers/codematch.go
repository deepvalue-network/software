package parsers

type codeMatch struct {
	content          string
	section          string
	patternVariables []string
}

func createCodeMatch(
	content string,
	section string,
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
func (obj *codeMatch) Content() string {
	return obj.content
}

// Section returns the section
func (obj *codeMatch) Section() string {
	return obj.section
}

// PatternVariables returns the patternVariables
func (obj *codeMatch) PatternVariables() []string {
	return obj.patternVariables
}
