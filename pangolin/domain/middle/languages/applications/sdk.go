package applications

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/heads"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/labels"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/tests"
)

// Application represents a language application
type Application interface {
	Head() heads.Head
	Labels() labels.Labels
	Main() instructions.Instructions
	HasTests() bool
	Tests() tests.Tests
}
