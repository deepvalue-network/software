package linkers

type script struct {
	language LanguageReference
	name     string
	version  string
	code     string
}

func createScript(
	language LanguageReference,
	name string,
	version string,
	code string,
) Script {
	out := script{
		language: language,
		name:     name,
		version:  version,
		code:     code,
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
