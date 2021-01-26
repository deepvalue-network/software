package parsers

type readFile struct {
	variable VariableName
	path     RelativePath
}

func createReadFile(
	variable VariableName,
	path RelativePath,
) ReadFile {
	out := readFile{
		variable: variable,
		path:     path,
	}

	return &out
}

// Variable returns the variable name
func (obj *readFile) Variable() VariableName {
	return obj.variable
}

// Path returns the path
func (obj *readFile) Path() RelativePath {
	return obj.path
}
