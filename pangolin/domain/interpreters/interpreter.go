package interpreters

type interpreter struct {
	app    Application
	lang   Language
	script Script
}

func createInterpreterWithApplication(
	app Application,
) Interpreter {
	return createInterpreterInternally(app, nil, nil)
}

func createInterpreterWithLanguage(
	lang Language,
) Interpreter {
	return createInterpreterInternally(nil, lang, nil)
}

func createInterpreterWithScript(
	script Script,
) Interpreter {
	return createInterpreterInternally(nil, nil, script)
}

func createInterpreterInternally(
	app Application,
	lang Language,
	script Script,
) Interpreter {
	out := interpreter{
		app:    app,
		lang:   lang,
		script: script,
	}

	return &out
}

// IsScript returns true if there is a script interpreter, false otherwise
func (app *interpreter) IsScript() bool {
	return app.script != nil
}

// Script returns the script interpreter, if any
func (app *interpreter) Script() Script {
	return app.script
}

// IsAppplication returns true if there is an application interpreter, false otherwise
func (app *interpreter) IsApplication() bool {
	return app.app != nil
}

// Application returns the application interpreter, if any
func (app *interpreter) Application() Application {
	return app.app
}

// IsLanguage returns true if there is a language interpreter, false otherwise
func (app *interpreter) IsLanguage() bool {
	return app.lang != nil
}

// Language returns the language interpreter, if any
func (app *interpreter) Language() Language {
	return app.lang
}
