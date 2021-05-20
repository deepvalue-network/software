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

// NewScript creates a new script instance
func NewScript(
	parser parsers.Parser,
	linker linkers.Linker,
) Script {
	application := NewApplication()
	computableBuilder := computable.NewBuilder()
	programBuilder := linkers.NewProgramBuilder()
	languageBuilder := linkers.NewLanguageBuilder()
	return createScript(
		application,
		computableBuilder,
		programBuilder,
		languageBuilder,
		linker,
	)
}

// NewLanguage creates a new language instance
func NewLanguage(
	linker linkers.Linker,
	lexerAdapterBuilder lexers.AdapterBuilder,
	events []lexers.Event,
) Language {
	composerBuilder := composers.NewBuilder(linker)
	machineStateFactory := machines.NewLanguageStateFactory()
	stackFrameBuilder := stackframes.NewBuilder()
	machineLangTestInsBuilder := machines.NewLanguageTestInstructionBuilder(lexerAdapterBuilder, events)
	machineLangInsBuilder := machines.NewLanguageInstructionBuilder(lexerAdapterBuilder, events)
	return createLanguage(
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

// Script represents a script interpreter
type Script interface {
	Execute(script linkers.Script) (linkers.Application, error)
}

// Language represents a language interpreter
type Language interface {
	Execute(linkedLangDef linkers.LanguageDefinition, input map[string]computable.Value) (linkers.Application, error)
	TestsAll(linkedLangDef linkers.LanguageDefinition) error
	TestByNames(linkedLangDef linkers.LanguageDefinition, names []string) error
}

// Application represents an application interpreter
type Application interface {
	Execute(linkedApp linkers.Application, input map[string]computable.Value) (stackframes.StackFrame, error)
	TestsAll(linkedApp linkers.Application) error
	TestByNames(linkedApp linkers.Application, names []string) error
}
