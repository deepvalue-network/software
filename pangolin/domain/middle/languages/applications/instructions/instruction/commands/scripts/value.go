package scripts

import "github.com/deepvalue-network/software/pangolin/domain/parsers"

type value struct {
	name       string
	version    string
	langPath   parsers.RelativePath
	scriptPath parsers.RelativePath
}

func createValueWithName(
	name string,
) Value {
	return createValueInternally(name, "", nil, nil)
}

func createValueWithVersion(
	version string,
) Value {
	return createValueInternally("", version, nil, nil)
}

func createValueWithLanguagePath(
	langPath parsers.RelativePath,
) Value {
	return createValueInternally("", "", langPath, nil)
}

func createValueWithScriptPath(
	scriptPath parsers.RelativePath,
) Value {
	return createValueInternally("", "", nil, scriptPath)
}

func createValueInternally(
	name string,
	version string,
	langPath parsers.RelativePath,
	scriptPath parsers.RelativePath,
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
	return obj.langPath != nil
}

// LanguagePath returns the language path, if any
func (obj *value) LanguagePath() parsers.RelativePath {
	return obj.langPath
}

// IsScriptPath returns true if there is a script path, false otherwise
func (obj *value) IsScriptPath() bool {
	return obj.scriptPath != nil
}

// ScriptPath returns the script path, if any
func (obj *value) ScriptPath() parsers.RelativePath {
	return obj.scriptPath
}
