package identities

type connectionsBuilder struct {
	list []Connection
}

func createConnectionsBuilder() ConnectionsBuilder {
	out := connectionsBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *connectionsBuilder) Create() ConnectionsBuilder {
	return createConnectionsBuilder()
}

// WithConnections add connections to the builder
func (app *connectionsBuilder) WithConnections(connections []Connection) ConnectionsBuilder {
	app.list = connections
	return app
}

// Now builds a new Connections instance
func (app *connectionsBuilder) Now() (Connections, error) {
	if app.list == nil {
		app.list = []Connection{}
	}

	return createConnections(app.list), nil
}
