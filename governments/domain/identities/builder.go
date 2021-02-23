package identities

import "errors"

type builder struct {
	connectionsBuilder ConnectionsBuilder
	name               string
	seed               string
	holders            ShareHolders
	conns              Connections
}

func createBuilder(
	connectionsBuilder ConnectionsBuilder,
) Builder {
	out := builder{
		connectionsBuilder: connectionsBuilder,
		name:               "",
		seed:               "",
		holders:            nil,
		conns:              nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.connectionsBuilder)
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

// WithShareHolders adds shareHolders to the builder
func (app *builder) WithShareHolders(shareHolders ShareHolders) Builder {
	app.holders = shareHolders
	return app
}

// WithConnections adds connections to the builder
func (app *builder) WithConnections(connections Connections) Builder {
	app.conns = connections
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

	if app.conns == nil {
		conns, err := app.connectionsBuilder.Create().Now()
		if err != nil {
			return nil, err
		}

		app.conns = conns
	}

	return createIdentity(app.name, app.seed, app.holders, app.conns), nil
}
