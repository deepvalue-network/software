package servers

// Builder represents a server builder
type Builder interface {
	Create() Builder
	WithHost(host string) Builder
	WithPort(port uint) Builder
	Now() (Server, error)
}

// Server represents a server
type Server interface {
	Content() Content
	Port() uint
	String() string
}

// Content represents a server content
type Content interface {
	IsClear() bool
	Clear() string
	IsOnion() bool
	Onion() string
}
