package externals

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the external builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithPath(path string) Builder
	Now() (External, error)
}

// External represents an external
type External interface {
	Name() string
	Path() string
}
