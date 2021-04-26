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
	valueBuilder := NewValueBuilder()
	return createAdapter(externalAdapter, builder, valueBuilder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewValueBuilder creates a new value builder instance
func NewValueBuilder() ValueBuilder {
	return createValueBuilder()
}

// Adapter represents an adapter
type Adapter interface {
	ToHead(parsed parsers.HeadCommand) (Head, error)
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
	WithImports(imports []externals.External) ValueBuilder
	Now() (Value, error)
}

// Value represents an head value
type Value interface {
	IsName() bool
	Name() string
	IsVersion() bool
	Version() string
	IsImports() bool
	Imports() []externals.External
}
