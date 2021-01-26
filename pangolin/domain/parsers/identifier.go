package parsers

type identifier struct {
	variableName VariableName
	constant     string
}

func createIdentifierWithVariableName(variableName VariableName) Identifier {
	return createIdentifierInternally(variableName, "")
}

func createIdentifierWithConstant(constant string) Identifier {
	return createIdentifierInternally(nil, constant)
}

func createIdentifierInternally(variableName VariableName, constant string) Identifier {
	out := identifier{
		variableName: variableName,
		constant:     constant,
	}

	return &out
}

// IsVariable returns true if there is a variable, false otherwise
func (obj *identifier) IsVariable() bool {
	return obj.variableName != nil
}

// Variable returns the variableName, if any
func (obj *identifier) Variable() VariableName {
	return obj.variableName
}

// IsConstant returns true if there is a constant, false otherwise
func (obj *identifier) IsConstant() bool {
	return obj.constant != ""
}

// Constant returns the constant, if any
func (obj *identifier) Constant() string {
	return obj.constant
}

// String returns the stirng
func (obj *identifier) String() string {
	if obj.IsVariable() {
		return obj.Variable().String()
	}

	return obj.Constant()
}
