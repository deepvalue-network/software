package parsers

import "errors"

type variableIncomingBuilder struct {
	isMandatory bool
	def         Value
}

func createVariableIncomingBuilder() VariableIncomingBuilder {
	out := variableIncomingBuilder{
		isMandatory: false,
		def:         nil,
	}

	return &out
}

// Create initializes the builder
func (app *variableIncomingBuilder) Create() VariableIncomingBuilder {
	return createVariableIncomingBuilder()
}

// IsMandatory flags the builder as mandatory
func (app *variableIncomingBuilder) IsMandatory() VariableIncomingBuilder {
	app.isMandatory = true
	return app
}

// WithOptionalDefaultValue adds an optional default value
func (app *variableIncomingBuilder) WithOptionalDefaultValue(def Value) VariableIncomingBuilder {
	app.def = def
	return app
}

// Now builds a new VariableIncoming instance
func (app *variableIncomingBuilder) Now() (VariableIncoming, error) {
	if app.isMandatory {
		return createVariableIncomingWithMandatory(), nil
	}

	if app.def != nil {
		return createVariableIncomingWithOptional(app.def), nil
	}

	return nil, errors.New("the VariableIncoming is invalid")
}
