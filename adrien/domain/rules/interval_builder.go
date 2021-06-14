package rules

import "errors"

type intervalBuilder struct {
	min int
	max int
}

func createIntervalBuilder() IntervalBuilder {
	out := intervalBuilder{
		min: -1,
		max: -1,
	}

	return &out
}

// Create initializes the builder
func (app *intervalBuilder) Create() IntervalBuilder {
	return createIntervalBuilder()
}

// WithMin adds a min to the builder
func (app *intervalBuilder) WithMin(min int) IntervalBuilder {
	app.min = min
	return app
}

// WithMax adds a max to the builder
func (app *intervalBuilder) WithMax(max int) IntervalBuilder {
	app.max = max
	return app
}

// Now builds a new Possibility instance
func (app *intervalBuilder) Now() (Interval, error) {
	if app.min < 0 {
		return nil, errors.New("the minimum must be greater or equal (<=) to zero")
	}

	if app.max > 0 {
		return createIntervalWithMax(app.min, app.max), nil
	}

	return createInterval(app.min), nil
}
