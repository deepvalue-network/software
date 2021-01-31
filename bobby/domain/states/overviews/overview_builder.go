package overviews

type overviewBuilder struct {
	valid      []ValidTransaction
	invalid    []InvalidTransaction
	canBeSaved bool
}

func createOverviewBuilder() OverviewBuilder {
	out := overviewBuilder{
		valid:      nil,
		invalid:    nil,
		canBeSaved: false,
	}

	return &out
}

// Create initializes the builder
func (app *overviewBuilder) Create() OverviewBuilder {
	return createOverviewBuilder()
}

// WithValid adds valid transactions to the builder
func (app *overviewBuilder) WithValid(valid []ValidTransaction) OverviewBuilder {
	app.valid = valid
	return app
}

// WithInvalid adds invalid transactions to the builder
func (app *overviewBuilder) WithInvalid(invalid []InvalidTransaction) OverviewBuilder {
	app.invalid = invalid
	return app
}

// CanBeSaved flags the builder as canBeSaved
func (app *overviewBuilder) CanBeSaved() OverviewBuilder {
	app.canBeSaved = true
	return app
}

// Now builds a new Overview instance
func (app *overviewBuilder) Now() (Overview, error) {
	if app.valid == nil {
		app.valid = []ValidTransaction{}
	}

	if app.invalid == nil {
		app.invalid = []InvalidTransaction{}
	}

	return createOverview(app.valid, app.invalid, app.canBeSaved), nil
}
