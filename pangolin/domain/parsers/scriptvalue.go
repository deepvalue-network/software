package parsers

type scriptValue struct {
	name       string
	version    string
	scriptPath RelativePath
	langPath   RelativePath
	output     string
	tests      ScriptTests
}

func createScriptValueWithName(name string) ScriptValue {
	return createScriptValueInternally(name, "", nil, nil, "", nil)
}

func createScriptValueWithVersion(version string) ScriptValue {
	return createScriptValueInternally("", version, nil, nil, "", nil)
}

func createScriptValueWithScriptPath(scriptPath RelativePath) ScriptValue {
	return createScriptValueInternally("", "", scriptPath, nil, "", nil)
}

func createScriptValueWithLanguagePath(langPath RelativePath) ScriptValue {
	return createScriptValueInternally("", "", nil, langPath, "", nil)
}

func createScriptValueWithOutput(output string) ScriptValue {
	return createScriptValueInternally("", "", nil, nil, output, nil)
}

func createScriptValueWithScriptTests(tests ScriptTests) ScriptValue {
	return createScriptValueInternally("", "", nil, nil, "", tests)
}

func createScriptValueInternally(
	name string,
	version string,
	scriptPath RelativePath,
	langPath RelativePath,
	output string,
	tests ScriptTests,
) ScriptValue {
	out := scriptValue{
		name:       name,
		version:    version,
		scriptPath: scriptPath,
		langPath:   langPath,
		output:     output,
		tests:      tests,
	}

	return &out
}

// IsName returns true if there is a name, false otherwise
func (obj *scriptValue) IsName() bool {
	return obj.name != ""
}

// Name returns the name, if any
func (obj *scriptValue) Name() string {
	return obj.name
}

// IsVersion returns true if there is a version, false otherwise
func (obj *scriptValue) IsVersion() bool {
	return obj.version != ""
}

// Version returns the version, if any
func (obj *scriptValue) Version() string {
	return obj.version
}

// IsScript returns true if there is a script, false otherwise
func (obj *scriptValue) IsScript() bool {
	return obj.scriptPath != nil
}

// Script returns the script, if any
func (obj *scriptValue) Script() RelativePath {
	return obj.scriptPath
}

// IsLanguage returns true if there is a language, false otherwise
func (obj *scriptValue) IsLanguage() bool {
	return obj.langPath != nil
}

// Language returns the langPath, if any
func (obj *scriptValue) Language() RelativePath {
	return obj.langPath
}

// IsOutput returns true if there is an output variable, false otherwise
func (obj *scriptValue) IsOutput() bool {
	return obj.output != ""
}

// Output returns the output variable, if any
func (obj *scriptValue) Output() string {
	return obj.output
}

// IsScriptTests returns true if there is scriptTests, false otherwise
func (obj *scriptValue) IsScriptTests() bool {
	return obj.tests != nil
}

// ScriptTests returns the scriptTests, if any
func (obj *scriptValue) ScriptTests() ScriptTests {
	return obj.tests
}
