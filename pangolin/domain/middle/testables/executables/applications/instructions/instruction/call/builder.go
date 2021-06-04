package call

import "errors"

type builder struct {
	name       string
	stackFrame string
	condition  string
}

func createBuilder() Builder {
	out := builder{
		name:       "",
		stackFrame: "",
		condition:  "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// WithStackFrame adds a stackFrame to the builder
func (app *builder) WithStackFrame(stackFrame string) Builder {
	app.stackFrame = stackFrame
	return app
}

// WithCondition adds a condition to the builder
func (app *builder) WithCondition(condition string) Builder {
	app.condition = condition
	return app
}

// Now builds a new Call instance
func (app *builder) Now() (Call, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Call instance")
	}

	if app.stackFrame == "" {
		return nil, errors.New("the stackFrame is mandatory in order to build a Call instance")
	}

	if app.condition != "" {
		return createCallWithCondition(app.name, app.stackFrame, app.condition), nil
	}

	return createCall(app.name, app.stackFrame), nil
}
