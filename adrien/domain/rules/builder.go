package rules

import "errors"

type builder struct {
	list []Rule
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

// WithRules add rules to the builder
func (app *builder) WithRules(rules []Rule) Builder {
	app.list = rules
	return app
}

// Now builds a new Rules instance
func (app *builder) Now() (Rules, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Rule in order to build a Rules instance")
	}

	mp := map[string]Rule{}
	for _, oneRule := range app.list {
		name := oneRule.Name()
		mp[name] = oneRule
	}

	return createRules(app.list, mp), nil
}
