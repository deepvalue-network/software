package servers

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
