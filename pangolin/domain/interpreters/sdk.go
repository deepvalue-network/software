package interpreters

import (
	"github.com/deepvalue-network/software/pangolin/domain/lexers"
	"github.com/deepvalue-network/software/pangolin/domain/lexers/grammar"
	lexer_parser "github.com/deepvalue-network/software/pangolin/domain/lexers/parser"
	"github.com/deepvalue-network/software/pangolin/domain/linkers"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/remaining"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/standard"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/transform"
	var_variable "github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/variable"
	var_value "github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/variable/value"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/variable/value/computable"
	label_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/applications/labels/label/instructions/instruction"
	language_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions/instruction"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions/instruction/commands"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions/instruction/match"
	language_label_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/labels/label/instructions/instruction"
)

// CallLabelFunc represents a call label func
type CallLabelFunc func(name string) error

// FetchStackFrameFunc represents a fetch stackframe func
type FetchStackFrameFunc func() StackFrame

// NewBuilder creates a new interpreter builder instance
func NewBuilder(machineLanguageBuilder MachineLanguageBuilder) Builder {
	stackFrameBuilder := NewStackFrameBuilder()
	machineBuilder := NewMachineBuilder()
	valueBuilder := computable.NewBuilder()
	return createBuilder(stackFrameBuilder, machineBuilder, machineLanguageBuilder, valueBuilder)
}

// NewMachineLanguageBuilder creates a new machine language builder
func NewMachineLanguageBuilder(
	lexerAdapterBuilder lexers.AdapterBuilder,
	events []lexers.Event,
) MachineLanguageBuilder {
	variableBuilder := var_variable.NewBuilder()
	valueBuilder := var_value.NewBuilder()
	computableValueBuilder := computable.NewBuilder()
	lexerParserApplication := lexer_parser.NewApplication()
	lexerParserBuilder := lexer_parser.NewBuilder()
	grammarRetrieverCriteriaBuilder := grammar.NewRetrieverCriteriaBuilder()
	stackFrameBuilder := NewStackFrameBuilder()
	return createMachineLanguageBuilder(
		variableBuilder,
		valueBuilder,
		computableValueBuilder,
		lexerParserApplication,
		lexerParserBuilder,
		lexerAdapterBuilder,
		grammarRetrieverCriteriaBuilder,
		stackFrameBuilder,
		events,
	)
}

// NewMachineBuilder creates a new machineBuilder instance
func NewMachineBuilder() MachineBuilder {
	computableValueBuilder := computable.NewBuilder()
	return createMachineBuilder(
		computableValueBuilder,
	)
}

// NewStackFrameBuilder creates a new stackFrame builder
func NewStackFrameBuilder() StackFrameBuilder {
	frameBuilder := NewFrameBuilder()
	return createStackFrameBuilder(frameBuilder)
}

// NewFrameBuilder creates a new frameBuilder instance
func NewFrameBuilder() FrameBuilder {
	computableBuilder := computable.NewBuilder()
	computer := createComputer(computableBuilder)
	return createFrameBuilder(computer, computableBuilder)
}

// Builder represets an interpreter builder
type Builder interface {
	Create() Builder
	WithProgram(program linkers.Program) Builder
	Now() (Interpreter, error)
}

// Interpreter represents an interpreter
type Interpreter interface {
	IsScript() bool
	Script() Script
	IsApplication() bool
	Application() Application
	IsLanguage() bool
	Language() Language
}

// Script represents a script interpreter
type Script interface {
	Execute(input map[string]computable.Value) (string, error)
}

// Application represents an application interpreter
type Application interface {
	Execute(input map[string]computable.Value) (StackFrame, error)
}

// Language represents a language interpreter
type Language interface {
	TestsAll() error
	Tests(names []string) error
}

// MachineBuilder represents a machine builder
type MachineBuilder interface {
	Create() MachineBuilder
	WithCallLabelFunc(callLabelFunc CallLabelFunc) MachineBuilder
	WithFetchStackFunc(fetchStackFunc FetchStackFrameFunc) MachineBuilder
	Now() (Machine, error)
}

// Machine represents a machine that receives 1 instruction at a time
type Machine interface {
	Receive(ins instruction.Instruction) error
	ReceiveLbl(lblIns label_instruction.Instruction) (bool, error)
}

// MachineLanguageBuilder represents a machine language builder
type MachineLanguageBuilder interface {
	Create() MachineLanguageBuilder
	WithLanguage(lang linkers.LanguageDefinition) MachineLanguageBuilder
	WithInput(input map[string]computable.Value) MachineLanguageBuilder
	WithFetchStackFunc(fetchStackFunc FetchStackFrameFunc) MachineLanguageBuilder
	WithMachine(machine Machine) MachineLanguageBuilder
	Now() (MachineLanguage, error)
}

// MachineLanguage represents a language machine that receives 1 instruction at a time
type MachineLanguage interface {
	Match(match match.Match) error
	Command(command commands.Command) error
	Receive(ins language_instruction.Instruction) error
	ReceiveLbl(lblIns language_label_instruction.Instruction) (bool, error)
}

// StackFrameBuilder represents a stackframe builder
type StackFrameBuilder interface {
	Create() StackFrameBuilder
	WithVariables(variables map[string]computable.Value) StackFrameBuilder
	Now() StackFrame
}

// StackFrame represents a stackframe
type StackFrame interface {
	Push()
	Pop() error
	Index() int
	Skip(index int) error
	Current() Frame
}

// FrameBuilder represents a frame builder
type FrameBuilder interface {
	Create() FrameBuilder
	WithVariables(variables map[string]computable.Value) FrameBuilder
	Now() Frame
}

// Frame represents a frame
type Frame interface {
	Standard(first string, second string, result string, operation standard.Operation) error
	Remaining(first string, second string, result string, remaining string, operation remaining.Operation) error
	Transform(input string, result string, operation transform.Operation) error
	PushTo(name string, frame Frame) error
	Insert(operation var_variable.Variable) error
	Update(operation var_variable.Variable) error
	UpdateValue(name string, val computable.Value) error
	Delete(name string) error
	Fetch(name string) (computable.Value, error)
	Stop()
	IsStopped() bool
}

// Computer represents a computer
type Computer interface {
	Add(first computable.Value, second computable.Value) (computable.Value, error)
	Substract(first computable.Value, second computable.Value) (computable.Value, error)
	Multiply(first computable.Value, second computable.Value) (computable.Value, error)
	Divide(first computable.Value, second computable.Value) (computable.Value, computable.Value, error)
	IsLessThan(first computable.Value, second computable.Value) (computable.Value, error)
	IsEqual(first computable.Value, second computable.Value) (computable.Value, error)
	IsNotEqual(first computable.Value, second computable.Value) (computable.Value, error)
	And(first computable.Value, second computable.Value) (computable.Value, error)
	Or(first computable.Value, second computable.Value) (computable.Value, error)
	Concat(first computable.Value, second computable.Value) (computable.Value, error)
	Match(pattern computable.Value, value computable.Value) ([]computable.Value, computable.Value, error)
}
