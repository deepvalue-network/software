package interpreters

import (
	"github.com/deepvalue-network/software/pangolin/domain/lexers"
	"github.com/deepvalue-network/software/pangolin/domain/linkers"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

type builder struct {
	executableBuilder ExecutableBuilder
	testableBuilder   TestableBuilder
	parser            parsers.Parser
	linker            linkers.Linker
	events            []lexers.Event
}

func createBuilder(
	executableBuilder ExecutableBuilder,
	testableBuilder TestableBuilder,
) Builder {
	out := builder{
		executableBuilder: executableBuilder,
		testableBuilder:   testableBuilder,
		parser:            nil,
		linker:            nil,
		events:            nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.executableBuilder,
		app.testableBuilder,
	)
}

// WithParser adds a parser to the builder
func (app *builder) WithParser(parser parsers.Parser) Builder {
	app.parser = parser
	return app
}

// WithLinker adds a linker to the builder
func (app *builder) WithLinker(linker linkers.Linker) Builder {
	app.linker = linker
	return app
}

// WithEvents add events to the builder
func (app *builder) WithEvents(events []lexers.Event) Builder {
	app.events = events
	return app
}

// Now builds a new Interpreter instance
func (app *builder) Now() (Interpreter, error) {
	executableBuilder := app.executableBuilder.Create()
	testableBuilder := app.testableBuilder.Create()
	if app.linker != nil {
		executableBuilder.WithLinker(app.linker)
		testableBuilder.WithLinker(app.linker)
	}

	if app.parser != nil {
		executableBuilder.WithParser(app.parser)
		testableBuilder.WithParser(app.parser)
	}

	if app.events != nil {
		testableBuilder.WithEvents(app.events)
	}

	executable, err := executableBuilder.Now()
	if err != nil {
		return nil, err
	}

	testable, err := testableBuilder.Now()
	if err != nil {
		return nil, err
	}

	return createInterpreter(testable, executable), nil
}
