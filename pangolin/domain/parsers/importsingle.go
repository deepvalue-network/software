package parsers

type importSingle struct {
	name string
	path RelativePath
}

func createImportSingle(
	name string,
	path RelativePath,
) ImportSingle {
	out := importSingle{
		name: name,
		path: path,
	}

	return &out
}

// Name returns the name
func (obj *importSingle) Name() string {
	return obj.name
}

// Path returns the path
func (obj *importSingle) Path() RelativePath {
	return obj.path
}
