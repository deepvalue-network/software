package parsers

type token struct {
	codeMatch    CodeMatch
	tokenSection TokenSection
}

func createTokenWithCodeMatch(
	codeMatch CodeMatch,
) Token {
	return createTokenInternally(codeMatch, nil)
}

func createTokenWithTokenSection(
	tokenSection TokenSection,
) Token {
	return createTokenInternally(nil, tokenSection)
}

func createTokenInternally(
	codeMatch CodeMatch,
	tokenSection TokenSection,
) Token {
	out := token{
		codeMatch:    codeMatch,
		tokenSection: tokenSection,
	}

	return &out
}

// IsCodeMatch returns true if there is a codeMatch, false otherwise
func (obj *token) IsCodeMatch() bool {
	return obj.codeMatch != nil
}

// CodeMatch returns the codeMatch, if any
func (obj *token) CodeMatch() CodeMatch {
	return obj.codeMatch
}

// IsTokenSection returns true if there is a tokenSection, false otherwise
func (obj *token) IsTokenSection() bool {
	return obj.tokenSection != nil
}

// TokenSection returns the tokenSection, if any
func (obj *token) TokenSection() TokenSection {
	return obj.tokenSection
}
