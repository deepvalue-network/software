package tokens

import "errors"

type blockBuilder struct {
	must Lines
	not  Lines
}

func createBlockBuilder() BlockBuilder {
	out := blockBuilder{
		must: nil,
		not:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *blockBuilder) Create() BlockBuilder {
	return createBlockBuilder()
}

// WithMust add must lines to the builder
func (app *blockBuilder) WithMust(must Lines) BlockBuilder {
	app.must = must
	return app
}

// WithNot add not lines to the builder
func (app *blockBuilder) WithNot(not Lines) BlockBuilder {
	app.not = not
	return app
}

// Now builds a new Block instance
func (app *blockBuilder) Now() (Block, error) {
	if app.must == nil {
		return nil, errors.New("the must lines are mandatory in order to build a Block instance")
	}

	if app.not != nil {
		return createBlockWithNot(app.must, app.not), nil
	}

	return createBlock(app.must), nil
}
