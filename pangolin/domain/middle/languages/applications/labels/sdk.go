package labels

import "github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/labels/label"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a labels builder
type Builder interface {
	Create() Builder
	WithList(list []label.Label) Builder
	Now() (Labels, error)
}

// Labels represents labels
type Labels interface {
	All() []label.Label
}
