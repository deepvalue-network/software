package machines

import (
	"errors"

	"github.com/deepvalue-network/software/pangolin/domain/interpreters/stackframes"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/variable/value/computable"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/labels"
)

type testInstructionBuilder struct {
	instructionBuilder InstructionBuilder
	computableBuilder  computable.Builder
	stackFrame         stackframes.StackFrame
	labelFn            CallLabelByNameFn
	labels             labels.Labels
	baseDir            string
}

func createTestInstructionBuilder(
	instructionBuilder InstructionBuilder,
	computableBuilder computable.Builder,
) TestInstructionBuilder {
	out := testInstructionBuilder{
		instructionBuilder: instructionBuilder,
		computableBuilder:  computableBuilder,
		stackFrame:         nil,
		labelFn:            nil,
		labels:             nil,
	}

	return &out
}

// Create initializes the builder
func (app *testInstructionBuilder) Create() TestInstructionBuilder {
	return createTestInstructionBuilder(app.instructionBuilder, app.computableBuilder)
}

// WithCallLabelFn adds a call label func to the builder
func (app *testInstructionBuilder) WithCallLabelFn(labelFn CallLabelByNameFn) TestInstructionBuilder {
	app.labelFn = labelFn
	return app
}

// WithLabels add labels to the builder
func (app *testInstructionBuilder) WithLabels(labels labels.Labels) TestInstructionBuilder {
	app.labels = labels
	return app
}

// WithStackFrame adds a stackFrame func to the builder
func (app *testInstructionBuilder) WithStackFrame(stackFrame stackframes.StackFrame) TestInstructionBuilder {
	app.stackFrame = stackFrame
	return app
}

// WithBaseDir adds a baseDir to the builder
func (app *testInstructionBuilder) WithBaseDir(baseDir string) TestInstructionBuilder {
	app.baseDir = baseDir
	return app
}

// Now builds a new TestInstruction instance
func (app *testInstructionBuilder) Now() (TestInstruction, error) {
	if app.stackFrame == nil {
		return nil, errors.New("the stackFrame is mandatory in order to build a TestInstruction instance")
	}

	if app.labels != nil {
		fn, err := fromLabelsToCallLabelByNameFunc(app.stackFrame, app.labels)
		if err != nil {
			return nil, err
		}

		app.labelFn = fn
	}

	if app.labelFn == nil {
		return nil, errors.New("the CallLabelByNameFn are mandatory in order to build a TestInstruction instance")
	}

	insApp, err := app.instructionBuilder.Create().WithCallLabelFn(app.labelFn).WithStackFrame(app.stackFrame).Now()
	if err != nil {
		return nil, err
	}

	return createTestInstruction(app.computableBuilder, app.stackFrame, insApp, app.baseDir), nil
}
