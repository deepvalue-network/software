package parsers

import "errors"

type printBuilder struct {
	value Value
}

func createPrintBuilder() PrintBuilder {
	out := printBuilder{
		value: nil,
	}

	return &out
}

// Create initializes the builder
func (app *printBuilder) Create() PrintBuilder {
	return createPrintBuilder()
}

// WithValue adds value to the builder
func (app *printBuilder) WithValue(value Value) PrintBuilder {
	app.value = value
	return app
}

// Now builds a new Print instance
func (app *printBuilder) Now() (Print, error) {
	if app.value == nil {
		return nil, errors.New("the value is mandatory in order to build a Print instance")
	}

	return createPrint(app.value), nil
}
