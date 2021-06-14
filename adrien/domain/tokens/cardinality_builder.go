package tokens

import "errors"

type cardinalityBuilder struct {
	isNonZeroMultiple bool
	isZeroMultiple    bool
	specific          SpecificCardinality
}

func createCardinalityBuilder() CardinalityBuilder {
	out := cardinalityBuilder{
		isNonZeroMultiple: false,
		isZeroMultiple:    false,
		specific:          nil,
	}

	return &out
}

// Create initializes the builder
func (app *cardinalityBuilder) Create() CardinalityBuilder {
	return createCardinalityBuilder()
}

// IsNonZeroMultiple flags the builder as a non-zero multiple
func (app *cardinalityBuilder) IsNonZeroMultiple() CardinalityBuilder {
	app.isNonZeroMultiple = true
	return app
}

// IsZeroMultiple flags the builder as a zero multiple
func (app *cardinalityBuilder) IsZeroMultiple() CardinalityBuilder {
	app.isZeroMultiple = true
	return app
}

// WithSpecific adds a specific cardinality to the builder
func (app *cardinalityBuilder) WithSpecific(specific SpecificCardinality) CardinalityBuilder {
	app.specific = specific
	return app
}

// Now builds a new Cardinality instance
func (app *cardinalityBuilder) Now() (Cardinality, error) {
	if app.isNonZeroMultiple {
		return createCardinalityWithNonZeroMultiple(), nil
	}

	if app.isZeroMultiple {
		return createCardinalityWithZeroMultiple(), nil
	}

	if app.specific != nil {
		return createCardinalityWithSpecific(app.specific), nil
	}

	return nil, errors.New("the cardinality instance is invalid")
}
