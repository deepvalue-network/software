package variable

import (
	var_value "github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction/variable/value"
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
	Now() (Variable, error)
}

// Variable represents a variable
type Variable interface {
	Name() string
	Value() var_value.Value
}
