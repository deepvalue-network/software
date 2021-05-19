package composers

import (
	"errors"

	"github.com/deepvalue-network/software/pangolin/domain/interpreters/stackframes"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

type instructionAdapterBuilder struct {
	testInstructionBuilder     parsers.TestInstructionBuilder
	assertBuilder              parsers.AssertBuilder
	readFileBuilder            parsers.ReadFileBuilder
	relativePathBuilder        parsers.RelativePathBuilder
	labelInstructionBuilder    parsers.LabelInstructionBuilder
	instructionBuilder         parsers.InstructionBuilder
	exitBuilder                parsers.ExitBuilder
	callBuilder                parsers.CallBuilder
	printBuilder               parsers.PrintBuilder
	operationBuilder           parsers.OperationBuilder
	arythmeticBuilder          parsers.ArythmeticBuilder
	relationalBuilder          parsers.RelationalBuilder
	logicalBuilder             parsers.LogicalBuilder
	variableBuilder            parsers.VariableBuilder
	declarationBuilder         parsers.DeclarationBuilder
	typeBuilder                parsers.TypeBuilder
	valueRepresentationBuilder parsers.ValueRepresentationBuilder
	valueBuilder               parsers.ValueBuilder
	numericValueBuilder        parsers.NumericValueBuilder
	assignmentBuilder          parsers.AssignmentBuilder
	concatenationBuilder       parsers.ConcatenationBuilder
	remainingOperationBuilder  parsers.RemainingOperationBuilder
	standardOperationBuilder   parsers.StandardOperationBuilder
	jumpBuilder                parsers.JumpBuilder
	stackFrameBuilder          parsers.StackFrameBuilder
	indexBuilder               parsers.IndexBuilder
	skipBuilder                parsers.SkipBuilder
	intPointBuilder            parsers.IntPointerBuilder
	localStackFrame            stackframes.StackFrame
	stackFrame                 stackframes.StackFrame
}

func createInstructionAdapterBuilder(
	testInstructionBuilder parsers.TestInstructionBuilder,
	assertBuilder parsers.AssertBuilder,
	readFileBuilder parsers.ReadFileBuilder,
	relativePathBuilder parsers.RelativePathBuilder,
	labelInstructionBuilder parsers.LabelInstructionBuilder,
	instructionBuilder parsers.InstructionBuilder,
	exitBuilder parsers.ExitBuilder,
	callBuilder parsers.CallBuilder,
	printBuilder parsers.PrintBuilder,
	operationBuilder parsers.OperationBuilder,
	arythmeticBuilder parsers.ArythmeticBuilder,
	relationalBuilder parsers.RelationalBuilder,
	logicalBuilder parsers.LogicalBuilder,
	variableBuilder parsers.VariableBuilder,
	declarationBuilder parsers.DeclarationBuilder,
	typeBuilder parsers.TypeBuilder,
	valueRepresentationBuilder parsers.ValueRepresentationBuilder,
	valueBuilder parsers.ValueBuilder,
	numericValueBuilder parsers.NumericValueBuilder,
	assignmentBuilder parsers.AssignmentBuilder,
	concatenationBuilder parsers.ConcatenationBuilder,
	remainingOperationBuilder parsers.RemainingOperationBuilder,
	standardOperationBuilder parsers.StandardOperationBuilder,
	jumpBuilder parsers.JumpBuilder,
	stackFrameBuilder parsers.StackFrameBuilder,
	indexBuilder parsers.IndexBuilder,
	skipBuilder parsers.SkipBuilder,
	intPointBuilder parsers.IntPointerBuilder,
) InstructionAdapterBuilder {
	out := instructionAdapterBuilder{
		testInstructionBuilder:     testInstructionBuilder,
		assertBuilder:              assertBuilder,
		readFileBuilder:            readFileBuilder,
		relativePathBuilder:        relativePathBuilder,
		labelInstructionBuilder:    labelInstructionBuilder,
		instructionBuilder:         instructionBuilder,
		exitBuilder:                exitBuilder,
		callBuilder:                callBuilder,
		printBuilder:               printBuilder,
		operationBuilder:           operationBuilder,
		arythmeticBuilder:          arythmeticBuilder,
		relationalBuilder:          relationalBuilder,
		logicalBuilder:             logicalBuilder,
		variableBuilder:            variableBuilder,
		declarationBuilder:         declarationBuilder,
		typeBuilder:                typeBuilder,
		valueRepresentationBuilder: valueRepresentationBuilder,
		valueBuilder:               valueBuilder,
		numericValueBuilder:        numericValueBuilder,
		assignmentBuilder:          assignmentBuilder,
		concatenationBuilder:       concatenationBuilder,
		remainingOperationBuilder:  remainingOperationBuilder,
		standardOperationBuilder:   standardOperationBuilder,
		jumpBuilder:                jumpBuilder,
		stackFrameBuilder:          stackFrameBuilder,
		indexBuilder:               indexBuilder,
		skipBuilder:                skipBuilder,
		intPointBuilder:            intPointBuilder,
		localStackFrame:            nil,
		stackFrame:                 nil,
	}

	return &out
}

// Create initializes the builder
func (app *instructionAdapterBuilder) Create() InstructionAdapterBuilder {
	return createInstructionAdapterBuilder(
		app.testInstructionBuilder,
		app.assertBuilder,
		app.readFileBuilder,
		app.relativePathBuilder,
		app.labelInstructionBuilder,
		app.instructionBuilder,
		app.exitBuilder,
		app.callBuilder,
		app.printBuilder,
		app.operationBuilder,
		app.arythmeticBuilder,
		app.relationalBuilder,
		app.logicalBuilder,
		app.variableBuilder,
		app.declarationBuilder,
		app.typeBuilder,
		app.valueRepresentationBuilder,
		app.valueBuilder,
		app.numericValueBuilder,
		app.assignmentBuilder,
		app.concatenationBuilder,
		app.remainingOperationBuilder,
		app.standardOperationBuilder,
		app.jumpBuilder,
		app.stackFrameBuilder,
		app.indexBuilder,
		app.skipBuilder,
		app.intPointBuilder,
	)
}

// WithLocalStackFrame adds a local stackFrame to the builder
func (app *instructionAdapterBuilder) WithLocalStackFrame(localStackFrame stackframes.StackFrame) InstructionAdapterBuilder {
	app.localStackFrame = localStackFrame
	return app
}

// WithStackFrame adds a stackFrame to the builder
func (app *instructionAdapterBuilder) WithStackFrame(stackFrame stackframes.StackFrame) InstructionAdapterBuilder {
	app.stackFrame = stackFrame
	return app
}

// Now builds a new InstructionAdapter insatnce
func (app *instructionAdapterBuilder) Now() (InstructionAdapter, error) {
	if app.localStackFrame != nil {
		return nil, errors.New("the local StackFrame is mandatory in order to build an InstructionAdapter instance")
	}

	if app.stackFrame != nil {
		return nil, errors.New("the stackFrame is mandatory in order to build an InstructionAdapter instance")
	}

	return createInstructionAdapter(
		app.testInstructionBuilder,
		app.assertBuilder,
		app.readFileBuilder,
		app.relativePathBuilder,
		app.labelInstructionBuilder,
		app.instructionBuilder,
		app.exitBuilder,
		app.callBuilder,
		app.printBuilder,
		app.operationBuilder,
		app.arythmeticBuilder,
		app.relationalBuilder,
		app.logicalBuilder,
		app.variableBuilder,
		app.declarationBuilder,
		app.typeBuilder,
		app.valueRepresentationBuilder,
		app.valueBuilder,
		app.numericValueBuilder,
		app.assignmentBuilder,
		app.concatenationBuilder,
		app.remainingOperationBuilder,
		app.standardOperationBuilder,
		app.jumpBuilder,
		app.stackFrameBuilder,
		app.indexBuilder,
		app.skipBuilder,
		app.intPointBuilder,
		app.localStackFrame,
		app.stackFrame,
	), nil
}
