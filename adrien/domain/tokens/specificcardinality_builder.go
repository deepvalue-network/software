package tokens

import "errors"

type specificCardinalityBuilder struct {
	amount *uint
	rnge   Range
}

func createSpecificCardinalityBuilder() SpecificCardinalityBuilder {
	out := specificCardinalityBuilder{
		amount: nil,
		rnge:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *specificCardinalityBuilder) Create() SpecificCardinalityBuilder {
	return createSpecificCardinalityBuilder()
}

// WithAmount adds an amount to the builder
func (app *specificCardinalityBuilder) WithAmount(amount uint) SpecificCardinalityBuilder {
	app.amount = &amount
	return app
}

// WithRange adds a range to the builder
func (app *specificCardinalityBuilder) WithRange(rnge Range) SpecificCardinalityBuilder {
	app.rnge = rnge
	return app
}

// Now builds a new SpecificCardinality instance
func (app *specificCardinalityBuilder) Now() (SpecificCardinality, error) {
	if app.amount != nil {
		return createSpecificCardinalityWithAmount(app.amount), nil
	}

	if app.rnge != nil {
		return createSpecificCardinalityWithRange(app.rnge), nil
	}

	return nil, errors.New("the SpecificCardinality instance is invalid")
}
