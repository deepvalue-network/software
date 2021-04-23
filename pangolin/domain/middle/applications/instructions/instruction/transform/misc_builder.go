package transform

import "errors"

type miscBuilder struct {
	isPop bool
}

func createMiscBuilder() MiscBuilder {
	out := miscBuilder{
		isPop: true,
	}

	return &out
}

// Create initializes the builder
func (app *miscBuilder) Create() MiscBuilder {
	return createMiscBuilder()
}

// IsPop flags the builder as pop
func (app *miscBuilder) IsPop() MiscBuilder {
	app.isPop = true
	return app
}

// Now builds a new Misc instance
func (app *miscBuilder) Now() (Misc, error) {
	if app.isPop {
		return createMiscWithPop(), nil
	}

	return nil, errors.New("the Misc is invalid")
}
