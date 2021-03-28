package variable

import (
	var_value "github.com/deepvalue-network/software/pangolin/domain/middle/variables/variable/value"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the variable builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithValue(val var_value.Value) Builder
	IsIncoming() Builder
	IsOutgoing() Builder
	Now() (Variable, error)
}

// Variable represents a variable
type Variable interface {
	IsIncoming() bool
	IsOutgoing() bool
	Name() string
	Value() var_value.Value
}
