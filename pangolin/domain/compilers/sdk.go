package compilers

import (
	"github.com/deepvalue-network/software/pangolin/domain/interpreters"
	"github.com/deepvalue-network/software/pangolin/domain/linkers"
	"github.com/deepvalue-network/software/pangolin/domain/middle"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/variable/value/computable"
)

// NewApplication creates a new application instance
func NewApplication(
	middleAdapter middle.Adapter,
	interpreterBuilder interpreters.Builder,
) Application {
	computableBuilder := computable.NewBuilder()
	programBuilder := linkers.NewProgramBuilder()
	languageBuilder := linkers.NewLanguageBuilder()
	return createApplication(
		computableBuilder,
		interpreterBuilder,
		middleAdapter,
		programBuilder,
		languageBuilder,
	)
}

// Application represents a compiler application
type Application interface {
	Execute(script linkers.Script) (middle.Program, error)
}
