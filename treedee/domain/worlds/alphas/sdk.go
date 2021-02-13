package alphas

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents an alpha builder
type Builder interface {
	Create() Builder
	WithAlpha(alpha uint8) Builder
	WithVariable(variable string) Builder
	Now() (Alpha, error)
}

// Alpha represents an alpha
type Alpha interface {
	Alpha() uint8
	Variable() string
}
