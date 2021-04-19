package parsers

import "errors"

type labelCommandBuilder struct {
	variable string
	name     string
	ins      []LabelCommandInstruction
}

func createLabelCommandBuilder() LabelCommandBuilder {
	out := labelCommandBuilder{
		variable: "",
		name:     "",
		ins:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *labelCommandBuilder) Create() LabelCommandBuilder {
	return createLabelCommandBuilder()
}

// WithVariable adds a variable to the builder
func (app *labelCommandBuilder) WithVariable(variable string) LabelCommandBuilder {
	app.variable = variable
	return app
}

// WithName adds a name to the builder
func (app *labelCommandBuilder) WithName(name string) LabelCommandBuilder {
	app.name = name
	return app
}

// WithInstructions adds instructions to the builder
func (app *labelCommandBuilder) WithInstructions(ins []LabelCommandInstruction) LabelCommandBuilder {
	app.ins = ins
	return app
}

// Now builds a new LabelCommand instance
func (app *labelCommandBuilder) Now() (LabelCommand, error) {
	if app.variable == "" {
		return nil, errors.New("the variable is mandatory in order to build a LabelCommand instance")
	}

	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a LabelCommand instance")
	}

	if app.ins != nil && len(app.ins) <= 0 {
		app.ins = nil
	}

	if app.ins == nil {
		return nil, errors.New("the instruction is mandatory in order to build a LabelCommand instance")
	}

	return createLabelCommand(app.variable, app.name, app.ins), nil
}
