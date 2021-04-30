package linkers

type language struct {
	ref LanguageReference
	app LanguageApplication
}

func createLanguageWithReference(
	ref LanguageReference,
) Language {
	return createLanguageInternally(ref, nil)
}

func createLanguageWithLanguageApplication(
	app LanguageApplication,
) Language {
	return createLanguageInternally(nil, app)
}

func createLanguageInternally(
	ref LanguageReference,
	app LanguageApplication,
) Language {
	out := language{
		ref: ref,
		app: app,
	}

	return &out
}

// IsReference returns true if there is a reference, false otherwise
func (obj *language) IsReference() bool {
	return obj.ref != nil
}

// Reference returns the reference, if any
func (obj *language) Reference() LanguageReference {
	return obj.ref
}

// IsApplication returns true if there is an application, false otherwise
func (obj *language) IsApplication() bool {
	return obj.app != nil
}

// Application returns the application, if any
func (obj *language) Application() LanguageApplication {
	return obj.app
}
