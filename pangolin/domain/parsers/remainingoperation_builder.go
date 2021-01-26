package parsers

import "errors"

type remainingOperationBuilder struct {
	first     Identifier
	second    Identifier
	result    VariableName
	remaining VariableName
}

func createRemainingOperationBuilder() RemainingOperationBuilder {
	out := remainingOperationBuilder{
		first:     nil,
		second:    nil,
		result:    nil,
		remaining: nil,
	}

	return &out
}

// Create initializes the builder
func (app *remainingOperationBuilder) Create() RemainingOperationBuilder {
	return createRemainingOperationBuilder()
}

// WithFirst adds the first variable to the builder
func (app *remainingOperationBuilder) WithFirst(first Identifier) RemainingOperationBuilder {
	app.first = first
	return app
}

// WithSecond adds the second variable to the builder
func (app *remainingOperationBuilder) WithSecond(second Identifier) RemainingOperationBuilder {
	app.second = second
	return app
}

// WithResult adds the result variable to the builder
func (app *remainingOperationBuilder) WithResult(result VariableName) RemainingOperationBuilder {
	app.result = result
	return app
}

// WithRemaining adds the remaining variable to the builder
func (app *remainingOperationBuilder) WithRemaining(remaining VariableName) RemainingOperationBuilder {
	app.remaining = remaining
	return app
}

// Now builds a new RemainingOperation instance
func (app *remainingOperationBuilder) Now() (RemainingOperation, error) {
	if app.first == nil {
		return nil, errors.New("the first identifier is mandatory in order to build a RemainingOperation instance")
	}

	if app.second == nil {
		return nil, errors.New("the second identifier is mandatory in order to build a RemainingOperation instance")
	}

	if app.result == nil {
		return nil, errors.New("the result variableName is mandatory in order to build a RemainingOperation instance")
	}

	if app.remaining == nil {
		return nil, errors.New("the remaining variableName is mandatory in order to build a RemainingOperation instance")
	}

	return createRemainingOperation(app.first, app.second, app.result, app.remaining), nil
}
