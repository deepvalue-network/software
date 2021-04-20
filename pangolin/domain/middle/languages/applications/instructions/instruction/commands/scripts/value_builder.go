package scripts

import "errors"

type valueBuilder struct {
	name       string
	version    string
	langPath   string
	scriptPath string
}

func createValueBuilder() ValueBuilder {
	out := valueBuilder{
		name:       "",
		version:    "",
		langPath:   "",
		scriptPath: "",
	}

	return &out
}

// Create initializes the builder
func (app *valueBuilder) Create() ValueBuilder {
	return createValueBuilder()
}

// WithName adds a name to the builder
func (app *valueBuilder) WithName(name string) ValueBuilder {
	app.name = name
	return app
}

// WithVersion adds a version to the builder
func (app *valueBuilder) WithVersion(version string) ValueBuilder {
	app.version = version
	return app
}

// WithLanguagePath adds a language path to the builder
func (app *valueBuilder) WithLanguagePath(langPath string) ValueBuilder {
	app.langPath = langPath
	return app
}

// WithScriptPath adds a script path to the builder
func (app *valueBuilder) WithScriptPath(scriptPath string) ValueBuilder {
	app.scriptPath = scriptPath
	return app
}

// Now builds a new Value instance
func (app *valueBuilder) Now() (Value, error) {
	if app.name != "" {
		return createValueWithName(app.name), nil
	}

	if app.version != "" {
		return createValueWithVersion(app.version), nil
	}

	if app.langPath != "" {
		return createValueWithLanguagePath(app.langPath), nil
	}

	if app.scriptPath != "" {
		return createValueWithScriptPath(app.scriptPath), nil
	}

	return nil, errors.New("the Value is invalid")
}
