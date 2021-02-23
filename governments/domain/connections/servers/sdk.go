package servers

const onionTLD = "onion"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a server builder
type Builder interface {
	Create() Builder
	WithURL(str string) Builder
	Now() (Server, error)
}

// Server represents a server
type Server interface {
	Protocol() string
	Host() Host
	Port() uint
	String() string
}

// Host represents a server host
type Host interface {
	String() string
	IsClear() bool
	Clear() string
	IsOnion() bool
	Onion() string
}
