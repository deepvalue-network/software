package heads

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/externals"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

// NewAdapter creates a new adapter instance
func NewAdapter(
	externalAdapter externals.Adapter,
) Adapter {
	builder := NewBuilder()
	return createAdapter(externalAdapter, builder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents the adapter
type Adapter interface {
	ToHead(parsed parsers.HeadSection) (Head, error)
}

// Builder represents the head builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithVersion(version string) Builder
	WithImports(imports []externals.External) Builder
	Now() (Head, error)
}

// Head represents an head
type Head interface {
	Name() string
	Version() string
	HasImports() bool
	Imports() []externals.External
}
