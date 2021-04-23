package applications

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/labels"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/tests"
	"github.com/deepvalue-network/software/pangolin/domain/middle/externals"
	"github.com/deepvalue-network/software/pangolin/domain/middle/heads"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

// NewAdapter creates a new adapter instance
func NewAdapter(
	testsAdapter tests.Adapter,
	labelsAdapter labels.Adapter,
	instructionsAdapter instructions.Adapter,
	headAdapter heads.Adapter,
) Adapter {
	builder := NewBuilder()
	return createAdapter(testsAdapter, labelsAdapter, instructionsAdapter, headAdapter, builder)
}

// NewBuilder initializes the builder
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents the application adapter
type Adapter interface {
	ToApplication(parsed parsers.Application) (Application, error)
}

// Builder represents an application builder
type Builder interface {
	Create() Builder
	WithHead(head heads.Head) Builder
	WithMain(main instructions.Instructions) Builder
	WithTests(tests tests.Tests) Builder
	WithLabels(labels labels.Labels) Builder
	WithExtends(extends []externals.External) Builder
	Now() (Application, error)
}

// Application represents an application
type Application interface {
	Head() heads.Head
	Main() instructions.Instructions
	Tests() tests.Tests
	Labels() labels.Labels
	HasExtends() bool
	Extends() []externals.External
}
