package parsers

type script struct {
	name     string
	version  string
	script   RelativePath
	language RelativePath
}

func createScript(
	name string,
	version string,
	scriptPath RelativePath,
	language RelativePath,
) Script {
	out := script{
		name:     name,
		version:  version,
		script:   scriptPath,
		language: language,
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
