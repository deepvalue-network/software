package parsers

import "errors"

type frameAssignmentBuilder struct {
	standard StandardOperation
}

func createFrameAssignmentBuilder() FrameAssignmentBuiler {
	out := frameAssignmentBuilder{
		standard: nil,
	}

	return &out
}

// Create initializes the builder
func (app *frameAssignmentBuilder) Create() FrameAssignmentBuiler {
	return createFrameAssignmentBuilder()
}

// WithStandard adds a standardOperation to the builder
func (app *frameAssignmentBuilder) WithStandard(standard StandardOperation) FrameAssignmentBuiler {
	app.standard = standard
	return app
}

// Now builds a new FrameAssignment instance
func (app *frameAssignmentBuilder) Now() (FrameAssignment, error) {
	if app.standard == nil {
		return nil, errors.New("the standardOperation is mandatory in order to build a FrameAssignment instance")
	}

	return createFrameAssignment(app.standard), nil
}
