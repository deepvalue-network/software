package parsers

type variableIncoming struct {
	isMandatory bool
	def         Value
}

func createVariableIncomingWithMandatory() VariableIncoming {
	return createVariableIncomingInternally(true, nil)
}

func createVariableIncomingWithOptional(def Value) VariableIncoming {
	return createVariableIncomingInternally(false, def)
}

func createVariableIncomingInternally(
	isMandatory bool,
	def Value,
) VariableIncoming {
	out := variableIncoming{
		isMandatory: isMandatory,
		def:         def,
	}

	return &out
}

// IsMandatory returns true if the variable is mandatory, false otherwise
func (obj *variableIncoming) IsMandatory() bool {
	return obj.isMandatory
}

// IsOptional returns true if the variable is optional, false otherwise
func (obj *variableIncoming) IsOptional() bool {
	return obj.def != nil
}

// OptionalDefaultValue returns the optional's default value, if any
func (obj *variableIncoming) OptionalDefaultValue() Value {
	return obj.def
}
