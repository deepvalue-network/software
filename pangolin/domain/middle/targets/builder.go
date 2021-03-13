package targets

import "errors"

type builder struct {
	list []Target
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

// WithTargets add targets to the builder
func (app *builder) WithTargets(targets []Target) Builder {
	app.list = targets
	return app
}

// Now builds new Targets instance
func (app *builder) Now() (Targets, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("the []Target are mandatory in order to build a Targets instance")
	}

	mp := map[string]Target{}
	for _, oneElement := range app.list {
		name := oneElement.Name()
		mp[name] = oneElement
	}

	return createTargets(app.list, mp), nil
}
