package machines

type languageStateFactory struct {
}

func createLanguageStateFactory() LanguageStateFactory {
	out := languageStateFactory{}
	return &out
}

// Create creates a new language state instance
func (app *languageStateFactory) Create() LanguageState {
	return createLanguageState()
}
