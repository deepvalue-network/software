package machines

import (
	"github.com/deepvalue-network/software/pangolin/domain/interpreters/composers"
	"github.com/deepvalue-network/software/pangolin/domain/interpreters/stackframes"
	"github.com/deepvalue-network/software/pangolin/domain/lexers"
	"github.com/deepvalue-network/software/pangolin/domain/lexers/grammar"
	lexer_parser "github.com/deepvalue-network/software/pangolin/domain/lexers/parser"
	"github.com/deepvalue-network/software/pangolin/domain/linkers"
	application_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction"
	var_variable "github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/variable"
	var_value "github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/variable/value"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/variable/value/computable"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/labels"
	label_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/applications/labels/label/instructions/instruction"
	test_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/applications/tests/test/instructions/instruction"
	language_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions/instruction"
	language_label_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/labels/label/instructions/instruction"
	language_test_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/tests/test/instructions/instruction"
)

// CallLabelByNameFn represents a func to call a label by name
type CallLabelByNameFn func(name string) error

// NewLanguageTestInstructionBuilder creates a language test instruction builder
func NewLanguageTestInstructionBuilder(
	lexerAdapterBuilder lexers.AdapterBuilder,
) LanguageTestInstructionBuilder {
	langCommonInsBuilder := NewLanguageInstructionCommonBuilder(lexerAdapterBuilder)
	testInsAppBuilder := NewTestInstructionBuilder()
	langInsBuilder := NewLanguageInstructionBuilder(lexerAdapterBuilder)
	return createLanguageTestInstructionBuilder(
		langCommonInsBuilder,
		testInsAppBuilder,
		langInsBuilder,
	)
}

// NewLanguageInstructionBuilder creates a new language instruction builder
func NewLanguageInstructionBuilder(
	lexerAdapterBuilder lexers.AdapterBuilder,
) LanguageInstructionBuilder {
	variableBuilder := var_variable.NewBuilder()
	valueBuilder := var_value.NewBuilder()
	computableValueBuilder := computable.NewBuilder()
	langCommonInsBuilder := NewLanguageInstructionCommonBuilder(lexerAdapterBuilder)
	insAppBuilder := NewInstructionBuilder()
	return createLanguageInstructionBuilder(
		variableBuilder,
		valueBuilder,
		computableValueBuilder,
		langCommonInsBuilder,
		insAppBuilder,
	)
}

// NewLanguageInstructionCommonBuilder creates a new language instruction common builder
func NewLanguageInstructionCommonBuilder(
	lexerAdapterBuilder lexers.AdapterBuilder,
) LanguageInstructionCommonBuilder {
	insAppBuilder := NewInstructionBuilder()
	grammarRetrieverCriteriaBuilder := grammar.NewRetrieverCriteriaBuilder()
	lexerParserApplication := lexer_parser.NewApplication()
	lexerParserBuilder := lexer_parser.NewBuilder()
	return createLanguageInstructionCommonBuilder(
		insAppBuilder,
		grammarRetrieverCriteriaBuilder,
		lexerParserApplication,
		lexerParserBuilder,
		lexerAdapterBuilder,
	)
}

// NewTestInstructionBuilder creates a new test instruction builder
func NewTestInstructionBuilder() TestInstructionBuilder {
	instructionBuilder := NewInstructionBuilder()
	computableValueBuilder := computable.NewBuilder()
	return createTestInstructionBuilder(instructionBuilder, computableValueBuilder)
}

// NewInstructionBuilder creates a new instruction builder
func NewInstructionBuilder() InstructionBuilder {
	computableValueBuilder := computable.NewBuilder()
	return createInstructionBuilder(computableValueBuilder)
}

// NewLanguageStateFactory creates a new language state factory
func NewLanguageStateFactory() LanguageStateFactory {
	return createLanguageStateFactory()
}

