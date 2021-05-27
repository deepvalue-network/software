package interpreters

import (
	"errors"

	"github.com/deepvalue-network/software/pangolin/domain/interpreters/composers"
	"github.com/deepvalue-network/software/pangolin/domain/interpreters/machines"
	"github.com/deepvalue-network/software/pangolin/domain/interpreters/stackframes"
	"github.com/deepvalue-network/software/pangolin/domain/lexers"
	"github.com/deepvalue-network/software/pangolin/domain/linkers"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

type testableBuilder struct {
	testInsMachineBuilder     machines.TestInstructionBuilder
	machineStateFactory       machines.LanguageStateFactory
	machineLangTestInsBuilder machines.LanguageTestInstructionBuilder
	stackFrameBuilder         stackframes.Builder
	composerBuilder           composers.Builder
	executableBuilder         ExecutableBuilder
	parser                    parsers.Parser
	linker                    linkers.Linker
	events                    []lexers.Event
}

func createTestableBuilder(
	testInsMachineBuilder machines.TestInstructionBuilder,
	machineStateFactory machines.LanguageStateFactory,
	machineLangTestInsBuilder machines.LanguageTestInstructionBuilder,
	stackFrameBuilder stackframes.Builder,
	composerBuilder composers.Builder,
	executableBuilder ExecutableBuilder,
) TestableBuilder {
	out := testableBuilder{
		testInsMachineBuilder:     testInsMachineBuilder,
		machineStateFactory:       machineStateFactory,
		machineLangTestInsBuilder: machineLangTestInsBuilder,
		stackFrameBuilder:         stackFrameBuilder,
		composerBuilder:           composerBuilder,
		executableBuilder:         executableBuilder,
		parser:                    nil,
		linker:                    nil,
		events:                    nil,
	}

	return &out
}

// Create initializes the builder
func (app *testableBuilder) Create() TestableBuilder {
	return createTestableBuilder(
		app.testInsMachineBuilder,
		app.machineStateFactory,
		app.machineLangTestInsBuilder,
		app.stackFrameBuilder,
		app.composerBuilder,
		app.executableBuilder,
	)
}

// WithParser adds a parser to the builder
func (app *testableBuilder) WithParser(parser parsers.Parser) TestableBuilder {
	app.parser = parser
	return app
}

// WithLinker adds a linker to the builder
func (app *testableBuilder) WithLinker(linker linkers.Linker) TestableBuilder {
	app.linker = linker
	return app
}

// WithEvents add events to the builder
func (app *testableBuilder) WithEvents(events []lexers.Event) TestableBuilder {
	app.events = events
	return app
}

// Now builds a new Script instance
func (app *testableBuilder) Now() (Testable, error) {
	if app.linker == nil {
		return nil, errors.New("the linker is mandatory in order to build a Testable instance")
	}

	executableBuilder := app.executableBuilder.Create().WithLinker(app.linker)
	if app.parser != nil {
		executableBuilder.WithParser(app.parser)
	}

	executable, err := executableBuilder.Now()
	if err != nil {
		return nil, err
	}

	return createTestable(
		app.testInsMachineBuilder,
		app.machineStateFactory,
		app.machineLangTestInsBuilder,
		app.stackFrameBuilder,
		app.composerBuilder,
		executable,
		app.linker,
		app.events,
	), nil
}
