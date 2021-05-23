package parsers

import "errors"

type scriptTestsBuilder struct {
	list []ScriptTest
}

func createScriptTestsBuilder() ScriptTestsBuilder {
	out := scriptTestsBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *scriptTestsBuilder) Create() ScriptTestsBuilder {
	return createScriptTestsBuilder()
}

// WithTests add tests to the builder
func (app *scriptTestsBuilder) WithTests(tests []ScriptTest) ScriptTestsBuilder {
	app.list = tests
	return app
}

// Now builds a new ScriptTests instance
func (app *scriptTestsBuilder) Now() (ScriptTests, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 ScriptTest instance in order to build a ScriptTests instance")
	}

	return createScriptTests(app.list), nil
}
