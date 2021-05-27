package labels

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/labels/label"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	labelAdapter := label.NewAdapter()
	builder := NewBuilder()
	return createAdapter(labelAdapter, builder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents an adapter
type Adapter interface {
	ToLabels(parsed parsers.LanguageLabelSection) (Labels, error)
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
	Fetch(name string) (label.Label, error)
}
