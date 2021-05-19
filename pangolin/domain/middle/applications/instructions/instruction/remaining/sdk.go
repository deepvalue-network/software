package remaining

// NewBuilder creates a new builder
func NewBuilder() Builder {
	arythmeticBuilder := createArythmeticBuilder()
	operationBuilder := createOperationBuilder()
	return createBuilder(arythmeticBuilder, operationBuilder)
}

// Builder represents the remaining builder
type Builder interface {
	Create() Builder
	WithResult(result string) Builder
	WithRemaining(remaining string) Builder
	WithFirst(first string) Builder
	WithSecond(second string) Builder
	WithOperation(operation Operation) Builder
	IsDiv() Builder
	Now() (Remaining, error)
}

// Remaining represents a remaining instruction
type Remaining interface {
	Operation() Operation
	Result() string
	Remaining() string
	First() string
	Second() string
}

// OperationBuilder represents an operation builder
type OperationBuilder interface {
	Create() OperationBuilder
	WithArythmetic(ary Arythmetic) OperationBuilder
	Now() (Operation, error)
}

// Operation represents an operation instruction
type Operation interface {
	IsArythmetic() bool
	Arythmetic() Arythmetic
}

// ArythmeticBuilder represents the arythmetic builder
type ArythmeticBuilder interface {
	Create() ArythmeticBuilder
	IsDiv() ArythmeticBuilder
	Now() (Arythmetic, error)
}

// Arythmetic represents an arythmetic operation instruction
type Arythmetic interface {
	IsDiv() bool
}
