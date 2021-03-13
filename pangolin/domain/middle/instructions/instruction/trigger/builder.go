package trigger

import "errors"

type builder struct {
	variable string
	event    string
}

func createBuilder() Builder {
	out := builder{
		variable: "",
		event:    "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithVariable adds a variable to the builder
func (app *builder) WithVariable(variable string) Builder {
	app.variable = variable
	return app
}

// WithEvent adds an event to the builder
func (app *builder) WithEvent(event string) Builder {
	app.event = event
	return app
}

// Now builds a new Trigger instance
func (app *builder) Now() (Trigger, error) {
	if app.variable == "" {
		return nil, errors.New("the variable is mandatory in order to build a Trigger instance")
	}

	if app.event == "" {
		return nil, errors.New("the event is mandatory in order to build a Trigger instance")
	}

	return createTrigger(app.variable, app.event), nil
}
