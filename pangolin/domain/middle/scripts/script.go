package scripts

type script struct {
	name    string
	version string
	lang    string
	script  string
	output  string
}

func createScript(
	name string,
	version string,
	lang string,
	scriptPath string,
	output string,
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
