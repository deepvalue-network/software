package middle

type program struct {
	app    Application
	lang   Language
	script Script
}

func createProgramWithApplication(app Application) Program {
	return createProgramInternally(app, nil, nil)
}

func createProgramWithLanguage(lang Language) Program {
	return createProgramInternally(nil, lang, nil)
}

func createProgramWithScript(script Script) Program {
	return createProgramInternally(nil, nil, script)
}

func createProgramInternally(
	app Application,
	lang Language,
	script Script,
) Program {
	out := program{
		app:    app,
		lang:   lang,
		script: script,
	}

	return &out
}

// IsApplication returns true if the program is an application, false otherwise
func (obj *program) IsApplication() bool {
	return obj.app != nil
}

// Application returns the application, if any
func (obj *program) Application() Application {
	return obj.app
}

// IsLanguage returns true if the program is a language, false otherwise
func (obj *program) IsLanguage() bool {
	return obj.lang != nil
}

// Language returns the language, if any
func (obj *program) Language() Language {
	return obj.lang
}

// IsScript returns true if the program is a script, false otherwise
func (obj *program) IsScript() bool {
	return obj.script != nil
}

// Script returns the script, if any
func (obj *program) Script() Script {
	return obj.script
}
