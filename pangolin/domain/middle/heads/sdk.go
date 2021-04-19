package heads

// Head represents an head
type Head interface {
	Name() string
	Version() string
	HasImports() bool
	Imports() []External
}

// ExternalBuilder represents the external builder
type ExternalBuilder interface {
	Create() ExternalBuilder
	WithName(name string) ExternalBuilder
	WithPath(path string) ExternalBuilder
	Now() (External, error)
}

// External represents an external
type External interface {
	Name() string
	Path() string
}
