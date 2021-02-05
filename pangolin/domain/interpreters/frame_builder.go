package interpreters

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/variables/variable/value/computable"
)

type frameBuilder struct {
	computer  Computer
	builder   computable.Builder
	variables map[string]computable.Value
	constants map[string]computable.Value
}

func createFrameBuilder(
	computer Computer,
	builder computable.Builder,
) FrameBuilder {
	out := frameBuilder{
		computer:  computer,
		builder:   builder,
		variables: map[string]computable.Value{},
		constants: map[string]computable.Value{},
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

// WithConstants add constants to the builder
func (app *frameBuilder) WithConstants(constants map[string]computable.Value) FrameBuilder {
	app.constants = constants
	return app
}

// Now builds a new Frame instance
func (app *frameBuilder) Now() Frame {
	return createFrame(app.computer, app.builder, app.variables, app.constants)
}
