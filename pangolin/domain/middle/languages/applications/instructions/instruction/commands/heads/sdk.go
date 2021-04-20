package heads

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/heads"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents an head builder
type Builder interface {
	Create() Builder
	WithVariable(variable string) Builder
	WithValues(values []Value) Builder
	Now() (Head, error)
}

// Head represents a head command
type Head interface {
	Variable() string
	Values() []Value
}

// ValueBuilder represents an headValue builder
type ValueBuilder interface {
	Create() ValueBuilder
	WithName(name string) ValueBuilder
	WithVersion(version string) ValueBuilder
	WithImports(imports []heads.External) ValueBuilder
	Now() (Value, error)
}

// Value represents an head value
type Value interface {
	IsName() bool
	Name() string
	IsVersion() bool
	Version() string
	IsImports() bool
	Imports() []heads.External
}
