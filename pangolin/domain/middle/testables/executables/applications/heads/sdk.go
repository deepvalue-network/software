package heads

import (
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	builder := NewBuilder()
	return createAdapter(builder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewLoadSingleBuilder creates a new load single builder
func NewLoadSingleBuilder() LoadSingleBuilder {
	return createLoadSingleBuilder()
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
	WithImports(imports []parsers.ImportSingle) Builder
	WithLoads(loads []LoadSingle) Builder
	Now() (Head, error)
}

// Head represents an head
type Head interface {
	Name() string
	Version() string
	HasImports() bool
	Imports() []parsers.ImportSingle
	HasLoads() bool
	Loads() []LoadSingle
}

// LoadSingleBuilder represents a load single builder
type LoadSingleBuilder interface {
	Create() LoadSingleBuilder
	WithInternal(internal string) LoadSingleBuilder
	WithExternal(external string) LoadSingleBuilder
	Now() (LoadSingle, error)
}

// LoadSingle represents a load single
type LoadSingle interface {
	Internal() string
	External() string
}
