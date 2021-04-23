package transform

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	miscBuilder := createMiscBuilder()
	operationBuilder := createOperationBuilder(miscBuilder)
	return createBuilder(
		operationBuilder,
	)
}

// Builder represents a transform builder
type Builder interface {
	Create() Builder
	WithResult(result string) Builder
	WithInput(input string) Builder
	WithOperation(operation Operation) Builder
	IsPop() Builder
	Now() (Transform, error)
}

// Transform represents a transform instruction
type Transform interface {
	Operation() Operation
	Result() string
	Input() string
}

// OperationBuilder represents an operation builder
type OperationBuilder interface {
	Create() OperationBuilder
	WithMisc(misc Misc) OperationBuilder
	IsPop() OperationBuilder
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
	IsPop() MiscBuilder
	Now() (Misc, error)
}

// Misc represents a misc operation
type Misc interface {
	IsPop() bool
}
