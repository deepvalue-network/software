package parsers

import "errors"

type scriptValueBuilder struct {
	name       string
	version    string
	scriptPath RelativePath
	langPath   RelativePath
}

func createScriptValueBuilder() ScriptValueBuilder {
	out := scriptValueBuilder{
		name:       "",
		version:    "",
		scriptPath: nil,
		langPath:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *scriptValueBuilder) Create() ScriptValueBuilder {
	return createScriptValueBuilder()
}

// WithName adds a name to the builder
func (app *scriptValueBuilder) WithName(name string) ScriptValueBuilder {
	app.name = name
	return app
}

// WithVersion adds a version to the builder
func (app *scriptValueBuilder) WithVersion(version string) ScriptValueBuilder {
	app.version = version
	return app
}

// WithScriptPath adds a scriptPath to the builder
func (app *scriptValueBuilder) WithScriptPath(scriptPath RelativePath) ScriptValueBuilder {
	app.scriptPath = scriptPath
	return app
}

// WithLanguagePath adds a langPath to the builder
func (app *scriptValueBuilder) WithLanguagePath(langPath RelativePath) ScriptValueBuilder {
	app.langPath = langPath
	return app
}

// Now builds a new ScriptValue instance
func (app *scriptValueBuilder) Now() (ScriptValue, error) {
	if app.name != "" {
		return createScriptValueWithName(app.name), nil
	}

	if app.version != "" {
		return createScriptValueWithVersion(app.version), nil
	}

	if app.scriptPath != nil {
		return createScriptValueWithScriptPath(app.scriptPath), nil
	}

	if app.langPath != nil {
		return createScriptValueWithLanguagePath(app.langPath), nil
	}

	return nil, errors.New("the ScriptValue is invalid")
}
