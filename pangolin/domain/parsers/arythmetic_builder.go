package parsers

import "errors"

type arythmeticBuilder struct {
	add StandardOperation
	sub StandardOperation
	mul StandardOperation
	div RemainingOperation
}

func createArythmeticBuilder() ArythmeticBuilder {
	out := arythmeticBuilder{
		add: nil,
		sub: nil,
		mul: nil,
		div: nil,
	}

	return &out
}

// Create initializes the builder
func (app *arythmeticBuilder) Create() ArythmeticBuilder {
	return createArythmeticBuilder()
}

// WithAddition adds an addition to the builder
func (app *arythmeticBuilder) WithAddition(add StandardOperation) ArythmeticBuilder {
	app.add = add
	return app
}

// WithSubstraction adds a substraction to the builder
func (app *arythmeticBuilder) WithSubstraction(sub StandardOperation) ArythmeticBuilder {
	app.sub = sub
	return app
}

// WithMultiplication adds a multiplication to the builder
func (app *arythmeticBuilder) WithMultiplication(mul StandardOperation) ArythmeticBuilder {
	app.mul = mul
	return app
}

// WithDivision adds a division to the builder
func (app *arythmeticBuilder) WithDivision(div RemainingOperation) ArythmeticBuilder {
	app.div = div
	return app
}

// Now builds a new Arythmetic instance
func (app *arythmeticBuilder) Now() (Arythmetic, error) {
	if app.add != nil {
		return createArythmeticWithAddition(app.add), nil
	}

	if app.sub != nil {
		return createArythmeticWithSubstraction(app.sub), nil
	}

	if app.mul != nil {
		return createArythmeticWithMultiplication(app.mul), nil
	}

	if app.div != nil {
		return createArythmeticWithDivision(app.div), nil
	}

	return nil, errors.New("the Arythmetic is invalid")
}
