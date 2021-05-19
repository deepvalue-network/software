package scripts

import (
	"errors"

	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

type valueBuilder struct {
	name       string
	version    string
	langPath   parsers.RelativePath
	scriptPath parsers.RelativePath
}

func createValueBuilder() ValueBuilder {
	out := valueBuilder{
		name:       "",
		version:    "",
		langPath:   nil,
		scriptPath: nil,
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
func (app *valueBuilder) WithLanguagePath(langPath parsers.RelativePath) ValueBuilder {
	app.langPath = langPath
	return app
}

// WithScriptPath adds a script path to the builder
func (app *valueBuilder) WithScriptPath(scriptPath parsers.RelativePath) ValueBuilder {
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

	if app.langPath != nil {
		return createValueWithLanguagePath(app.langPath), nil
	}

	if app.scriptPath != nil {
		return createValueWithScriptPath(app.scriptPath), nil
	}

	return nil, errors.New("the Value is invalid")
}
