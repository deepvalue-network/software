package parsers

import "errors"

type eventBuilder struct {
	token               string
	set                 SetFn
	enter               EventFn
	exit                EventFn
	retrieveReplacement RetrieveReplacementsFn
}

func createEventBuilder() EventBuilder {
	out := eventBuilder{
		token:               "",
		set:                 nil,
		enter:               nil,
		exit:                nil,
		retrieveReplacement: nil,
	}

	return &out
}

// Create initializes the builder
func (app *eventBuilder) Create() EventBuilder {
	return createEventBuilder()
}

// WithToken add a token to the builder
func (app *eventBuilder) WithToken(token string) EventBuilder {
	app.token = token
	return app
}

// WithSet add a setFn to the builder
func (app *eventBuilder) WithSet(set SetFn) EventBuilder {
	app.set = set
	return app
}

// WithOnEnter adds an onEnter event func to the builder
func (app *eventBuilder) WithOnEnter(onEnter EventFn) EventBuilder {
	app.enter = onEnter
	return app
}

// WithOnExit adds an onExit event func to the builder
func (app *eventBuilder) WithOnExit(onExit EventFn) EventBuilder {
	app.exit = onExit
	return app
}

// WithRetrieveReplacement adds a retrieveReplacementFn to the builder
func (app *eventBuilder) WithRetrieveReplacement(retrieveReplacement RetrieveReplacementsFn) EventBuilder {
	app.retrieveReplacement = retrieveReplacement
	return app
}

// Now builds a new Event instance
func (app *eventBuilder) Now() (Event, error) {
	if app.token == "" {
		return nil, errors.New("the token string is mandatory in order to build an Event instance")
	}

	if app.enter != nil {
		if app.retrieveReplacement != nil {
			if app.set != nil {
				return createEventWithOnEnterAndRetrieveReplacementAndSet(app.token, app.enter, app.retrieveReplacement, app.set), nil
			}

			return createEventWithOnEnterAndRetrieveReplacement(app.token, app.enter, app.retrieveReplacement), nil
		}

		if app.set != nil {
			return createEventWithOnEnterAndSet(app.token, app.enter, app.set), nil
		}

		return createEventWithOnEnter(app.token, app.enter), nil
	}

	if app.exit != nil {
		if app.retrieveReplacement != nil {
			if app.set != nil {
				return createEventWithOnExitAndRetrieveReplacementAndSet(app.token, app.exit, app.retrieveReplacement, app.set), nil
			}

			return createEventWithOnExitAndRetrieveReplacement(app.token, app.exit, app.retrieveReplacement), nil
		}

		if app.set != nil {
			return createEventWithOnExitAndSet(app.token, app.set, app.exit), nil
		}

		return createEventWithOnExit(app.token, app.exit), nil
	}

	return nil, errors.New("the enter/exit func is mandatory in order to build an Event instance")
}
