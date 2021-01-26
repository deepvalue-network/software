package compilers

import (
	"github.com/steve-care-software/products/pangolin/domain/interpreters"
	"github.com/steve-care-software/products/pangolin/domain/linkers"
	"github.com/steve-care-software/products/pangolin/domain/middle"
	"github.com/steve-care-software/products/pangolin/domain/middle/variables/variable/value/computable"
)

// NewApplication creates a new application instance
func NewApplication(
	middleAdapter middle.Adapter,
	interpreterBuilder interpreters.Builder,
) Application {
	computableBuilder := computable.NewBuilder()
	programBuilder := linkers.NewProgramBuilder()
	return createApplication(computableBuilder, interpreterBuilder, middleAdapter, programBuilder)
}

// Application represents a compiler application
type Application interface {
	Execute(script linkers.Script) (middle.Program, error)
}
