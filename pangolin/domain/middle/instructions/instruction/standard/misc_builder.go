package standard

import "errors"

type miscBuilder struct {
	isConcat bool
}

func createMiscBuilder() MiscBuilder {
	out := miscBuilder{
		isConcat: false,
	}

	return &out
}

// Create initializes the builder
func (app *miscBuilder) Create() MiscBuilder {
	return createMiscBuilder()
}

// IsConcatenation flags the builder as a concatenation
func (app *miscBuilder) IsConcatenation() MiscBuilder {
	app.isConcat = true
	return app
}

// Now builds a new Misc instance
func (app *miscBuilder) Now() (Misc, error) {
	if app.isConcat {
		return createMiscWithConcatenation(), nil
	}

	return nil, errors.New("the Misc is invalid")
}
