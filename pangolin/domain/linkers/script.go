package linkers

type script struct {
	language LanguageReference
	name     string
	version  string
	code     string
	output   string
	tests    []Test
}

func createScript(
	language LanguageReference,
	name string,
	version string,
	code string,
	output string,
) Script {
	return createScriptInternally(language, name, version, code, output, nil)
}

func createScriptWithTests(
	language LanguageReference,
	name string,
	version string,
	code string,
	output string,
	tests []Test,
) Script {
	return createScriptInternally(language, name, version, code, output, tests)
}

func createScriptInternally(
	language LanguageReference,
	name string,
	version string,
	code string,
	output string,
	tests []Test,
) Script {
	out := script{
		language: language,
		name:     name,
		version:  version,
		code:     code,
		output:   output,
		tests:    tests,
	}

	return &out
}

// Language returns the language
func (obj *script) Language() LanguageReference {
	return obj.language
}

// Name returns the name
func (obj *script) Name() string {
	return obj.name
}

// Version returns the version
func (obj *script) Version() string {
	return obj.version
}

// Code returns the code
func (obj *script) Code() string {
	return obj.code
}

// Output returns the output
func (obj *script) Output() string {
	return obj.output
}

// HasTests returns true if there is tests, false otherwise
func (obj *script) HasTests() bool {
	return obj.tests != nil
}

// Tests returns the tests, if any
func (obj *script) Tests() []Test {
	return obj.tests
}
