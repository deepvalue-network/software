package remaining

import "errors"

type arythmeticBuilder struct {
	isDiv bool
}

func createArythmeticBuilder() ArythmeticBuilder {
	out := arythmeticBuilder{
		isDiv: false,
	}

	return &out
}

// Create initializes the builder
func (app *arythmeticBuilder) Create() ArythmeticBuilder {
	return createArythmeticBuilder()
}

// IsDiv flags the builder as a division
func (app *arythmeticBuilder) IsDiv() ArythmeticBuilder {
	app.isDiv = true
	return app
}

// Now builds a new Arythmetic instance
func (app *arythmeticBuilder) Now() (Arythmetic, error) {
	if app.isDiv {
		return createArythmeticWithDiv(), nil
	}

	return nil, errors.New("the Arythmetic is invalid")
}
