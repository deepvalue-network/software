package value

import (
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
	"github.com/deepvalue-network/software/pangolin/domain/middle/variables/variable/value/computable"
)

const (
	defaultInt    = 0
	defaultUint   = 0
	defaultFloat  = 0.0
	defaultString = ""
	defaultBool   = false
)

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	computableBuilder := computable.NewBuilder()
	builder := NewBuilder()
	return createAdapter(computableBuilder, builder)
}

// NewFactory creates a new factory instance
func NewFactory() Factory {
	computableBuilder := computable.NewBuilder()
	builder := NewBuilder()
	return createFactory(computableBuilder, builder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents the value adapter
type Adapter interface {
	ToValue(parsed parsers.Value) (Value, error)
	ToValueWithType(typ parsers.Type, parsed parsers.Value) (Value, error)
}

// Factory represents a value factory
type Factory interface {
	Create(typ parsers.Type) (Value, error)
}

// Builder represents a value builder
type Builder interface {
	Create() Builder
	WithComputable(computable computable.Value) Builder
	WithGlobalVariable(globalVariable string) Builder
	WithLocalVariable(localVariable string) Builder
	Now() (Value, error)
}

// Value represents a value
type Value interface {
	IsComputable() bool
	Computable() computable.Value
	IsGlobalVariable() bool
	GlobalVariable() string
	IsLocalVariable() bool
	LocalVariable() string
}
