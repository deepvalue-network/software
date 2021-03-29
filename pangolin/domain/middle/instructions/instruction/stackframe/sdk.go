package stackframe

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewSkipBuilder creates a new skip instance
func NewSkipBuilder() SkipBuilder {
	return createSkipBuilder()
}

// Builder represents the stackframe builder
type Builder interface {
	Create() Builder
	IsPush() Builder
	IsPop() Builder
	WithSkip(skip Skip) Builder
	WithIndex(indexVariable string) Builder
	Now() (Stackframe, error)
}

// Stackframe represents a stackframe instruction
type Stackframe interface {
	IsPush() bool
	IsPop() bool
	IsIndex() bool
	Index() string
	IsSkip() bool
	Skip() Skip
}

// SkipBuilder represents a skip builder
type SkipBuilder interface {
	Create() SkipBuilder
	WithInt(intVal int64) SkipBuilder
	WithVariable(variable string) SkipBuilder
	Now() (Skip, error)
}

// Skip represents a skip
type Skip interface {
	IsInt() bool
	Int() int64
	IsVariable() bool
	Variable() string
}
