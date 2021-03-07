package exit

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents an exit builder
type Builder interface {
	Create() Builder
	WithCondition(condition string) Builder
	Now() (Exit, error)
}

// Exit represents an exit
type Exit interface {
	HasCondition() bool
	Condition() string
}
