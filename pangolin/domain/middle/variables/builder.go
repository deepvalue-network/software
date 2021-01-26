package variables

import (
	"github.com/steve-care-software/products/pangolin/domain/middle/variables/variable"
)

type builder struct {
	lst []variable.Variable
}

func createBuilder() Builder {
	out := builder{
		lst: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithVariables add variables to the builder
func (app *builder) WithVariables(vrs []variable.Variable) Builder {
	app.lst = vrs
	return app
}

// Now builds a new Variables instance
func (app *builder) Now() (Variables, error) {
	mp := map[string]variable.Variable{}
	if app.lst != nil {
		for _, oneElement := range app.lst {
			mp[oneElement.Name()] = oneElement
		}
	}

	return createVariables(mp), nil
}
