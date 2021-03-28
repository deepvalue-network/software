package parsers

import "errors"

type readFileBuilder struct {
	variable string
	path     RelativePath
}

func createReadFileBuilder() ReadFileBuilder {
	out := readFileBuilder{
		variable: "",
		path:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *readFileBuilder) Create() ReadFileBuilder {
	return createReadFileBuilder()
}

// WithVariable adds a variable to the builder
func (app *readFileBuilder) WithVariable(variable string) ReadFileBuilder {
	app.variable = variable
	return app
}

// WithPath adds a path to the builder
func (app *readFileBuilder) WithPath(path RelativePath) ReadFileBuilder {
	app.path = path
	return app
}

// Now builds a new ReadFile instance
func (app *readFileBuilder) Now() (ReadFile, error) {
	if app.variable == "" {
		return nil, errors.New("the Variablename is mandatory in order to build a ReadFile instance")
	}

	if app.path == nil {
		return nil, errors.New("the RelativePath is mandatory in order to build a ReadFile instance")
	}

	return createReadFile(app.variable, app.path), nil
}
