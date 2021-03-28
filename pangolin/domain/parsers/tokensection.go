package parsers

type tokenSection struct {
	variableName string
	specific     SpecificTokenCode
}

func createTokenSectionWithstring(
	variableName string,
) TokenSection {
	return createTokenSectionInternally(variableName, nil)
}

func createTokenSectionWithSpecific(
	specific SpecificTokenCode,
) TokenSection {
	return createTokenSectionInternally("", specific)
}

func createTokenSectionInternally(
	variableName string,
	specific SpecificTokenCode,
) TokenSection {
	out := tokenSection{
		variableName: variableName,
		specific:     specific,
	}

	return &out
}

// IsVariableName returns true if there is a variableName, false otherwise
func (obj *tokenSection) IsVariableName() bool {
	return obj.variableName != ""
}

// VariableName returns the variableName, if any
func (obj *tokenSection) VariableName() string {
	return obj.variableName
}

// IsSpecific returns true if there is specific code, false otherwise
func (obj *tokenSection) IsSpecific() bool {
	return obj.specific != nil
}

// Specific returns the specific code, if any
func (obj *tokenSection) Specific() SpecificTokenCode {
	return obj.specific
}
