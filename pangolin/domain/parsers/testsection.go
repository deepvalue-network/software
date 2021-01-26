package parsers

type testSection struct {
	decl []TestDeclaration
}

func createTestSection(decl []TestDeclaration) TestSection {
	out := testSection{
		decl: decl,
	}

	return &out
}

// Declarations return the test declarations
func (obj *testSection) Declarations() []TestDeclaration {
	return obj.decl
}
