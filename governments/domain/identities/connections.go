package identities

type conns struct {
	list []Connection
}

func createConnections(
	list []Connection,
) Connections {
	out := conns{
		list: list,
	}

	return &out
}

// All returns the list of connections
func (obj *conns) All() []Connection {
	return obj.list
}
