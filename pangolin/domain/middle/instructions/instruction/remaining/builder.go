package remaining

import "errors"

type builder struct {
	arythmeticBuilder ArythmeticBuilder
	miscBuilder       MiscBuilder
	operationBuilder  OperationBuilder
	result            string
	remaining         string
	first             string
	second            string
	operation         Operation
	isDiv             bool
	isMatch           bool
}

func createBuilder(
	arythmeticBuilder ArythmeticBuilder,
	miscBuilder MiscBuilder,
	operationBuilder OperationBuilder,
) Builder {
	out := builder{
		arythmeticBuilder: arythmeticBuilder,
		miscBuilder:       miscBuilder,
		operationBuilder:  operationBuilder,
		result:            "",
		remaining:         "",
		first:             "",
		second:            "",
		operation:         nil,
		isDiv:             false,
		isMatch:           false,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.arythmeticBuilder, app.miscBuilder, app.operationBuilder)
}

// WithResult add result to the builder
func (app *builder) WithResult(result string) Builder {
	app.result = result
	return app
}

// WithRemaining add remaining to the builder
func (app *builder) WithRemaining(remaining string) Builder {
	app.remaining = remaining
	return app
}

// WithFirst add first to the builder
func (app *builder) WithFirst(first string) Builder {
	app.first = first
	return app
}

// WithSecond add second to the builder
func (app *builder) WithSecond(second string) Builder {
	app.second = second
	return app
}

// WithOperation adds an operation to the builder
func (app *builder) WithOperation(operation Operation) Builder {
	app.operation = operation
	return app
}

// IsDiv flags the builder as a division
func (app *builder) IsDiv() Builder {
	app.isDiv = true
	return app
}

// IsMatch flags the builder as a match
func (app *builder) IsMatch() Builder {
	app.isMatch = true
	return app
}

// Now builds a new Remaining instance
func (app *builder) Now() (Remaining, error) {
	if app.result == "" {
		return nil, errors.New("the result is mandatory in order to build a Remaining instance")
	}

	if app.remaining == "" {
		return nil, errors.New("the remaining is mandatory in order to build a Remaining instance")
	}

	if app.first == "" {
		return nil, errors.New("the first is mandatory in order to build a Remaining instance")
	}

	if app.second == "" {
		return nil, errors.New("the second is mandatory in order to build a Remaining instance")
	}

	if app.operation != nil {
		return createRemaining(app.operation, app.result, app.remaining, app.first, app.second), nil
	}

	operationBuilder := app.operationBuilder.Create()
	if app.isDiv {
		ary, err := app.arythmeticBuilder.Create().IsDiv().Now()
		if err != nil {
			return nil, err
		}

		operationBuilder.WithArythmetic(ary)
	}

	if app.isMatch {
		misc, err := app.miscBuilder.Create().IsMatch().Now()
		if err != nil {
			return nil, err
		}

		operationBuilder.WithMisc(misc)
	}

	op, err := operationBuilder.Now()
	if err != nil {
		return nil, err
	}

	return createRemaining(op, app.result, app.remaining, app.first, app.second), nil
}
