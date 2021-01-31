package overviews

type builder struct {
	valid      []ValidTransaction
	invalid    []InvalidTransaction
	canBeSaved bool
}

func createBuilder() Builder {
	out := builder{
		valid:      nil,
		invalid:    nil,
		canBeSaved: false,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithValid adds valid transactions to the builder
func (app *builder) WithValid(valid []ValidTransaction) Builder {
	app.valid = valid
	return app
}

// WithInvalid adds invalid transactions to the builder
func (app *builder) WithInvalid(invalid []InvalidTransaction) Builder {
	app.invalid = invalid
	return app
}

// CanBeSaved flags the builder as canBeSaved
func (app *builder) CanBeSaved() Builder {
	app.canBeSaved = true
	return app
}

// Now builds a new Overview instance
func (app *builder) Now() (Overview, error) {
	if app.valid == nil {
		app.valid = []ValidTransaction{}
	}

	if app.invalid == nil {
		app.invalid = []InvalidTransaction{}
	}

	return createOverview(app.valid, app.invalid, app.canBeSaved), nil
}
