package linkers

type test struct {
	name   string
	script Script
}

func createTest(
	name string,
	script Script,
) Test {
	out := test{
		name:   name,
		script: script,
	}

	return &out
}

// Name returns the name
func (obj *test) Name() string {
	return obj.name
}

// Script returns the script
func (obj *test) Script() Script {
	return obj.script
}
