package parsers

import "errors"

type eventsBuilder struct {
	evts []Event
}

func createEventsBuilder() EventsBuilder {
	out := eventsBuilder{
		evts: nil,
	}

	return &out
}

// Create initializes the builder
func (app *eventsBuilder) Create() EventsBuilder {
	return createEventsBuilder()
}

// WithEvents add []Event instances to the builder
func (app *eventsBuilder) WithEvents(evts []Event) EventsBuilder {
	app.evts = evts
	return app
}

// Now builds a new Events instance
func (app *eventsBuilder) Now() (Events, error) {
	if app.evts == nil {
		return nil, errors.New("the []Event are mandatory in order to build an Events instance")
	}

	return createEvents(app.evts), nil
}
