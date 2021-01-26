package parsers

import "errors"

type numericValueBuilder struct {
	isNegative bool
	intValue   *int
	floatValue *float64
}

func createNumericValueBuilder() NumericValueBuilder {
	out := numericValueBuilder{
		isNegative: false,
		intValue:   nil,
		floatValue: nil,
	}

	return &out
}

// Create initializes the builder
func (app *numericValueBuilder) Create() NumericValueBuilder {
	return createNumericValueBuilder()
}

// IsNegative flags the builder as negative
func (app *numericValueBuilder) IsNegative() NumericValueBuilder {
	app.isNegative = true
	return app
}

// WithInt adds an int to the builder
func (app *numericValueBuilder) WithInt(intVal int) NumericValueBuilder {
	app.intValue = &intVal
	return app
}

// WithFloat adds a float to the builder
func (app *numericValueBuilder) WithFloat(floatVal float64) NumericValueBuilder {
	app.floatValue = &floatVal
	return app
}

// Now builds a new NumericValue instance
func (app *numericValueBuilder) Now() (NumericValue, error) {
	if app.intValue != nil {
		return createNumericValueWithInt(app.isNegative, app.intValue), nil
	}

	if app.floatValue != nil {
		return createNumericValueWithFloat(app.isNegative, app.floatValue), nil
	}

	return nil, errors.New("the NumericValue is invalid")
}
