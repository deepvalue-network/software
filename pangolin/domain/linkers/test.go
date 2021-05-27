package linkers

type test struct {
	name       string
	executable Executable
}

func createTest(
	name string,
	executable Executable,
) Test {
	out := test{
		name:       name,
		executable: executable,
	}

	return &out
}

// Name returns the name
func (obj *test) Name() string {
	return obj.name
}

// Executable returns the executable
func (obj *test) Executable() Executable {
	return obj.executable
}
