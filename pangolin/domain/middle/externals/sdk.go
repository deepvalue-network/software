package externals

import "github.com/deepvalue-network/software/pangolin/domain/parsers"

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	builder := NewBuilder()
	return createAdapter(builder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents an adapter
type Adapter interface {
	ToExternals(parsed []parsers.ImportSingle) ([]External, error)
	ToExternal(parsed parsers.ImportSingle) (External, error)
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
