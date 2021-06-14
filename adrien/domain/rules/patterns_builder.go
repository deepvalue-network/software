package rules

import "errors"

type patternsBuilder struct {
	list []Pattern
}

func createPatternsBuilder() PatternsBuilder {
	out := patternsBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *patternsBuilder) Create() PatternsBuilder {
	return createPatternsBuilder()
}

// WithPatterns add patterns to the builder
func (app *patternsBuilder) WithPatterns(patterns []Pattern) PatternsBuilder {
	app.list = patterns
	return app
}

// Now builds a new Patterns instance
func (app *patternsBuilder) Now() (Patterns, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Pattern in order to build a Patterns instance")
	}

	mp := map[string]Pattern{}
	for _, onePattern := range app.list {
		name := onePattern.Name()
		mp[name] = onePattern
	}

	return createPatterns(app.list, mp), nil
}
