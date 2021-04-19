package parsers

import "errors"

type languageApplicationBuilder struct {
	head   HeadSection
	labels LanguageLabelSection
	main   LanguageMainSection
	tests  LanguageTestSection
}

func createLanguageApplicationBuilder() LanguageApplicationBuilder {
	out := languageApplicationBuilder{
		head:   nil,
		labels: nil,
		main:   nil,
		tests:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *languageApplicationBuilder) Create() LanguageApplicationBuilder {
	return createLanguageApplicationBuilder()
}

// WithHead adds an head to the builder
func (app *languageApplicationBuilder) WithHead(head HeadSection) LanguageApplicationBuilder {
	app.head = head
	return app
}

// WithLabels add labels to the builder
func (app *languageApplicationBuilder) WithLabels(labels LanguageLabelSection) LanguageApplicationBuilder {
	app.labels = labels
	return app
}

// WithMain add main to the builder
func (app *languageApplicationBuilder) WithMain(main LanguageMainSection) LanguageApplicationBuilder {
	app.main = main
	return app
}

// WithTests add tests to the builder
func (app *languageApplicationBuilder) WithTests(tests LanguageTestSection) LanguageApplicationBuilder {
	app.tests = tests
	return app
}

// Now builds a new LanguageApplication instance
func (app *languageApplicationBuilder) Now() (LanguageApplication, error) {
	if app.head == nil {
		return nil, errors.New("the head section is mandatory in order to build a LanguageApplication instance")
	}

	if app.labels == nil {
		return nil, errors.New("the labels section is mandatory in order to build a LanguageApplication instance")
	}

	if app.main == nil {
		return nil, errors.New("the main section is mandatory in order to build a LanguageApplication instance")
	}

	if app.tests != nil {
		return createLanguageApplicationWithTests(app.head, app.labels, app.main, app.tests), nil
	}

	return createLanguageApplication(app.head, app.labels, app.main), nil
}
