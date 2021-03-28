package variables

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/variables/variable"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the variables builder
type Builder interface {
	Create() Builder
	WithVariables(vrs []variable.Variable) Builder
	Now() (Variables, error)
}

// Variables represents variables
type Variables interface {
	All() []variable.Variable
	Merge(vr Variables) Variables
}
