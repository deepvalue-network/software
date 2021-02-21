package swaps

import "errors"

type builder struct {
	in  Incoming
	out Complete
}

func createBuilder() Builder {
	out := builder{
		in:  nil,
		out: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithIncoming adds an incoming to the builder
func (app *builder) WithIncoming(incoming Incoming) Builder {
	app.in = incoming
	return app
}

// WithOutgoing adds an outgoing to the builder
func (app *builder) WithOutgoing(outgoing Complete) Builder {
	app.out = outgoing
	return app
}

// Now builds a new Swap instance
func (app *builder) Now() (Swap, error) {
	if app.in != nil {
		return createSwapWithIncoming(app.in), nil
	}

	if app.out != nil {
		return createSwapWithOutgoing(app.out), nil
	}

	return nil, errors.New("the Swap is invalid")

}
