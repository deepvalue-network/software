package tests

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/tests/test"
)

type builder struct {
	lst []test.Test
	mp  map[string]test.Test
}

func createBuilder() Builder {
	out := builder{
		lst: nil,
		mp:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithList add list to the builder
func (app *builder) WithList(lst []test.Test) Builder {
	app.lst = lst
	return app
}

// WithMap add map to the builder
func (app *builder) WithMap(mp map[string]test.Test) Builder {
	app.mp = mp
	return app
}

// Now builds a new Tests instance
func (app *builder) Now() (Tests, error) {
	if app.mp != nil {
		lst := []test.Test{}
		for _, oneLabel := range app.mp {
			lst = append(lst, oneLabel)
		}

		app.lst = lst
	}

	if app.lst == nil {
		app.lst = []test.Test{}
	}

	return createTests(app.lst), nil
}
