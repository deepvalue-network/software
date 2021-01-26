package parsers

type popBuilder struct {
	stackframe TransformOperation
}

func createPopBuilder() PopBuilder {
	out := popBuilder{
		stackframe: nil,
	}

	return &out
}

// Create initializes the popBuilder instance
func (app *popBuilder) Create() PopBuilder {
	return createPopBuilder()
}

// WithStackframe adds a stackFrame to the builder
func (app *popBuilder) WithStackframe(stackframe TransformOperation) PopBuilder {
	app.stackframe = stackframe
	return app
}

// Now builds a new Pop instance
func (app *popBuilder) Now() (Pop, error) {
	if app.stackframe != nil {
		return createPopWithStackframe(app.stackframe), nil
	}

	return createPop(), nil
}
