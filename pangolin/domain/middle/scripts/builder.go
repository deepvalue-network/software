package scripts

import "errors"

type builder struct {
	name    string
	version string
	lang    string
	script  string
}

func createBuilder() Builder {
	out := builder{
		name:    "",
		version: "",
		lang:    "",
		script:  "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// WithVersion adds a version to the builder
func (app *builder) WithVersion(version string) Builder {
	app.version = version
	return app
}

// WithLanguagePath adds a language path to the builder
func (app *builder) WithLanguagePath(lang string) Builder {
	app.lang = lang
	return app
}

// WithScriptPath adds a script path to the builder
func (app *builder) WithScriptPath(script string) Builder {
	app.script = script
	return app
}

// Now builds  new script instance
func (app *builder) Now() (Script, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Script instance")
	}

	if app.version == "" {
		return nil, errors.New("the version is mandatory in order to build a Script instance")
	}

	if app.lang == "" {
		return nil, errors.New("the language path is mandatory in order to build a Script instance")
	}

	if app.script == "" {
		return nil, errors.New("the script path is mandatory in order to build a Script instance")
	}

	return createScript(app.name, app.version, app.lang, app.script), nil
}
