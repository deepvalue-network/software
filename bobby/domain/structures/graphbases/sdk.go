package graphbases

import (
	"github.com/deepvalue-network/software/blockchain/domain/chains"
	"github.com/deepvalue-network/software/bobby/domain/resources"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a graphbase builder
type Builder interface {
	Create() Builder
	WithResource(res resources.Accessible) Builder
	WithMetaData(metaData string) Builder
	WithParent(parent resources.Accessible) Builder
	OnChain(chain chains.Chain) Builder
	Now() (Graphbase, error)
}

// Graphbase represents a graphbase
type Graphbase interface {
	Resource() resources.Accessible
	MetaData() string
	Chain() chains.Chain
	HasParent() bool
	Parent() resources.Accessible
}
