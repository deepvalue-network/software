package parsers

type language struct {
	app LanguageApplication
	def LanguageDefinition
}

func createLanguageWithApplication(
	app LanguageApplication,
) Language {
	return createLanguageInternally(app, nil)
}

func createLanguageWithDefinition(
	def LanguageDefinition,
) Language {
	return createLanguageInternally(nil, def)
}

func createLanguageInternally(
	app LanguageApplication,
	def LanguageDefinition,
) Language {
	out := language{
		app: app,
		def: def,
	}

	return &out
}

// IsApplication returns true if there is an application, false otherwise
func (obj *language) IsApplication() bool {
	return obj.app != nil
}

// Application returns the application, if any
func (obj *language) Application() LanguageApplication {
	return obj.app
}

// IsDefinition returns true if there is a definition, false otherwise
func (obj *language) IsDefinition() bool {
	return obj.def != nil
}

// Definition returns the definition, if any
func (obj *language) Definition() LanguageDefinition {
	return obj.def
}
