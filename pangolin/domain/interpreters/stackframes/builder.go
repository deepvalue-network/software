package stackframes

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/variable/value/computable"
)

type builder struct {
	frameBuilder FrameBuilder
	variables    map[string]computable.Value
}

func createBuilder(frameBuilder FrameBuilder) Builder {
	out := builder{
		frameBuilder: frameBuilder,
		variables:    map[string]computable.Value{},
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.frameBuilder)
}

// WithVariables add variables to the builder
func (app *builder) WithVariables(variables map[string]computable.Value) Builder {
	app.variables = variables
	return app
}

// Now builds a new StackFrame instance
func (app *builder) Now() StackFrame {
	registry := createRegistry()
	return createStackFrame(app.frameBuilder, registry, app.variables)
}
