package trigger

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a trigger builder
type Builder interface {
	Create() Builder
	WithVariable(variable string) Builder
	WithEvent(event string) Builder
	Now() (Trigger, error)
}

// Trigger represents a trigger
type Trigger interface {
	Variable() string
	Event() string
}
