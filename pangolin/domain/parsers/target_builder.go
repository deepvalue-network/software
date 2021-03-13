package parsers

import "errors"

type targetBuilder struct {
	name    string
	singles []TargetSingle
}

func createTargetBuilder() TargetBuilder {
	out := targetBuilder{
		name:    "",
		singles: nil,
	}

	return &out
}

// Create initializes the builder
func (app *targetBuilder) Create() TargetBuilder {
	return createTargetBuilder()
}

// WithName adds a name to the builder
func (app *targetBuilder) WithName(name string) TargetBuilder {
	app.name = name
	return app
}

// WithSingles adds single targets to the builder
func (app *targetBuilder) WithSingles(singles []TargetSingle) TargetBuilder {
	app.singles = singles
	return app
}

// Now builds a new Target instance
func (app *targetBuilder) Now() (Target, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Target instance")
	}

	if app.singles == nil {
		app.singles = []TargetSingle{}
	}

	var path RelativePath
	events := []Event{}

	for _, oneSingle := range app.singles {
		if oneSingle.IsEvents() {
			events = oneSingle.Events()
			continue
		}

		if oneSingle.IsPath() {
			path = oneSingle.Path()
			continue
		}
	}

	if events == nil {
		return nil, errors.New("the events are mandatory in order to build a Target instance")
	}

	if path == nil {
		return nil, errors.New("the path is mandatory in order to build a Target instance")
	}

	return createTarget(app.name, path, events), nil

}
