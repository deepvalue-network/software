package scripts

type value struct {
	name       string
	version    string
	langPath   string
	scriptPath string
}

func createValueWithName(
	name string,
) Value {
	return createValueInternally(name, "", "", "")
}

func createValueWithVersion(
	version string,
) Value {
	return createValueInternally("", version, "", "")
}

func createValueWithLanguagePath(
	langPath string,
) Value {
	return createValueInternally("", "", langPath, "")
}

func createValueWithScriptPath(
	scriptPath string,
) Value {
	return createValueInternally("", "", "", scriptPath)
}

func createValueInternally(
	name string,
	version string,
	langPath string,
	scriptPath string,
) Value {
	out := value{
		name:       name,
		version:    version,
		langPath:   langPath,
		scriptPath: scriptPath,
	}

	return &out
}

// IsName returns true if there is a name, false otherwise
func (obj *value) IsName() bool {
	return obj.name != ""
}

// Name returns the name, if any
func (obj *value) Name() string {
	return obj.name
}

// IsVersion returns true if there is a version, false otherwise
func (obj *value) IsVersion() bool {
	return obj.version != ""
}

// Version returns the version, if any
func (obj *value) Version() string {
	return obj.version
}

// IsLanguagePath returns true if there is a language path, false otherwise
func (obj *value) IsLanguagePath() bool {
	return obj.langPath != ""
}

// LanguagePath returns the language path, if any
func (obj *value) LanguagePath() string {
	return obj.langPath
}

// IsScriptPath returns true if there is a script path, false otherwise
func (obj *value) IsScriptPath() bool {
	return obj.scriptPath != ""
}

// ScriptPath returns the script path, if any
func (obj *value) ScriptPath() string {
	return obj.scriptPath
}
