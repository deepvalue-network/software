package labels

import (
	"github.com/steve-care-software/products/pangolin/domain/parsers"
	"github.com/steve-care-software/products/pangolin/domain/middle/labels/label"
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
}
