package parsers

import "errors"

type scriptTestBuilder struct {
	name string
	path RelativePath
}

func createScriptTestBuilder() ScriptTestBuilder {
	out := scriptTestBuilder{
		name: "",
		path: nil,
	}

	return &out
}

// Create initializes the builder
func (app *scriptTestBuilder) Create() ScriptTestBuilder {
	return createScriptTestBuilder()
}

// WithName adds a name to the builder
func (app *scriptTestBuilder) WithName(name string) ScriptTestBuilder {
	app.name = name
	return app
}

// WithPath adds a path to the builder
func (app *scriptTestBuilder) WithPath(path RelativePath) ScriptTestBuilder {
	app.path = path
	return app
}

// Now builds a new ScriptTest instance
func (app *scriptTestBuilder) Now() (ScriptTest, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a ScriptTest instance")
	}

	if app.path == nil {
		return nil, errors.New("the RelativePath is mandatory in order to build a ScriptTest instance")
	}

	return createScriptTest(app.name, app.path), nil
}
