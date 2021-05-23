package parsers

type scriptTest struct {
	name string
	path RelativePath
}

func createScriptTest(
	name string,
	path RelativePath,
) ScriptTest {
	out := scriptTest{
		name: name,
		path: path,
	}

	return &out
}

// Name returns the name
func (obj *scriptTest) Name() string {
	return obj.name
}

// Path returns the path
func (obj *scriptTest) Path() RelativePath {
	return obj.path
}
