package middle

type script struct {
	name    string
	version string
	lang    string
	script  string
}

func createScript(
	name string,
	version string,
	lang string,
	scriptPath string,
) Script {
	out := script{
		name:    name,
		version: version,
		lang:    lang,
		script:  scriptPath,
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

// script returns the script path
func (obj *script) ScriptPath() string {
	return obj.script
}
