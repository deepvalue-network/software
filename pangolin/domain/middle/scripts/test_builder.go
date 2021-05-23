package scripts

import "errors"

type testBuilder struct {
	name string
	path string
}

func createTestBuilder() TestBuilder {
	out := testBuilder{
		name: "",
		path: "",
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

// WithPath adds a path to the builder
func (app *testBuilder) WithPath(path string) TestBuilder {
	app.path = path
	return app
}

// Now builds a new Test instance
func (app *testBuilder) Now() (Test, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Test instance")
	}

	if app.path == "" {
		return nil, errors.New("the path is mandatory in order to build a Test instance")
	}

	return createTest(app.name, app.path), nil
}
