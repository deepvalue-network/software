package parsers

type scriptValue struct {
	name       string
	version    string
	scriptPath RelativePath
	langPath   RelativePath
}

func createScriptValueWithName(name string) ScriptValue {
	return createScriptValueInternally(name, "", nil, nil)
}

func createScriptValueWithVersion(version string) ScriptValue {
	return createScriptValueInternally("", version, nil, nil)
}

func createScriptValueWithScriptPath(scriptPath RelativePath) ScriptValue {
	return createScriptValueInternally("", "", scriptPath, nil)
}

func createScriptValueWithLanguagePath(langPath RelativePath) ScriptValue {
	return createScriptValueInternally("", "", nil, langPath)
}

func createScriptValueInternally(
	name string,
	version string,
	scriptPath RelativePath,
	langPath RelativePath,
) ScriptValue {
	out := scriptValue{
		name:       name,
		version:    version,
		scriptPath: scriptPath,
		langPath:   langPath,
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
