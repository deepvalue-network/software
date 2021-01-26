package parsers

import "errors"

type variableBuilder struct {
	declaration Declaration
	assignment  Assignment
	concat      Concatenation
	delete      VariableName
}

func createVariableBuilder() VariableBuilder {
	out := variableBuilder{
		declaration: nil,
		assignment:  nil,
		concat:      nil,
		delete:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *variableBuilder) Create() VariableBuilder {
	return createVariableBuilder()
}

// WithDeclaration adds a declaration to the builder
func (app *variableBuilder) WithDeclaration(declaration Declaration) VariableBuilder {
	app.declaration = declaration
	return app
}

// WithAssigment adds an assignment to the builder
func (app *variableBuilder) WithAssigment(assignment Assignment) VariableBuilder {
	app.assignment = assignment
	return app
}

// WithConcatenation adds a concatenation to the builder
func (app *variableBuilder) WithConcatenation(concatenation Concatenation) VariableBuilder {
	app.concat = concatenation
	return app
}

// WithDelete adds a delete to the builder
func (app *variableBuilder) WithDelete(delete VariableName) VariableBuilder {
	app.delete = delete
	return app
}

// Now builds a new Variable instance
func (app *variableBuilder) Now() (Variable, error) {
	if app.declaration != nil {
		return createVariableWithDeclaration(app.declaration), nil
	}

	if app.assignment != nil {
		return createVariableWithAssignment(app.assignment), nil
	}

	if app.concat != nil {
		return createVariableWithConcatenation(app.concat), nil
	}

	if app.delete != nil {
		return createVariableWithDelete(app.delete), nil
	}

	return nil, errors.New("the Variable is invalid")
}
