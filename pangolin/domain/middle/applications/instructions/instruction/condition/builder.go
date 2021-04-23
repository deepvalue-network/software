package condition

type builder struct {
	operationBuilder OperationBuilder
	proposition      Proposition
	isJump           bool
}

func createBuilder(operationBuilder OperationBuilder) Builder {
	out := builder{
		operationBuilder: operationBuilder,
		proposition:      nil,
		isJump:           false,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.operationBuilder)
}

// WithProposition adds a proposition to the builder
func (app *builder) WithProposition(proposition Proposition) Builder {
	app.proposition = proposition
	return app
}

// IsJump flags the builder with a isJump operation
func (app *builder) IsJump() Builder {
	app.isJump = true
	return app
}

// Now builds a new Condition instance
func (app *builder) Now() (Condition, error) {
	operationBuilder := app.operationBuilder.Create()
	if app.isJump {
		operationBuilder.IsJump()
	}

	operation, err := operationBuilder.Now()
	if err != nil {
		return nil, err
	}

	return createCondition(app.proposition, operation), nil
}
