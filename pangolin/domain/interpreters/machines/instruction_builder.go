package machines

import (
	"errors"

	"github.com/deepvalue-network/software/pangolin/domain/interpreters/stackframes"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/variable/value/computable"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/labels"
)

type instructionBuilder struct {
	computableValueBuilder computable.Builder
	labelFn                CallLabelByNameFn
	labels                 labels.Labels
	stackFrame             stackframes.StackFrame
}

func createInstructionBuilder(
	computableValueBuilder computable.Builder,
) InstructionBuilder {
	out := instructionBuilder{
		computableValueBuilder: computableValueBuilder,
		labelFn:                nil,
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

// WithCallLabelFn adds a callLabelByNameFn to the builder
func (app *instructionBuilder) WithCallLabelFn(labelFn CallLabelByNameFn) InstructionBuilder {
	app.labelFn = labelFn
	return app
}

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
	if app.stackFrame == nil {
		return nil, errors.New("the stackFrame is mandatory in order to build a Instruction instance")
	}

	if app.labels != nil {
		fn, err := fromLabelsToCallLabelByNameFunc(app.stackFrame, app.labels)
		if err != nil {
			return nil, err
		}

		app.labelFn = fn
	}

	if app.labelFn == nil {
		return nil, errors.New("the CallLabelByNameFn are mandatory in order to build a Instruction instance")
	}

	return createInstruction(app.computableValueBuilder, app.labelFn, app.stackFrame), nil
}
