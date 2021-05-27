package linkers

import "errors"

type testBuilder struct {
	name       string
	executable Executable
}

func createTestBuilder() TestBuilder {
	out := testBuilder{
		name:       "",
		executable: nil,
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

// WithExecutable adds an executable to the builder
func (app *testBuilder) WithExecutable(executable Executable) TestBuilder {
	app.executable = executable
	return app
}

// Now builds a new Test instance
func (app *testBuilder) Now() (Test, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Test instance")
	}

	if app.executable == nil {
		return nil, errors.New("the executable is mandatory in order to build a Test instance")
	}

	return createTest(app.name, app.executable), nil
}
