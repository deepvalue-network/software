package interpreters

import (
	"github.com/deepvalue-network/software/pangolin/domain/interpreters/composers"
	"github.com/deepvalue-network/software/pangolin/domain/interpreters/machines"
	"github.com/deepvalue-network/software/pangolin/domain/interpreters/stackframes"
	"github.com/deepvalue-network/software/pangolin/domain/lexers"
	"github.com/deepvalue-network/software/pangolin/domain/linkers"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction/variable/value/computable"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

const delimiter = "\n++++++++++++++++++++++++++++++++++\n"
const printTestStr = "Test: %s\n"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	testableBuilder := NewTestableBuilder()
	executableBuilder := NewExecutableBuilder()
	return createBuilder(executableBuilder, testableBuilder)
}

// NewTestableBuilder creates a new testable builder instance
func NewTestableBuilder() TestableBuilder {
	testInsMachineBuilder := machines.NewTestInstructionBuilder()
	machineStateFactory := machines.NewLanguageStateFactory()
	machineLangTestInsBuilder := machines.NewLanguageTestInstructionBuilder()
	stackFrameBuilder := stackframes.NewBuilder()
	composerBuilder := composers.NewBuilder()
	executableBuilder := NewExecutableBuilder()
	return createTestableBuilder(
		testInsMachineBuilder,
		machineStateFactory,
		machineLangTestInsBuilder,
		stackFrameBuilder,
		composerBuilder,
		executableBuilder,
	)
}

// NewExecutableBuilder creates a new executable builder instance
func NewExecutableBuilder() ExecutableBuilder {
	insMachineBuilder := machines.NewInstructionBuilder()
	stackFrameBuilder := stackframes.NewBuilder()
	computableBuilder := computable.NewBuilder()
	programBuilder := linkers.NewProgramBuilder()
	testableBuilder := linkers.NewTestableBuilder()
	return createExecutableBuilder(
		insMachineBuilder,
		stackFrameBuilder,
		computableBuilder,
		programBuilder,
		testableBuilder,
	)
}

// Builder represents a builder
type Builder interface {
	Create() Builder
	WithParser(parser parsers.Parser) Builder
	WithLinker(linker linkers.Linker) Builder
	WithEvents(events []lexers.Event) Builder
	Now() (Interpreter, error)
}

// Interpreter represents the interpreter interface
type Interpreter interface {
	Execute(excutable linkers.Executable, input stackframes.StackFrame) (stackframes.StackFrame, error)
	Tests(testable linkers.Testable) error
}

// ExecutableBuilder represents an executable builder
type ExecutableBuilder interface {
	Create() ExecutableBuilder
	WithParser(parser parsers.Parser) ExecutableBuilder
	WithLinker(linker linkers.Linker) ExecutableBuilder
	Now() (Executable, error)
}

// Executable represents the executable interface
type Executable interface {
	Execute(excutable linkers.Executable, input stackframes.StackFrame) (stackframes.StackFrame, error)
	Application(appli linkers.Application, input stackframes.StackFrame) (stackframes.StackFrame, error)
	Script(appli linkers.Script, input stackframes.StackFrame) (stackframes.StackFrame, error)
}

// TestableBuilder represents a testable builder
type TestableBuilder interface {
	Create() TestableBuilder
	WithParser(parser parsers.Parser) TestableBuilder
	WithLinker(linker linkers.Linker) TestableBuilder
	WithEvents(events []lexers.Event) TestableBuilder
	Now() (Testable, error)
}

// Testable represents the testable interface
type Testable interface {
	Execute(testable linkers.Testable) error
	Executable(executable linkers.Executable) error
	Application(appli linkers.Application) error
	Script(script linkers.Script) error
	Language(language linkers.LanguageDefinition) error
}
