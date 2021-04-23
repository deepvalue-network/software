package variablename

import "errors"

type miscBuilder struct {
	isPush bool
}

func createMiscBuilder() MiscBuilder {
	out := miscBuilder{
		isPush: true,
	}

	return &out
}

// Create initializes the builder
func (app *miscBuilder) Create() MiscBuilder {
	return createMiscBuilder()
}

// IsPush flags the builder as pop
func (app *miscBuilder) IsPush() MiscBuilder {
	app.isPush = true
	return app
}

// Now builds a new Misc instance
func (app *miscBuilder) Now() (Misc, error) {
	if app.isPush {
		return createMiscWithPop(), nil
	}

	return nil, errors.New("the Misc is invalid")
}
