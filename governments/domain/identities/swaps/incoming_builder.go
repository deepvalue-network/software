package swaps

import (
	"errors"

	"github.com/deepvalue-network/software/governments/domain/identities/transfers"
)

type incomingBuilder struct {
	complete Complete
	incoming transfers.Transfer
}

func createIncomingBuilder() IncomingBuilder {
	out := incomingBuilder{
		complete: nil,
		incoming: nil,
	}

	return &out
}

// Create initializes the builder
func (app *incomingBuilder) Create() IncomingBuilder {
	return createIncomingBuilder()
}

// WithComplete adds a complete instance to the builder
func (app *incomingBuilder) WithComplete(complete Complete) IncomingBuilder {
	app.complete = complete
	return app
}

// WithIncoming adds a incoming transfer instance to the builder
func (app *incomingBuilder) WithIncoming(incoming transfers.Transfer) IncomingBuilder {
	app.incoming = incoming
	return app
}

// Now builds a new Incomign instance
func (app *incomingBuilder) Now() (Incoming, error) {
	if app.complete == nil {
		return nil, errors.New("the complete instance is mandatory in order to build an Incoming instance")
	}

	if app.incoming == nil {
		return nil, errors.New("the incoming transfer is mandatory in order to build an Incoming instance")
	}

	return createIncoming(app.complete, app.incoming), nil
}
