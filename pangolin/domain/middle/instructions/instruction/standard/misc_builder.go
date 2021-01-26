package standard

import "errors"

type miscBuilder struct {
	isConcat          bool
	isFrameAssignment bool
}

func createMiscBuilder() MiscBuilder {
	out := miscBuilder{
		isConcat:          false,
		isFrameAssignment: false,
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

// IsFrameAssignment flags the builder as a frameAssignment
func (app *miscBuilder) IsFrameAssignment() MiscBuilder {
	app.isFrameAssignment = true
	return app
}

// Now builds a new Misc instance
func (app *miscBuilder) Now() (Misc, error) {
	if app.isConcat {
		return createMiscWithConcatenation(), nil
	}

	if app.isFrameAssignment {
		return createMiscWithFrameAssignment(), nil
	}

	return nil, errors.New("the Misc is invalid")
}
