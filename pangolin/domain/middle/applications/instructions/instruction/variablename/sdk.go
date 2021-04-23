package variablename

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	miscBuilder := createMiscBuilder()
	operationBuilder := createOperationBuilder(miscBuilder)
	return createBuilder(
		operationBuilder,
	)
}

// Builder represents a variableName builder
type Builder interface {
	Create() Builder
	WithVariableName(vr string) Builder
	WithOperation(operation Operation) Builder
	IsPush() Builder
	Now() (VariableName, error)
}

// VariableName represents a variableName instruction
type VariableName interface {
	Operation() Operation
	Variable() string
}

// OperationBuilder represents an operation builder
type OperationBuilder interface {
	Create() OperationBuilder
	WithMisc(misc Misc) OperationBuilder
	IsPush() OperationBuilder
	Now() (Operation, error)
}

// Operation represents an operation instruction
type Operation interface {
	IsMisc() bool
	Misc() Misc
}

// MiscBuilder represents a misc builder
type MiscBuilder interface {
	Create() MiscBuilder
	IsPush() MiscBuilder
	Now() (Misc, error)
}

// Misc represents a misc operation
type Misc interface {
	IsPush() bool
}
