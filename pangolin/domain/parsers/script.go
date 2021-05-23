package parsers

type script struct {
	name     string
	version  string
	script   RelativePath
	language RelativePath
	output   string
	tests    ScriptTests
}

func createScript(
	name string,
	version string,
	scriptPath RelativePath,
	language RelativePath,
	output string,
) Script {
	return createScriptInternally(name, version, scriptPath, language, output, nil)
}

func createScriptWithTests(
	name string,
	version string,
	scriptPath RelativePath,
	language RelativePath,
	output string,
	tests ScriptTests,
) Script {
	return createScriptInternally(name, version, scriptPath, language, output, tests)
}

func createScriptInternally(
	name string,
	version string,
	scriptPath RelativePath,
	language RelativePath,
	output string,
	tests ScriptTests,
) Script {
	out := script{
		name:     name,
		version:  version,
		script:   scriptPath,
		language: language,
		output:   output,
		tests:    tests,
	}

	return &out
}

// Name returns the name
func (obj *script) Name() string {
	return obj.name
}

// Version returns the version
func (obj *script) Version() string {
	return obj.version
}

// Script returns the script
func (obj *script) Script() RelativePath {
	return obj.script
}

// Language returns the language
func (obj *script) Language() RelativePath {
	return obj.language
}

// Output returns the output variable
func (obj *script) Output() string {
	return obj.output
}

// HasTests returns true if there is tests, false otherwise
func (obj *script) HasTests() bool {
	return obj.tests != nil
}

// Tests returns the tests, if any
func (obj *script) Tests() ScriptTests {
	return obj.tests
}
