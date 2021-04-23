package heads

import "github.com/deepvalue-network/software/pangolin/domain/middle/externals"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
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
