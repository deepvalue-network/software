package machines

import (
	"errors"

	"github.com/deepvalue-network/software/pangolin/domain/interpreters/stackframes"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/variable/value/computable"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/labels"
)

type instructionBuilder struct {
	computableValueBuilder computable.Builder
	labels                 labels.Labels
	stackFrame             stackframes.StackFrame
}

func createInstructionBuilder(
	computableValueBuilder computable.Builder,
) InstructionBuilder {
	out := instructionBuilder{
		computableValueBuilder: computableValueBuilder,
		labels:                 nil,
		stackFrame:             nil,
	}

	return &out
}

// Create initializes the builder
func (app *instructionBuilder) Create() InstructionBuilder {
	return createInstructionBuilder(
		app.computableValueBuilder,
	)
}

// WithLabels add labels to the builder
func (app *instructionBuilder) WithLabels(labels labels.Labels) InstructionBuilder {
	app.labels = labels
	return app
}

// WithStackFrame adds a stackframe func to the builder
func (app *instructionBuilder) WithStackFrame(stackFrame stackframes.StackFrame) InstructionBuilder {
	app.stackFrame = stackFrame
	return app
}

// Now builds a new Instruction instance
func (app *instructionBuilder) Now() (Instruction, error) {
	if app.labels == nil {
		return nil, errors.New("the Labels are mandatory in order to build a Instruction instance")
	}

	if app.stackFrame == nil {
		return nil, errors.New("the stackFrame is mandatory in order to build a Instruction instance")
	}

	return createInstruction(app.computableValueBuilder, app.labels, app.stackFrame), nil
}
