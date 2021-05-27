package linkers

type external struct {
	name       string
	executable Executable
}

func createExternal(
	name string,
	executable Executable,
) External {
	out := external{
		name:       name,
		executable: executable,
	}

	return &out
}

// Name returns the name
func (obj *external) Name() string {
	return obj.name
}

// Executable returns the executable
func (obj *external) Executable() Executable {
	return obj.executable
}
