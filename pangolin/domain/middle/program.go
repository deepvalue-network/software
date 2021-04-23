package middle

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages"
	"github.com/deepvalue-network/software/pangolin/domain/middle/scripts"
)

type program struct {
	app    applications.Application
	lang   languages.Language
	script scripts.Script
}

func createProgramWithApplication(app applications.Application) Program {
	return createProgramInternally(app, nil, nil)
}

func createProgramWithLanguage(lang languages.Language) Program {
	return createProgramInternally(nil, lang, nil)
}

func createProgramWithScript(script scripts.Script) Program {
	return createProgramInternally(nil, nil, script)
}

func createProgramInternally(
	app applications.Application,
	lang languages.Language,
	script scripts.Script,
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
func (obj *program) Application() applications.Application {
	return obj.app
}

// IsLanguage returns true if the program is a language, false otherwise
func (obj *program) IsLanguage() bool {
	return obj.lang != nil
}

// Language returns the language, if any
func (obj *program) Language() languages.Language {
	return obj.lang
}

// IsScript returns true if the program is a script, false otherwise
func (obj *program) IsScript() bool {
	return obj.script != nil
}

// Script returns the script, if any
func (obj *program) Script() scripts.Script {
	return obj.script
}
