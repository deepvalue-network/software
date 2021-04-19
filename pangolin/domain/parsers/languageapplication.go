package parsers

type languageApplication struct {
	head   HeadSection
	labels LanguageLabelSection
	main   LanguageMainSection
	tests  LanguageTestSection
}

func createLanguageApplication(
	head HeadSection,
	labels LanguageLabelSection,
	main LanguageMainSection,
) LanguageApplication {
	return createLanguageApplicationInternally(head, labels, main, nil)
}

func createLanguageApplicationWithTests(
	head HeadSection,
	labels LanguageLabelSection,
	main LanguageMainSection,
	tests LanguageTestSection,
) LanguageApplication {
	return createLanguageApplicationInternally(head, labels, main, tests)
}

func createLanguageApplicationInternally(
	head HeadSection,
	labels LanguageLabelSection,
	main LanguageMainSection,
	tests LanguageTestSection,
) LanguageApplication {
	out := languageApplication{
		head:   head,
		labels: labels,
		main:   main,
		tests:  tests,
	}

	return &out
}

// Head returns the head section
func (obj *languageApplication) Head() HeadSection {
	return obj.head
}

// Labels returns the labels section
func (obj *languageApplication) Labels() LanguageLabelSection {
	return obj.labels
}

// Main returns the main section
func (obj *languageApplication) Main() LanguageMainSection {
	return obj.main
}

// HasTests returns true if there is tests, false otherwise
func (obj *languageApplication) HasTests() bool {
	return obj.tests != nil
}

// Tests returns the tests, if any
func (obj *languageApplication) Tests() LanguageTestSection {
	return obj.tests
}
