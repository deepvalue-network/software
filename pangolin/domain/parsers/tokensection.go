package parsers

type tokenSection struct {
	code     TokenCode
	specific SpecificTokenCode
}

func createTokenSectionWithCode(
	code TokenCode,
) TokenSection {
	return createTokenSectionInternally(code, nil)
}

func createTokenSectionWithSpecific(
	specific SpecificTokenCode,
) TokenSection {
	return createTokenSectionInternally(nil, specific)
}

func createTokenSectionInternally(
	code TokenCode,
	specific SpecificTokenCode,
) TokenSection {
	out := tokenSection{
		code:     code,
		specific: specific,
	}

	return &out
}

// IsCode returns true if there is code, false otherwise
func (obj *tokenSection) IsCode() bool {
	return obj.code != nil
}

// Code returns the code, if any
func (obj *tokenSection) Code() TokenCode {
	return obj.code
}

// IsSpecific returns true if there is specific code, false otherwise
func (obj *tokenSection) IsSpecific() bool {
	return obj.specific != nil
}

// Specific returns the specific code, if any
func (obj *tokenSection) Specific() SpecificTokenCode {
	return obj.specific
}