// LanguageTestInstructionBuilder represents a language test instruction builder
type LanguageTestInstructionBuilder interface {
	Create() LanguageTestInstructionBuilder
	WithComposer(composerApp composers.Composer) LanguageTestInstructionBuilder
	WithLanguage(langDef linkers.LanguageDefinition) LanguageTestInstructionBuilder
	WithStackFrame(stackFrame stackframes.StackFrame) LanguageTestInstructionBuilder
	WithState(state LanguageState) LanguageTestInstructionBuilder
	WithEvents(events []lexers.Event) LanguageTestInstructionBuilder
	Now() (LanguageTestInstruction, error)
}

// LanguageTestInstruction represents a language test instruction machine
type LanguageTestInstruction interface {
	Receive(ins language_test_instruction.Instruction) (bool, error)
}

// LanguageInstructionBuilder represents a language instruction builder
type LanguageInstructionBuilder interface {
	Create() LanguageInstructionBuilder
	WithComposer(composerApp composers.Composer) LanguageInstructionBuilder
	WithLanguage(langDef linkers.LanguageDefinition) LanguageInstructionBuilder
	WithStackFrame(stackFrame stackframes.StackFrame) LanguageInstructionBuilder
	WithState(state LanguageState) LanguageInstructionBuilder
	WithEvents(events []lexers.Event) LanguageInstructionBuilder
	Now() (LanguageInstruction, error)
}

// LanguageInstruction represents a language instruction machine that receives 1 instruction at a time
type LanguageInstruction interface {
	Receive(ins language_instruction.Instruction) error
	ReceiveLbl(lblIns language_label_instruction.Instruction) (bool, error)
}

// LanguageInstructionCommonBuilder represents a language instruction common builder
type LanguageInstructionCommonBuilder interface {
	Create() LanguageInstructionCommonBuilder
	WithLanguage(langDef linkers.LanguageDefinition) LanguageInstructionCommonBuilder
	WithCallLabelFn(labelFn CallLabelByNameFn) LanguageInstructionCommonBuilder
	WithStackFrame(stackFrame stackframes.StackFrame) LanguageInstructionCommonBuilder
	WithState(state LanguageState) LanguageInstructionCommonBuilder
	WithEvents(events []lexers.Event) LanguageInstructionCommonBuilder
	Now() (LanguageInstructionCommon, error)
}

// LanguageInstructionCommon represents a language instruction common machine that receives 1 instruction at a time
type LanguageInstructionCommon interface {
	Receive(ins language_instruction.CommonInstruction) error
}

// TestInstructionBuilder represents a test instruction builder
type TestInstructionBuilder interface {
	Create() TestInstructionBuilder
	WithCallLabelFn(labelFn CallLabelByNameFn) TestInstructionBuilder
	WithLabels(labels labels.Labels) TestInstructionBuilder
	WithStackFrame(stackFrame stackframes.StackFrame) TestInstructionBuilder
	WithBaseDir(baseDir string) TestInstructionBuilder
	Now() (TestInstruction, error)
}

// TestInstruction represents a test instruction machine
type TestInstruction interface {
	Receive(ins test_instruction.Instruction) (bool, error)
}

// InstructionBuilder represents an instruction machine builder
type InstructionBuilder interface {
	Create() InstructionBuilder
	WithCallLabelFn(labelFn CallLabelByNameFn) InstructionBuilder
	WithLabels(labels labels.Labels) InstructionBuilder
	WithStackFrame(stackFrame stackframes.StackFrame) InstructionBuilder
	Now() (Instruction, error)
}

// Instruction represents an instruction machine that receives 1 instruction at a time
type Instruction interface {
	Receive(ins application_instruction.Instruction) error
	ReceiveLbl(lblIns label_instruction.Instruction) (bool, error)
}

// LanguageStateFactory represents a language state factory
type LanguageStateFactory interface {
	Create() LanguageState
}

// LanguageState represents a machine language state
type LanguageState interface {
	ChangeCurrentToken(tok lexers.NodeTree)
	CurrentToken() lexers.NodeTree
}
