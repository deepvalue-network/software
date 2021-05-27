package interpreters

import (
	"errors"

	"github.com/deepvalue-network/software/pangolin/domain/interpreters/machines"
	"github.com/deepvalue-network/software/pangolin/domain/interpreters/stackframes"
	"github.com/deepvalue-network/software/pangolin/domain/linkers"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction/variable/value/computable"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

type executableBuilder struct {
	insMachineBuilder machines.InstructionBuilder
	stackFrameBuilder stackframes.Builder
	computableBuilder computable.Builder
	programBuilder    linkers.ProgramBuilder
	testableBuilder   linkers.TestableBuilder
	parser            parsers.Parser
	linker            linkers.Linker
}

func createExecutableBuilder(
	insMachineBuilder machines.InstructionBuilder,
	stackFrameBuilder stackframes.Builder,
	computableBuilder computable.Builder,
	programBuilder linkers.ProgramBuilder,
	testableBuilder linkers.TestableBuilder,
) ExecutableBuilder {
	out := executableBuilder{
		insMachineBuilder: insMachineBuilder,
		stackFrameBuilder: stackFrameBuilder,
		computableBuilder: computableBuilder,
		programBuilder:    programBuilder,
		testableBuilder:   testableBuilder,
		parser:            nil,
		linker:            nil,
	}

	return &out
}

// Create initializes the builder
func (app *executableBuilder) Create() ExecutableBuilder {
	return createExecutableBuilder(
		app.insMachineBuilder,
		app.stackFrameBuilder,
		app.computableBuilder,
		app.programBuilder,
		app.testableBuilder,
	)
}

// WithParser adds a parser to the builder
func (app *executableBuilder) WithParser(parser parsers.Parser) ExecutableBuilder {
	app.parser = parser
	return app
}

// WithLinker adds a linker to the builder
func (app *executableBuilder) WithLinker(linker linkers.Linker) ExecutableBuilder {
	app.linker = linker
	return app
}

// Now builds a new Executable instance
func (app *executableBuilder) Now() (Executable, error) {
	if app.parser == nil {
		return nil, errors.New("the parser is mandatory in order to build an Executable instance")
	}

	if app.linker == nil {
		return nil, errors.New("the linker is mandatory in order to build an Executable instance")
	}

	return createExecutable(
		app.insMachineBuilder,
		app.stackFrameBuilder,
		app.computableBuilder,
		app.programBuilder,
		app.testableBuilder,
		app.parser,
		app.linker,
	), nil
}
