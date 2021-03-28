package parsers

type readFile struct {
	variable string
	path     RelativePath
}

func createReadFile(
	variable string,
	path RelativePath,
) ReadFile {
	out := readFile{
		variable: variable,
		path:     path,
	}

	return &out
}

// Variable returns the variable name
func (obj *readFile) Variable() string {
	return obj.variable
}

// Path returns the path
func (obj *readFile) Path() RelativePath {
	return obj.path
}
