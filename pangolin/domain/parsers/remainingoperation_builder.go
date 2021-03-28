package parsers

import "errors"

type remainingOperationBuilder struct {
	first     string
	second    string
	result    string
	remaining string
}

func createRemainingOperationBuilder() RemainingOperationBuilder {
	out := remainingOperationBuilder{
		first:     "",
		second:    "",
		result:    "",
		remaining: "",
	}

	return &out
}

// Create initializes the builder
func (app *remainingOperationBuilder) Create() RemainingOperationBuilder {
	return createRemainingOperationBuilder()
}

// WithFirst adds the first variable to the builder
func (app *remainingOperationBuilder) WithFirst(first string) RemainingOperationBuilder {
	app.first = first
	return app
}

// WithSecond adds the second variable to the builder
func (app *remainingOperationBuilder) WithSecond(second string) RemainingOperationBuilder {
	app.second = second
	return app
}

// WithResult adds the result variable to the builder
func (app *remainingOperationBuilder) WithResult(result string) RemainingOperationBuilder {
	app.result = result
	return app
}

// WithRemaining adds the remaining variable to the builder
func (app *remainingOperationBuilder) WithRemaining(remaining string) RemainingOperationBuilder {
	app.remaining = remaining
	return app
}

// Now builds a new RemainingOperation instance
func (app *remainingOperationBuilder) Now() (RemainingOperation, error) {
	if app.first == "" {
		return nil, errors.New("the first variableName is mandatory in order to build a RemainingOperation instance")
	}

	if app.second == "" {
		return nil, errors.New("the second variableName is mandatory in order to build a RemainingOperation instance")
	}

	if app.result == "" {
		return nil, errors.New("the result variableName is mandatory in order to build a RemainingOperation instance")
	}

	if app.remaining == "" {
		return nil, errors.New("the remaining variableName is mandatory in order to build a RemainingOperation instance")
	}

	return createRemainingOperation(app.first, app.second, app.result, app.remaining), nil
}
