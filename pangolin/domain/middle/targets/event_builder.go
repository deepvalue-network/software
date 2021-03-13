package targets

import "errors"

type eventBuilder struct {
	name  string
	label string
}

func createEventBuilder() EventBuilder {
	out := eventBuilder{
		name:  "",
		label: "",
	}

	return &out
}

// Create initializes the builder
func (app *eventBuilder) Create() EventBuilder {
	return createEventBuilder()
}

// WithName adds a name to the builder
func (app *eventBuilder) WithName(name string) EventBuilder {
	app.name = name
	return app
}

// WithLabel adds a label to the builder
func (app *eventBuilder) WithLabel(label string) EventBuilder {
	app.label = label
	return app
}

// Now builds a new Event instance
func (app *eventBuilder) Now() (Event, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build an Event instance")
	}

	if app.label == "" {
		return nil, errors.New("the label is mandatory in order to build an Event instance")
	}

	return createEvent(app.name, app.label), nil
}
