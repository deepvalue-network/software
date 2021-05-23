package linkers

import "errors"

type testBuilder struct {
	name   string
	script Script
}

func createTestBuilder() TestBuilder {
	out := testBuilder{
		name:   "",
		script: nil,
	}

	return &out
}

// Create initializes the builder
func (app *testBuilder) Create() TestBuilder {
	return createTestBuilder()
}

// WithName adds a name to the builder
func (app *testBuilder) WithName(name string) TestBuilder {
	app.name = name
	return app
}

// WithScript adds a script to the builder
func (app *testBuilder) WithScript(script Script) TestBuilder {
	app.script = script
	return app
}

// Now builds a new Test instance
func (app *testBuilder) Now() (Test, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Test instance")
	}

	if app.script == nil {
		return nil, errors.New("the script is mandatory in order to build a Test instance")
	}

	return createTest(app.name, app.script), nil
}
