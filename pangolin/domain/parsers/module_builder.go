package parsers

import "errors"

type moduleBuilder struct {
	stackFrame string
	name       string
	symbol     string
}

func createModuleBuilder() ModuleBuilder {
	out := moduleBuilder{
		stackFrame: "",
		name:       "",
		symbol:     "",
	}

	return &out
}

// Create initializes the builder
func (app *moduleBuilder) Create() ModuleBuilder {
	return createModuleBuilder()
}

// WithStackFrame adds a stackFrame to the builder
func (app *moduleBuilder) WithStackFrame(stackFrame string) ModuleBuilder {
	app.stackFrame = stackFrame
	return app
}

// WithName adds a name to the builder
func (app *moduleBuilder) WithName(name string) ModuleBuilder {
	app.name = name
	return app
}

// WithSymbol adds a symbol to the builder
func (app *moduleBuilder) WithSymbol(symbol string) ModuleBuilder {
	app.symbol = symbol
	return app
}

// Now builds a new Module instance
func (app *moduleBuilder) Now() (Module, error) {
	if app.stackFrame == "" {
		return nil, errors.New("the stackFrame is mandatory in order to build a Module instance")
	}

	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Module instance")
	}

	if app.symbol == "" {
		return nil, errors.New("the symbol is mandatory in order to build a Module instance")
	}

	return createModule(app.stackFrame, app.name, app.symbol), nil
}
