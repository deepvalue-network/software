package rules

import "errors"

type amountBuilder struct {
	exactly  int
	interval Interval
}

func createAmountBuilder() AmountBuilder {
	out := amountBuilder{
		exactly:  -1,
		interval: nil,
	}

	return &out
}

// Create initializes the builder
func (app *amountBuilder) Create() AmountBuilder {
	return createAmountBuilder()
}

// WithExactly adds an exact amount to the builder
func (app *amountBuilder) WithExactly(exactly int) AmountBuilder {
	app.exactly = exactly
	return app
}

// WithInterval adds an interval to the builder
func (app *amountBuilder) WithInterval(interval Interval) AmountBuilder {
	app.interval = interval
	return app
}

// Now builds a new Amount instance
func (app *amountBuilder) Now() (Amount, error) {
	if app.exactly != -1 {
		return createAmountWithExactly(app.exactly), nil
	}

	if app.interval != nil {
		return createAmountWithInterval(app.interval), nil
	}

	return nil, errors.New("the amount is invalid")
}
