package scripts

type script struct {
	name    string
	version string
	lang    string
	script  string
	output  string
	tests   Tests
}

func createScript(
	name string,
	version string,
	lang string,
	scriptPath string,
	output string,
) Script {
	return createScriptInternally(name, version, lang, scriptPath, output, nil)
}

func createScriptWithTests(
	name string,
	version string,
	lang string,
	scriptPath string,
	output string,
	tests Tests,
) Script {
	return createScriptInternally(name, version, lang, scriptPath, output, tests)
}

func createScriptInternally(
	name string,
	version string,
	lang string,
	scriptPath string,
	output string,
	tests Tests,
) Script {
	out := script{
		name:    name,
		version: version,
		lang:    lang,
		script:  scriptPath,
		output:  output,
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

// LanguagePath returns the language path
func (obj *script) LanguagePath() string {
	return obj.lang
}

// ScriptPath returns the script path
func (obj *script) ScriptPath() string {
	return obj.script
}

// Output returns the output
func (obj *script) Output() string {
	return obj.script
}

// HasTests returns true if there is tests, false otherwise
func (obj *script) HasTests() bool {
	return obj.tests != nil
}

// Tests returns the tests, if any
func (obj *script) Tests() Tests {
	return obj.tests
}
