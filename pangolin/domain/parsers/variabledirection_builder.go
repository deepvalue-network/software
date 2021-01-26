package parsers

import "errors"

type variableDirectionBuilder struct {
	incoming   VariableIncoming
	isOutgoing bool
}

func createVariableDirectionBuilder() VariableDirectionBuilder {
	out := variableDirectionBuilder{
		incoming:   nil,
		isOutgoing: false,
	}

	return &out
}

// Create initializes the builder
func (app *variableDirectionBuilder) Create() VariableDirectionBuilder {
	return createVariableDirectionBuilder()
}

// IsIncoming flags the builder as incoming
func (app *variableDirectionBuilder) WithIncoming(incoming VariableIncoming) VariableDirectionBuilder {
	app.incoming = incoming
	return app
}

// IsOutgoing flags the builder as outgoing
func (app *variableDirectionBuilder) IsOutgoing() VariableDirectionBuilder {
	app.isOutgoing = true
	return app
}

// Now builds a new VariableDirection instance
func (app *variableDirectionBuilder) Now() (VariableDirection, error) {
	if app.incoming != nil && app.isOutgoing {
		return createVariableDirectionWithIncomingAndOutgoing(app.incoming), nil
	}

	if app.incoming != nil {
		return createVariableDirectionWithIncoming(app.incoming), nil
	}

	if app.isOutgoing {
		return createVariableDirectionWithOutgoing(), nil
	}

	return nil, errors.New("the VariableDirection instance is invalid")
}
