package languages

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/definitions"
)

type language struct {
	def definitions.Definition
	app applications.Application
}

func createLanguageWithDefinition(
	def definitions.Definition,
) Language {
	return createLanguageInternally(def, nil)
}

func createLanguageWithApplication(
	app applications.Application,
) Language {
	return createLanguageInternally(nil, app)
}

func createLanguageInternally(
	def definitions.Definition,
	app applications.Application,
) Language {
	out := language{
		def: def,
		app: app,
	}

	return &out
}

// IsDefinition returns true if there is a definition, false otherwise
func (obj *language) IsDefinition() bool {
	return obj.def != nil
}

// Definition returns the definition, if any
func (obj *language) Definition() definitions.Definition {
	return obj.def
}

// IsApplication returns true if there is an application, false otherwise
func (obj *language) IsApplication() bool {
	return obj.app != nil
}

// Application returns the application, if any
func (obj *language) Application() applications.Application {
	return obj.app
}
