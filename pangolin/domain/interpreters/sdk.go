package interpreters

import (
	"github.com/deepvalue-network/software/pangolin/domain/interpreters/composers"
	"github.com/deepvalue-network/software/pangolin/domain/interpreters/machines"
	"github.com/deepvalue-network/software/pangolin/domain/interpreters/stackframes"
	"github.com/deepvalue-network/software/pangolin/domain/lexers"
	"github.com/deepvalue-network/software/pangolin/domain/linkers"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/variable/value/computable"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

const delimiter = "\n++++++++++++++++++++++++++++++++++\n"
const printTestStr = "Test: %s\n"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	scriptBuilder := NewScriptBuilder()
	application := NewApplication()
	return createBuilder(scriptBuilder, application)
}

// NewScriptBuilder creates a new script builder
func NewScriptBuilder() ScriptBuilder {
	computableBuilder := computable.NewBuilder()
	programBuilder := linkers.NewProgramBuilder()
	linkerLnguageBuilder := linkers.NewLanguageBuilder()
	languageBuilder := NewLanguageBuilder()
	application := NewApplication()
	return createScriptBuilder(
		computableBuilder,
		programBuilder,
		linkerLnguageBuilder,
		languageBuilder,
		application,
	)
}

// NewLanguageBuilder creates a new language builder
func NewLanguageBuilder() LanguageBuilder {
	lexerAdapterBuilder := lexers.NewAdapterBuilder()
	composerBuilder := composers.NewBuilder()
	machineStateFactory := machines.NewLanguageStateFactory()
	stackFrameBuilder := stackframes.NewBuilder()
	machineLangTestInsBuilder := machines.NewLanguageTestInstructionBuilder(lexerAdapterBuilder)
	machineLangInsBuilder := machines.NewLanguageInstructionBuilder(lexerAdapterBuilder)
	return createLanguageBuilder(
		lexerAdapterBuilder,
		composerBuilder,
		machineStateFactory,
		stackFrameBuilder,
		machineLangTestInsBuilder,
		machineLangInsBuilder,
	)
}

// NewApplication creates a new application instance
func NewApplication() Application {
	insMachineBuilder := machines.NewInstructionBuilder()
	testInsMachineBuilder := machines.NewTestInstructionBuilder()
	stackFrameBuilder := stackframes.NewBuilder()
	return createApplication(insMachineBuilder, testInsMachineBuilder, stackFrameBuilder)
}

// Builder represents an interpreter builder
type Builder interface {
	Create() Builder
	WithParser(parser parsers.Parser) Builder
	WithLinker(linker linkers.Linker) Builder
	WithEvents(events []lexers.Event) Builder
	Now() (Interpreter, error)
}

// Interpreter represents an interpreter
type Interpreter interface {
	Execute(excutable linkers.Executable, input map[string]computable.Value) (stackframes.StackFrame, error)
	Tests(excutable linkers.Executable) error
}

// ScriptBuilder represents a script builder
type ScriptBuilder interface {
	Create() ScriptBuilder
	WithParser(parser parsers.Parser) ScriptBuilder
	WithLinker(linker linkers.Linker) ScriptBuilder
	WithEvents(events []lexers.Event) ScriptBuilder
	Now() (Script, error)
}

// Script represents a script interpreter
type Script interface {
	Execute(script linkers.Script) (linkers.Application, error)
	Tests(script linkers.Script) error
}

// LanguageBuilder represents a language builder
type LanguageBuilder interface {
	Create() LanguageBuilder
	WithLinker(linker linkers.Linker) LanguageBuilder
	WithEvents(events []lexers.Event) LanguageBuilder
	Now() (Language, error)
}

// Language represents a language interpreter
type Language interface {
	Execute(linkedLangDef linkers.LanguageDefinition, input map[string]computable.Value) (linkers.Application, error)
	Tests(linkedLangDef linkers.LanguageDefinition) error
}

// Application represents an application interpreter
type Application interface {
	Execute(linkedApp linkers.Application, input map[string]computable.Value) (stackframes.StackFrame, error)
	Tests(linkedApp linkers.Application) error
}
