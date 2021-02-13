package glfw

import (
	"github.com/deepvalue-network/software/treedee/application/windows"
)

// NewBuilder creates a new glfw application builder
func NewBuilder() windows.Builder {
	return createBuilder()
}
