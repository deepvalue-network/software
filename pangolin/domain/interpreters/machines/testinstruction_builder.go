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
		labels:             nil,
	}

	return &out
}

// Create initializes the builder
func (app *testInstructionBuilder) Create() TestInstructionBuilder {
	return createTestInstructionBuilder(app.instructionBuilder, app.computableBuilder)
}

// WithLabels add Labels to the builder
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
	if app.labels == nil {
		return nil, errors.New("the Labels are mandatory in order to build a TestInstruction instance")
	}

	if app.stackFrame == nil {
		return nil, errors.New("the stackFrame is mandatory in order to build a TestInstruction instance")
	}

	insApp, err := app.instructionBuilder.Create().WithLabels(app.labels).WithStackFrame(app.stackFrame).Now()
	if err != nil {
		return nil, err
	}

	return createTestInstruction(app.computableBuilder, app.stackFrame, insApp, app.baseDir), nil
}
