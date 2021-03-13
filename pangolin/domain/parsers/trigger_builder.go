package parsers

import "errors"

type triggerBuilder struct {
	variableName VariableName
	event        string
}

func createTriggerBuilder() TriggerBuilder {
	out := triggerBuilder{
		variableName: nil,
		event:        "",
	}

	return &out
}

// Create initializes the builder
func (app *triggerBuilder) Create() TriggerBuilder {
	return createTriggerBuilder()
}

// WithVariableName adds a variable name to the builder
func (app *triggerBuilder) WithVariableName(variableName VariableName) TriggerBuilder {
	app.variableName = variableName
	return app
}

// WithEvent adds an event to the builder
func (app *triggerBuilder) WithEvent(event string) TriggerBuilder {
	app.event = event
	return app
}

// Now builds a new Trigger instance
func (app *triggerBuilder) Now() (Trigger, error) {
	if app.variableName == nil {
		return nil, errors.New("the VariableName is mandatory in order to build a Trigger instance")
	}

	if app.event == "" {
		return nil, errors.New("the event is mandatory in order to build a Trigger instance")
	}

	return createTrigger(app.variableName, app.event), nil
}
