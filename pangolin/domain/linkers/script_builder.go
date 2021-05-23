package linkers

import "errors"

type scriptBuilder struct {
	language LanguageReference
	name     string
	version  string
	code     string
	output   string
	tests    []Test
}

func createScriptBuilder() ScriptBuilder {
	out := scriptBuilder{
		language: nil,
		name:     "",
		version:  "",
		code:     "",
		output:   "",
		tests:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *scriptBuilder) Create() ScriptBuilder {
	return createScriptBuilder()
}

// WithLanguage adds a language to the builder
func (app *scriptBuilder) WithLanguage(language LanguageReference) ScriptBuilder {
	app.language = language
	return app
}

// WithName adds a name to the builder
func (app *scriptBuilder) WithName(name string) ScriptBuilder {
	app.name = name
	return app
}

// WithVersion adds a version to the builder
func (app *scriptBuilder) WithVersion(version string) ScriptBuilder {
	app.version = version
	return app
}

// WithCode adds code to the builder
func (app *scriptBuilder) WithCode(code string) ScriptBuilder {
	app.code = code
	return app
}

// WithOutput adds an output to the builder
func (app *scriptBuilder) WithOutput(output string) ScriptBuilder {
	app.output = output
	return app
}

// WithTests add tests to the builder
func (app *scriptBuilder) WithTests(tests []Test) ScriptBuilder {
	app.tests = tests
	return app
}

// Now builds a new Script instance
func (app *scriptBuilder) Now() (Script, error) {
	if app.language == nil {
		return nil, errors.New("the language is mandatory in order to build a Script instance")
	}

	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Script instance")
	}

	if app.version == "" {
		return nil, errors.New("the version is mandatory in order to build a Script instance")
	}

	if app.code == "" {
		return nil, errors.New("the code is mandatory in order to build a Script instance")
	}

	if app.output == "" {
		return nil, errors.New("the output is mandatory in order to build a Script instance")
	}

	if app.tests != nil && len(app.tests) <= 0 {
		app.tests = nil
	}

	if app.tests != nil {
		return createScriptWithTests(app.language, app.name, app.version, app.code, app.output, app.tests), nil
	}

	return createScript(app.language, app.name, app.version, app.code, app.output), nil
}
