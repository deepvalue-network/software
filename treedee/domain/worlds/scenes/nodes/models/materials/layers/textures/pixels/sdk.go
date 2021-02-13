package pixels

import (
	"github.com/deepvalue-network/software/treedee/domain/worlds/colors"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a pixel builder
type Builder interface {
	Create() Builder
	WithColor(color colors.Color) Builder
	WithAlpha(alpha uint8) Builder
	Now() (Pixel, error)
}

// Pixel represents a pixel
type Pixel interface {
	Color() colors.Color
	Alpha() uint8
}
