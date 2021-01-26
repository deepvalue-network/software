package middle

type external struct {
	name string
	path string
}

func createExternal(
	name string,
	path string,
) External {
	out := external{
		name: name,
		path: path,
	}

	return &out
}

// Name returns the name
func (obj *external) Name() string {
	return obj.name
}

// Path returns the path
func (obj *external) Path() string {
	return obj.path
}
