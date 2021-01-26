package labels

import (
	"github.com/steve-care-software/products/pangolin/domain/middle/labels/label"
)

type builder struct {
	lst []label.Label
	mp  map[string]label.Label
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
func (app *builder) WithList(lst []label.Label) Builder {
	app.lst = lst
	return app
}

// WithMap add map to the builder
func (app *builder) WithMap(mp map[string]label.Label) Builder {
	app.mp = mp
	return app
}

// Now builds a new Labels instance
func (app *builder) Now() (Labels, error) {
	if app.mp != nil {
		lst := []label.Label{}
		for _, oneLabel := range app.mp {
			lst = append(lst, oneLabel)
		}

		app.lst = lst
	}

	if app.lst == nil {
		app.lst = []label.Label{}
	}

	return createLabels(app.lst), nil
}
