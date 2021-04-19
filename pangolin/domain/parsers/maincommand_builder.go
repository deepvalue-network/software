package parsers

import "errors"

type mainCommandBuilder struct {
	variable string
	ins      []MainCommandInstruction
}

func createMainCommandBuilder() MainCommandBuilder {
	out := mainCommandBuilder{
		variable: "",
		ins:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *mainCommandBuilder) Create() MainCommandBuilder {
	return createMainCommandBuilder()
}

// WithVariable adds a variable to the builder
func (app *mainCommandBuilder) WithVariable(variable string) MainCommandBuilder {
	app.variable = variable
	return app
}

// WithInstructions adds instructions to the builder
func (app *mainCommandBuilder) WithInstructions(ins []MainCommandInstruction) MainCommandBuilder {
	app.ins = ins
	return app
}

// Now builds a new MaiNCommand instance
func (app *mainCommandBuilder) Now() (MainCommand, error) {
	if app.variable == "" {
		return nil, errors.New("the variable is mandatory in order to build a MainCommand instance")
	}

	if app.ins != nil && len(app.ins) <= 0 {
		app.ins = nil
	}

	if app.ins == nil {
		return nil, errors.New("the instructions are mandatory in order to build a MainCommand instance")
	}

	return createMainCommand(app.variable, app.ins), nil
}
