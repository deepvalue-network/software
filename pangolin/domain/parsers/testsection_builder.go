package parsers

type testSectionBuilder struct {
	decl []TestDeclaration
}

func createTestSectionBuilder() TestSectionBuilder {
	out := testSectionBuilder{
		decl: nil,
	}

	return &out
}

// Create initializes the builder
func (app *testSectionBuilder) Create() TestSectionBuilder {
	return createTestSectionBuilder()
}

// WithDeclarations add []TestDeclaration to the builder
func (app *testSectionBuilder) WithDeclarations(declarations []TestDeclaration) TestSectionBuilder {
	app.decl = declarations
	return app
}

// Now builds a new TestSection instance
func (app *testSectionBuilder) Now() (TestSection, error) {
	if app.decl == nil {
		app.decl = []TestDeclaration{}
	}

	return createTestSection(app.decl), nil
}
