package rules

import "errors"

type possibilityBuilder struct {
	list   []string
	amount Amount
}

func createPossibilityBuilder() PossibilityBuilder {
	out := possibilityBuilder{
		list:   nil,
		amount: nil,
	}

	return &out
}

// Create initializes the builder
func (app *possibilityBuilder) Create() PossibilityBuilder {
	return createPossibilityBuilder()
}

// WithList adds a list to the builder
func (app *possibilityBuilder) WithList(list []string) PossibilityBuilder {
	app.list = list
	return app
}

// WithAmount adds an amount to the builder
func (app *possibilityBuilder) WithAmount(amount Amount) PossibilityBuilder {
	app.amount = amount
	return app
}

// Now builds a new Possibility instance
func (app *possibilityBuilder) Now() (Possibility, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 possibility in the list in order to build a Possibility instance")
	}

	if app.amount == nil {
		return nil, errors.New("the amount is mandatory in order to build a Possibility instance")
	}

	return createPossibility(app.list, app.amount), nil
}
