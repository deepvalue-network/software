package interpreters

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/variable/value/computable"
)

type stackFrameBuilder struct {
	frameBuilder FrameBuilder
	variables    map[string]computable.Value
}

func createStackFrameBuilder(frameBuilder FrameBuilder) StackFrameBuilder {
	out := stackFrameBuilder{
		frameBuilder: frameBuilder,
		variables:    map[string]computable.Value{},
	}

	return &out
}

// Create initializes the builder
func (app *stackFrameBuilder) Create() StackFrameBuilder {
	return createStackFrameBuilder(app.frameBuilder)
}

// WithVariables add variables to the builder
func (app *stackFrameBuilder) WithVariables(variables map[string]computable.Value) StackFrameBuilder {
	app.variables = variables
	return app
}

// Now builds a new StackFrame instance
func (app *stackFrameBuilder) Now() StackFrame {
	return createStackFrame(app.frameBuilder, app.variables)
}
