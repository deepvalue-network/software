package parsers

type scriptTests struct {
	list []ScriptTest
}

func createScriptTests(
	list []ScriptTest,
) ScriptTests {
	out := scriptTests{
		list: list,
	}

	return &out
}

// All returns the script tests
func (obj *scriptTests) All() []ScriptTest {
	return obj.list
}
