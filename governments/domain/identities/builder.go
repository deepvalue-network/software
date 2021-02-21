package identities

import "errors"

type builder struct {
	name    string
	seed    string
	holders ShareHolders
}

func createBuilder() Builder {
	out := builder{
		name:    "",
		seed:    "",
		holders: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// WithSeed adds a seed to the builder
func (app *builder) WithSeed(seed string) Builder {
	app.seed = seed
	return app
}

// WithShareHolders adds a shareHolders to the builder
func (app *builder) WithShareHolders(shareHolders ShareHolders) Builder {
	app.holders = shareHolders
	return app
}

// Now builds a new Identity instance
func (app *builder) Now() (Identity, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build an Identity instance")
	}

	if app.seed == "" {
		return nil, errors.New("the seed is mandatory in order to build an Identity instance")
	}

	if app.holders == nil {
		return nil, errors.New("the shareHolders are mandatory in order to build an Identity instance")
	}

	return createIdentity(app.name, app.seed, app.holders), nil
}
