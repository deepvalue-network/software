package test

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/tests/test/instructions"
)

// Test represents a test
type Test interface {
	Name() string
	Instructions() instructions.Instructions
}
