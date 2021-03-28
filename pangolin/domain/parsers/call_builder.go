package parsers

import "errors"

type callBuilder struct {
	name      string
	condition string
}

func createCallBuilder() CallBuilder {
	out := callBuilder{
		name:      "",
		condition: "",
	}

	return &out
}

// Create initializes the builder
func (app *callBuilder) Create() CallBuilder {
	return createCallBuilder()
}

// WithName adds a name to the builder
func (app *callBuilder) WithName(name string) CallBuilder {
	app.name = name
	return app
}

// WithCondition adds a condition to the builder
func (app *callBuilder) WithCondition(condition string) CallBuilder {
	app.condition = condition
	return app
}

// Now builds a new Call instance
func (app *callBuilder) Now() (Call, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Call instance")
	}

	if app.condition != "" {
		return createCallWithCondition(app.name, app.condition), nil
	}

	return createCall(app.name), nil
}
