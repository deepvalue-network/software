package stackframe

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the stackframe builder
type Builder interface {
	Create() Builder
	IsPush() Builder
	IsPop() Builder
	Now() (Stackframe, error)
}

// Stackframe represents a stackframe instruction
type Stackframe interface {
	IsPush() bool
	IsPop() bool
}
