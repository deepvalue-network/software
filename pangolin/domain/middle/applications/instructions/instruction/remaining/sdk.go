package remaining

// NewBuilder creates a new builder
func NewBuilder() Builder {
	arythmeticBuilder := createArythmeticBuilder()
	miscBuilder := createMiscBuilder()
	operationBuilder := createOperationBuilder()
	return createBuilder(arythmeticBuilder, miscBuilder, operationBuilder)
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
	IsMatch() Builder
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
	WithMisc(misc Misc) OperationBuilder
	Now() (Operation, error)
}

// Operation represents an operation instruction
type Operation interface {
	IsMisc() bool
	Misc() Misc
	IsArythmetic() bool
	Arythmetic() Arythmetic
}

// MiscBuilder represents a misc builder
type MiscBuilder interface {
	Create() MiscBuilder
	IsMatch() MiscBuilder
	Now() (Misc, error)
}

// Misc represents a misc operation
type Misc interface {
	IsMatch() bool
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
