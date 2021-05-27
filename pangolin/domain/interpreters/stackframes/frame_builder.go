package stackframes

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction/variable/value/computable"
)

type frameBuilder struct {
	computer  Computer
	builder   computable.Builder
	variables map[string]computable.Value
}

func createFrameBuilder(
	computer Computer,
	builder computable.Builder,
) FrameBuilder {
	out := frameBuilder{
		computer:  computer,
		builder:   builder,
		variables: map[string]computable.Value{},
	}

	return &out
}

// Create initializes the builder
func (app *frameBuilder) Create() FrameBuilder {
	return createFrameBuilder(app.computer, app.builder)
}

// WithVariables add variables to the builder
func (app *frameBuilder) WithVariables(variables map[string]computable.Value) FrameBuilder {
	app.variables = variables
	return app
}

// Now builds a new Frame instance
func (app *frameBuilder) Now() Frame {

	variables := map[string]computable.Value{}
	for keyname, value := range app.variables {
		variables[keyname] = value
	}

	return createFrame(app.computer, app.builder, variables)
}
