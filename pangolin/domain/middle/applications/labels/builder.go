package labels

import (
	"errors"

	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/labels/label"
)

type builder struct {
	list []label.Label
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

// WithList add labels to the builder
func (app *builder) WithList(list []label.Label) Builder {
	app.list = list
	return app
}

// Now builds a new Instructins instance
func (app *builder) Now() (Labels, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("the []Label are mandatory in order to build an Labels instance")
	}

	mp := map[string]label.Label{}
	for _, oneLabel := range app.list {
		name := oneLabel.Name()
		mp[name] = oneLabel
	}

	return createLabels(mp, app.list), nil
}
