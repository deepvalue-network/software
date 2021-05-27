package scripts

type test struct {
	name string
	path string
}

func createTest(
	name string,
	path string,
) Test {
	out := test{
		name: name,
		path: path,
	}

	return &out
}

// Name returns the name
func (obj *test) Name() string {
	return obj.name
}

// Path returns the path
func (obj *test) Path() string {
	return obj.path
}
