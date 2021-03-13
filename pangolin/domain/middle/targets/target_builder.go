package targets

import "errors"

type targetBuilder struct {
	name   string
	path   string
	events []Event
}

func createTargetBuilder() TargetBuilder {
	out := targetBuilder{
		name:   "",
		path:   "",
		events: nil,
	}

	return &out
}

// Create initializes the targetBuilder
func (app *targetBuilder) Create() TargetBuilder {
	return createTargetBuilder()
}

// WithName adds a name to the targetBuilder
func (app *targetBuilder) WithName(name string) TargetBuilder {
	app.name = name
	return app
}

// WithPath adds a path to the targetBuilder
func (app *targetBuilder) WithPath(path string) TargetBuilder {
	app.path = path
	return app
}

// WithEvents add events to the targetBuilder
func (app *targetBuilder) WithEvents(events []Event) TargetBuilder {
	app.events = events
	return app
}

// Now builds a new Target instance
func (app *targetBuilder) Now() (Target, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Target instance")
	}

	if app.path == "" {
		return nil, errors.New("the path is mandatory in order to build a Target instance")
	}

	if app.events != nil && len(app.events) <= 0 {
		return nil, errors.New("the events are mandatory in order to build a Target instance")
	}

	return createTarget(app.name, app.path, app.events), nil
}
