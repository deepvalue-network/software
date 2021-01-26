package value

import (
	var_value "github.com/steve-care-software/products/pangolin/domain/middle/variables/variable/value"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	operationBuilder := createOperationBuilder()
	return createBuilder(operationBuilder)
}

// Builder represents a value builder
type Builder interface {
	Create() Builder
	WithValue(val var_value.Value) Builder
	WithOperation(operation Operation) Builder
	IsPrint() Builder
	Now() (Value, error)
}

// Value represents a value instruction
type Value interface {
	Value() var_value.Value
	Operation() Operation
}

// OperationBuilder represents an operation builder
type OperationBuilder interface {
	Create() OperationBuilder
	IsPrint() OperationBuilder
	Now() (Operation, error)
}

// Operation represents a value operation
type Operation interface {
	IsPrint() bool
}
