package parsers

import "errors"

type assignmentBuilder struct {
	variable string
	value    ValueRepresentation
}

func createAssignmentBuilder() AssignmentBuilder {
	out := assignmentBuilder{
		variable: "",
		value:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *assignmentBuilder) Create() AssignmentBuilder {
	return createAssignmentBuilder()
}

// WithVariable adds a variable to the builder
func (app *assignmentBuilder) WithVariable(variable string) AssignmentBuilder {
	app.variable = variable
	return app
}

// WithValue adds a value to the builder
func (app *assignmentBuilder) WithValue(value ValueRepresentation) AssignmentBuilder {
	app.value = value
	return app
}

// Now builds a new Assignment instance
func (app *assignmentBuilder) Now() (Assignment, error) {
	if app.variable == "" {
		return nil, errors.New("the variable is mandatory in order to build an Assignment")
	}

	if app.value == nil {
		return nil, errors.New("the valueRepresentation is mandatory in order to build an Assignment")
	}

	return createAssignment(app.variable, app.value), nil
}
