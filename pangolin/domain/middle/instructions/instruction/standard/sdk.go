package standard

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	arythmeticBuilder := createArythmericBuilder()
	logicalBuilder := createLogicalBuilder()
	relationalBuilder := createRelationalBuilder()
	miscBuilder := createMiscBuilder()
	operationBuilder := createOperationBuilder()
	return createBuilder(
		arythmeticBuilder,
		logicalBuilder,
		relationalBuilder,
		miscBuilder,
		operationBuilder,
	)
}

// Builder represents a standard builder
type Builder interface {
	Create() Builder
	WithResult(result string) Builder
	WithFirst(first string) Builder
	WithSecond(second string) Builder
	WithOperation(operation Operation) Builder
	IsConcatenation() Builder
	IsFrameAssignment() Builder
	IsAdd() Builder
	IsSub() Builder
	IsMul() Builder
	IsLessThan() Builder
	IsEqual() Builder
	IsNotEqual() Builder
	IsAnd() Builder
	IsOr() Builder
	Now() (Standard, error)
}

// Standard represents a standard instruction
type Standard interface {
	Operation() Operation
	Result() string
	First() string
	Second() string
}

// OperationBuilder represents an operation builder
type OperationBuilder interface {
	Create() OperationBuilder
	WithArythmetic(arythmetic Arythmetic) OperationBuilder
	WithRelational(relational Relational) OperationBuilder
	WithLogical(logical Logical) OperationBuilder
	WithMisc(misc Misc) OperationBuilder
	Now() (Operation, error)
}

// Operation represents an operation instruction
type Operation interface {
	IsArythmetic() bool
	Arythmetic() Arythmetic
	IsRelational() bool
	Relational() Relational
	IsLogical() bool
	Logical() Logical
	IsMisc() bool
	Misc() Misc
}

// MiscBuilder represents a misc builder
type MiscBuilder interface {
	Create() MiscBuilder
	IsConcatenation() MiscBuilder
	IsFrameAssignment() MiscBuilder
	Now() (Misc, error)
}

// Misc represents a misc operation
type Misc interface {
	IsConcatenation() bool
	IsFrameAssignment() bool
}

// ArythmeticBuilder represents an arythmetic builder
type ArythmeticBuilder interface {
	Create() ArythmeticBuilder
	IsAdd() ArythmeticBuilder
	IsSub() ArythmeticBuilder
	IsMul() ArythmeticBuilder
	Now() (Arythmetic, error)
}

// Arythmetic represents an arythmetic operation instruction
type Arythmetic interface {
	IsAdd() bool
	IsSub() bool
	IsMul() bool
}

// RelationalBuilder represents the relational builder
type RelationalBuilder interface {
	Create() RelationalBuilder
	IsLessThan() RelationalBuilder
	IsEqual() RelationalBuilder
	IsNotEqual() RelationalBuilder
	Now() (Relational, error)
}

// Relational represents a relational operation instruction
type Relational interface {
	IsLessThan() bool
	IsEqual() bool
	IsNotEqual() bool
}

// LogicalBuilder represents the logical builder
type LogicalBuilder interface {
	Create() LogicalBuilder
	IsAnd() LogicalBuilder
	IsOr() LogicalBuilder
	Now() (Logical, error)
}

// Logical represents a logical operation instruction
type Logical interface {
	IsAnd() bool
	IsOr() bool
}
