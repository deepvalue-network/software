package module

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represets a module builder
type Builder interface {
	Create() Builder
	WithStackFrame(stackFrame string) Builder
	WithName(name string) Builder
	WithSymbol(symbol string) Builder
	Now() (Module, error)
}

// Module represents a module
type Module interface {
	StackFrame() string
	Name() string
	Symbol() string
}
