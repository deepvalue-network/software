package applications

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/heads"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/labels"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/tests"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents an application builder
type Builder interface {
	Create() Builder
	WithHead(head heads.Head) Builder
	WithLabels(labels labels.Labels) Builder
	WithMain(main instructions.Instructions) Builder
	WithTests(tests tests.Tests) Builder
	Now() (Application, error)
}

// Application represents a language application
type Application interface {
	Head() heads.Head
	Labels() labels.Labels
	Main() instructions.Instructions
	HasTests() bool
	Tests() tests.Tests
}
