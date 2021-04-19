package parsers

import "errors"

type valueBuilder struct {
	isNil    bool
	numeric  NumericValue
	bl       *bool
	strValue string
}

func createValueBuilder() ValueBuilder {
	out := valueBuilder{
		isNil:    false,
		numeric:  nil,
		bl:       nil,
		strValue: "",
	}

	return &out
}

// Create initializes the builder
func (app *valueBuilder) Create() ValueBuilder {
	return createValueBuilder()
}

// IsNil flags the builder as nil
func (app *valueBuilder) IsNil() ValueBuilder {
	app.isNil = true
	return app
}

// WithNumeric adds a numeric value to the builder
func (app *valueBuilder) WithNumeric(numeric NumericValue) ValueBuilder {
	app.numeric = numeric
	return app
}

// WithBool adds a bool value to the builder
func (app *valueBuilder) WithBool(bl bool) ValueBuilder {
	app.bl = &bl
	return app
}

// WithString adds a string value to the builder
func (app *valueBuilder) WithString(str string) ValueBuilder {
	app.strValue = str
	return app
}

// Now builds a new Value instance
func (app *valueBuilder) Now() (Value, error) {
	if app.isNil {
		return createValueWithNil(), nil
	}

	if app.numeric != nil {
		return createValueWithNumeric(app.numeric), nil
	}

	if app.bl != nil {
		return createValueWithBool(app.bl), nil
	}

	if app.strValue != "" {
		return createValueWithString(app.strValue), nil
	}

	return nil, errors.New("the Value is invalid")
}
