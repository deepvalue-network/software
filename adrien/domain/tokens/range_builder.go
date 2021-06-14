package tokens

type rangeBuilder struct {
	min uint
	max *uint
}

func createRangeBuilder() RangeBuilder {
	out := rangeBuilder{
		min: 0,
		max: nil,
	}

	return &out
}

// Create initializes the builder
func (app *rangeBuilder) Create() RangeBuilder {
	return createRangeBuilder()
}

// WithMinimum adds a minimum to the builder
func (app *rangeBuilder) WithMinimum(min uint) RangeBuilder {
	app.min = min
	return app
}

// WithMaximum adds a maximum to the builder
func (app *rangeBuilder) WithMaximum(max uint) RangeBuilder {
	app.max = &max
	return app
}

// Now builds a new Range instance
func (app *rangeBuilder) Now() (Range, error) {
	if app.max != nil {
		return createRangeWithMaximum(app.min, app.max), nil
	}

	return createRange(app.min), nil
}
