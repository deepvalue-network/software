package tests

import (
	"errors"

	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/tests/test"
)

type builder struct {
	list []test.Test
}

func createBuilder() Builder {
	out := builder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithList add tests to the builder
func (app *builder) WithList(list []test.Test) Builder {
	app.list = list
	return app
}

// Now builds a new Instructins instance
func (app *builder) Now() (Tests, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("the []Test are mandatory in order to build an Tests instance")
	}

	return createTests(app.list), nil
}
