package heads

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the head builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithVersion(version string) Builder
	WithImports(imports []External) Builder
	Now() (Head, error)
}

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
