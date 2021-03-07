package call

import "errors"

type builder struct {
	name      string
	condition string
}

func createBuilder() Builder {
	out := builder{
		name:      "",
		condition: "",
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

	if app.condition != "" {
		return createCallWithCondition(app.name, app.condition), nil
	}

	return createCall(app.name), nil
}
