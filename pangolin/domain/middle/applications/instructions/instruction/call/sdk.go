package call

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a call builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithCondition(condition string) Builder
	Now() (Call, error)
}

// Call represents a call
type Call interface {
	Name() string
	HasCondition() bool
	Condition() string
}
