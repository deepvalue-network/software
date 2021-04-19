package token

type codeMatch struct {
	ret         string
	sectionName string
	token       string
	patterns    []string
}

func createCodeMatch(
	ret string,
	sectionName string,
	token string,
	patterns []string,
) CodeMatch {
	out := codeMatch{
		ret:         ret,
		sectionName: sectionName,
		token:       token,
		patterns:    patterns,
	}

	return &out
}

// Return returns the return variable
func (obj *codeMatch) Return() string {
	return obj.ret
}

// SectionName returns the sectionName
func (obj *codeMatch) SectionName() string {
	return obj.sectionName
}

// Token returns the token
func (obj *codeMatch) Token() string {
	return obj.token
}

// Patterns returns the patterns
func (obj *codeMatch) Patterns() []string {
	return obj.patterns
}
