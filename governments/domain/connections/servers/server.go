package servers

import "fmt"

type server struct {
	protocol string
	host     Host
	port     uint
}

func createServer(
	protocol string,
	host Host,
	port uint,
) Server {
	out := server{
		protocol: protocol,
		host:     host,
		port:     port,
	}

	return &out
}

// Protocol returns the protocol
func (obj *server) Protocol() string {
	return obj.protocol
}

// Host returns the host
func (obj *server) Host() Host {
	return obj.host
}

// Port returns the port
func (obj *server) Port() uint {
	return obj.port
}

// String returns the string representation of the server
func (obj *server) String() string {
	return fmt.Sprintf("%s://%s:%d", obj.protocol, obj.host.String(), obj.port)
}
