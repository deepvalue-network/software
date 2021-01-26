package parsers

type pushBuilder struct {
	stackframe VariableName
}

func createPushBuilder() PushBuilder {
	out := pushBuilder{
		stackframe: nil,
	}

	return &out
}

// Create initializes the pushBuilder instance
func (app *pushBuilder) Create() PushBuilder {
	return createPushBuilder()
}

// WithStackframe adds a stackFrame to the builder
func (app *pushBuilder) WithStackframe(stackframe VariableName) PushBuilder {
	app.stackframe = stackframe
	return app
}

// Now builds a new Push instance
func (app *pushBuilder) Now() (Push, error) {
	if app.stackframe != nil {
		return createPushWithStackframe(app.stackframe), nil
	}

	return createPush(), nil
}
