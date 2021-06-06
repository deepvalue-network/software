package module

import "errors"

type builder struct {
	stackFrame string
	name       string
	symbol     string
}

func createBuilder() Builder {
	out := builder{
		stackFrame: "",
		name:       "",
		symbol:     "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithStackFrame adds a stackFrame to the builder
func (app *builder) WithStackFrame(stackFrame string) Builder {
	app.stackFrame = stackFrame
	return app
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// WithSymbol adds a symbol to the builder
func (app *builder) WithSymbol(symbol string) Builder {
	app.symbol = symbol
	return app
}

// Now builds a new Module instance
func (app *builder) Now() (Module, error) {
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
