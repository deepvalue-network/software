package condition

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	operationBuilder := createOperationBuilder()
	return createBuilder(operationBuilder)
}

// NewPropositionBuilder creates a new proposition builder
func NewPropositionBuilder() PropositionBuilder {
	return createPropositionBuilder()
}

// Builder represents a condition builder
type Builder interface {
	Create() Builder
	WithProposition(proposition Proposition) Builder
	IsJump() Builder
	Now() (Condition, error)
}

// Condition represents a condition
type Condition interface {
	Proposition() Proposition
	Operation() Operation
}

// OperationBuilder represents an operation builder
type OperationBuilder interface {
	Create() OperationBuilder
	IsJump() OperationBuilder
	Now() (Operation, error)
}

// Operation represents a condition operation
type Operation interface {
	IsJump() bool
}

// PropositionBuilder represents a proposition builder
type PropositionBuilder interface {
	Create() PropositionBuilder
	WithName(name string) PropositionBuilder
	WithCondition(condition string) PropositionBuilder
	Now() (Proposition, error)
}

// Proposition represents a proposition
type Proposition interface {
	Name() string
	HasCondition() bool
	Condition() string
}
