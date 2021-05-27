package labels

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/labels/label"
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

// Adapter represents the labels adapter
type Adapter interface {
	ToLabels(section parsers.LabelSection) (Labels, error)
}

// Builder represents the labels builder
type Builder interface {
	Create() Builder
	WithList(lst []label.Label) Builder
	WithMap(mp map[string]label.Label) Builder
	Now() (Labels, error)
}

// Labels represents labels
type Labels interface {
	All() []label.Label
	Fetch(name string) (label.Label, error)
}
