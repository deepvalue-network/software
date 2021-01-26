package condition

import "errors"

type propositionBuilder struct {
	name      string
	condition string
}

func createPropositionBuilder() PropositionBuilder {
	out := propositionBuilder{
		name:      "",
		condition: "",
	}

	return &out
}

// Create initializes the propositionBuilder
func (app *propositionBuilder) Create() PropositionBuilder {
	return createPropositionBuilder()
}

// WithName adds a name to the propositionBuilder
func (app *propositionBuilder) WithName(name string) PropositionBuilder {
	app.name = name
	return app
}

// WithCondition adds a condition to the propositionBuilder
func (app *propositionBuilder) WithCondition(condition string) PropositionBuilder {
	app.condition = condition
	return app
}

// Now builds a new Proposition instance
func (app *propositionBuilder) Now() (Proposition, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Proposition instance")
	}

	if app.condition != "" {
		return createPropositionWithCondition(app.name, app.condition), nil
	}

	return createProposition(app.name), nil
}
