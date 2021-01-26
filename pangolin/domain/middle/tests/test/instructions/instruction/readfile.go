package instruction

type readFile struct {
	variable string
	path     string
}

func createReadFile(
	variable string,
	path string,
) ReadFile {
	out := readFile{
		variable: variable,
		path:     path,
	}

	return &out
}

// Variable returns the variable
func (obj *readFile) Variable() string {
	return obj.variable
}

// Path returns the path
func (obj *readFile) Path() string {
	return obj.path
}
