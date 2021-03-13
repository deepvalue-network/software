package parsers

import "errors"

type targetSingleBuilder struct {
	evts []Event
	path RelativePath
}

func createTargetSingleBuilder() TargetSingleBuilder {
	out := targetSingleBuilder{
		evts: nil,
		path: nil,
	}

	return &out
}

// Create initializes the builder
func (app *targetSingleBuilder) Create() TargetSingleBuilder {
	return createTargetSingleBuilder()
}

// WithEvents adds events to the builder
func (app *targetSingleBuilder) WithEvents(evts []Event) TargetSingleBuilder {
	app.evts = evts
	return app
}

// WithPath adds path to the builder
func (app *targetSingleBuilder) WithPath(path RelativePath) TargetSingleBuilder {
	app.path = path
	return app
}

// Now builds a new TargetSingle instance
func (app *targetSingleBuilder) Now() (TargetSingle, error) {
	if app.evts != nil && len(app.evts) <= 0 {
		app.evts = nil
	}

	if app.evts != nil {
		return createTargetSingleWithEvents(app.evts), nil
	}

	if app.path != nil {
		return createTargetSingleWithPath(app.path), nil
	}

	return nil, errors.New("the TargetSingle is invalid")
}
