package parsers

type tokenSection struct {
	variableName VariableName
	specific     SpecificTokenCode
}

func createTokenSectionWithVariableName(
	variableName VariableName,
) TokenSection {
	return createTokenSectionInternally(variableName, nil)
}

func createTokenSectionWithSpecific(
	specific SpecificTokenCode,
) TokenSection {
	return createTokenSectionInternally(nil, specific)
}

func createTokenSectionInternally(
	variableName VariableName,
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
	return obj.variableName != nil
}

// VariableName returns the variableName, if any
func (obj *tokenSection) VariableName() VariableName {
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
