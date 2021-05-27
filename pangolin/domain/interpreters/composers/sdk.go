package composers

import (
	"github.com/deepvalue-network/software/pangolin/domain/interpreters/stackframes"
	"github.com/deepvalue-network/software/pangolin/domain/linkers"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/commands"
	command_labels "github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/commands/labels"
	command_mains "github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/commands/mains"
	command_tests "github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/commands/tests"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

// NewBuilder creates a new composer builder
func NewBuilder() Builder {
	instructionAdapterBuilder := NewInstructionAdapterBuilder()
	stackFrameBuilder := stackframes.NewBuilder()
	programBuilder := parsers.NewProgramBuilder()
	testableBuilder := parsers.NewTestableBuilder()
	executableBuilder := parsers.NewExecutableBuilder()
	languageApplicationBuilder := parsers.NewLanguageApplicationBuilder()
	languageDefinitionBuilder := parsers.NewLanguageDefinitionBuilder()
	languageValueBuilder := parsers.NewLanguageValueBuilder()
	scriptBuilder := parsers.NewScriptBuilder()
	scriptValueBuilder := parsers.NewScriptValueBuilder()
	patternMatchBuilder := parsers.NewPatternMatchBuilder()
	patternLabelsBuilder := parsers.NewPatternLabelsBuilder()
	applicationBuilder := parsers.NewApplicationBuilder()
	testSectionBuilder := parsers.NewTestSectionBuilder()
	testDeclarationBuilder := parsers.NewTestDeclarationBuilder()
	headSectionBuilder := parsers.NewHeadSectionBuilder()
	headValueBuilder := parsers.NewHeadValueBuilder()
	labelSectionBuilder := parsers.NewLabelSectionBuilder()
	labelDeclarationBuilder := parsers.NewLabelDeclarationBuilder()
	mainSectionBuilder := parsers.NewMainSectionBuilder()

	return createBuilder(
		instructionAdapterBuilder,
		stackFrameBuilder,
		programBuilder,
		testableBuilder,
		executableBuilder,
		applicationBuilder,
		labelSectionBuilder,
		mainSectionBuilder,
		testSectionBuilder,
		languageApplicationBuilder,
		languageDefinitionBuilder,
		languageValueBuilder,
		patternMatchBuilder,
		patternLabelsBuilder,
		scriptBuilder,
		scriptValueBuilder,
		headSectionBuilder,
		headValueBuilder,
		testDeclarationBuilder,
		labelDeclarationBuilder,
	)
}

// NewInstructionAdapterBuilder creates a new instruction adapter builder
func NewInstructionAdapterBuilder() InstructionAdapterBuilder {
	relativePathBuilder := parsers.NewRelativePathBuilder()
	testInstructionBuilder := parsers.NewTestInstructionBuilder()
	assertBuilder := parsers.NewAssertBuilder()
	readFileBuilder := parsers.NewReadFileBuilder()
	labelInstructionBuilder := parsers.NewLabelInstructonBuilder()
	instructionBuilder := parsers.NewInstructionBuilder()
	variableBuilder := parsers.NewVariableBuilder()
	concatenationBuilder := parsers.NewConcatenationBuilder()
	declarationBuilder := parsers.NewDeclarationBuilder()
	assignmentBuilder := parsers.NewAssignmentBuilder()
	valueRepresentationBuilder := parsers.NewValueRepresentationBuilder()
	valueBuilder := parsers.NewValueBuilder()
	numericValueBuilder := parsers.NewNumericValueBuilder()
	typeBuilder := parsers.NewTypeBuilder()
	operationBuilder := parsers.NewOperationBuilder()
	arythmeticBuilder := parsers.NewArythmeticBuilder()
	relationalBuilder := parsers.NewRelationalBuilder()
	logicalBuilder := parsers.NewLogicalBuilder()
	standardOperationBuilder := parsers.NewStandardOperationBuilder()
	remainingOperationBuilder := parsers.NewRemainingOperationBuilder()
	printBuilder := parsers.NewPrintBuilder()
	jumpBuilder := parsers.NewJumpBuilder()
	exitBuilder := parsers.NewExitBuilder()
	callBuilder := parsers.NewCallBuilder()
	stackFrameBuilder := parsers.NewStackFrameBuilder()
	indexBuilder := parsers.NewIndexBuilder()
	skipBuilder := parsers.NewSkipBuilder()
	intPointerBuilder := parsers.NewIntPointerBuilder()

	return createInstructionAdapterBuilder(
		testInstructionBuilder,
		assertBuilder,
		readFileBuilder,
		relativePathBuilder,
		labelInstructionBuilder,
		instructionBuilder,
		exitBuilder,
		callBuilder,
		printBuilder,
		operationBuilder,
		arythmeticBuilder,
		relationalBuilder,
		logicalBuilder,
		variableBuilder,
		declarationBuilder,
		typeBuilder,
		valueRepresentationBuilder,
		valueBuilder,
		numericValueBuilder,
		assignmentBuilder,
		concatenationBuilder,
		remainingOperationBuilder,
		standardOperationBuilder,
		jumpBuilder,
		stackFrameBuilder,
		indexBuilder,
		skipBuilder,
		intPointerBuilder,
	)
}

// Builder represents a composer builder
type Builder interface {
	Create() Builder
	WithLinker(linker linkers.Linker) Builder
	WithStackFrame(stackFrame stackframes.StackFrame) Builder
	Now() (Composer, error)
}

// Composer represents a composer
type Composer interface {
	Receive(command commands.Command) error
	Now() (linkers.Application, error)
}

// InstructionAdapterBuilder represents an instruction adapter builder
type InstructionAdapterBuilder interface {
	Create() InstructionAdapterBuilder
	WithLocalStackFrame(localStackFrame stackframes.StackFrame) InstructionAdapterBuilder
	WithStackFrame(stackFrame stackframes.StackFrame) InstructionAdapterBuilder
	Now() (InstructionAdapter, error)
}

// InstructionAdapter represents an instruction adapter
type InstructionAdapter interface {
	Test(ins command_tests.Instruction) ([]parsers.TestInstruction, error)
	Label(ins command_labels.Instruction) ([]parsers.LabelInstruction, error)
	Application(ins command_mains.Instruction) ([]parsers.Instruction, error)
}
