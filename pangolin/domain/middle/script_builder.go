package middle

import "errors"

type scriptBuilder struct {
	name    string
	version string
	lang    string
	script  string
}

func createScriptBuilder() ScriptBuilder {
	out := scriptBuilder{
		name:    "",
		version: "",
		lang:    "",
		script:  "",
	}

	return &out
}

// Create initializes the builder
func (app *scriptBuilder) Create() ScriptBuilder {
	return createScriptBuilder()
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

// WithLanguagePath adds a language path to the builder
func (app *scriptBuilder) WithLanguagePath(lang string) ScriptBuilder {
	app.lang = lang
	return app
}

// WithScriptPath adds a script path to the builder
func (app *scriptBuilder) WithScriptPath(script string) ScriptBuilder {
	app.script = script
	return app
}

// Now builds  new script instance
func (app *scriptBuilder) Now() (Script, error) {
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
