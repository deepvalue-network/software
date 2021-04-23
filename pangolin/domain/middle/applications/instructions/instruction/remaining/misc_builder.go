package remaining

import "errors"

type miscBuilder struct {
	isMatch bool
}

func createMiscBuilder() MiscBuilder {
	out := miscBuilder{
		isMatch: false,
	}

	return &out
}

// Create initializes the builder
func (app *miscBuilder) Create() MiscBuilder {
	return createMiscBuilder()
}

// IsMatch flags the builder as a match
func (app *miscBuilder) IsMatch() MiscBuilder {
	app.isMatch = true
	return app
}

// Now builds a new Misc instance
func (app *miscBuilder) Now() (Misc, error) {
	if app.isMatch {
		return createMiscWithMatch(), nil
	}

	return nil, errors.New("the Misc is invalid")
}
