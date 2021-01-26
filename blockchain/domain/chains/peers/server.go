package peers

import "fmt"

type server struct {
	host string
	port uint
}

func createServer(
	host string,
	port uint,
) Server {
	out := server{
		host: host,
		port: port,
	}

	return &out
}

// Host returns the host
func (obj *server) Host() string {
	return obj.host
}

// Port returns the port
func (obj *server) Port() uint {
	return obj.port
}

// String returns the server as string
func (obj *server) String() string {
	return fmt.Sprintf("%s:%d", obj.host, obj.port)
}
