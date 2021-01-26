package lexers

import "errors"

type eventBuilder struct {
	token string
	fn    EventFn
}

func createEventBuilder() EventBuilder {
	out := eventBuilder{
		token: "",
		fn:    nil,
	}

	return &out
}

// Create initializes the builder
func (obj *eventBuilder) Create() EventBuilder {
	return createEventBuilder()
}

// WithToken adds a token to the builder
func (obj *eventBuilder) WithToken(token string) EventBuilder {
	obj.token = token
	return obj
}

// WithFn adds an event Func to the builder
func (obj *eventBuilder) WithFn(fn EventFn) EventBuilder {
	obj.fn = fn
	return obj
}

// Now builds a new Event instance
func (obj *eventBuilder) Now() (Event, error) {
	if obj.token == "" {
		return nil, errors.New("the token is mandatory in order to build an Event instance")
	}

	if obj.fn == nil {
		return nil, errors.New("the EventFn is mandatory in order to build an Event instance")
	}

	return createEvent(obj.token, obj.fn), nil
}
