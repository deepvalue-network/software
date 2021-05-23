package interpreters

import (
	"errors"

	"github.com/deepvalue-network/software/pangolin/domain/lexers"
	"github.com/deepvalue-network/software/pangolin/domain/linkers"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

type builder struct {
	scriptBuilder ScriptBuilder
	application   Application
	parser        parsers.Parser
	linker        linkers.Linker
	events        []lexers.Event
}

func createBuilder(
	scriptBuilder ScriptBuilder,
	application Application,
) Builder {
	out := builder{
		scriptBuilder: scriptBuilder,
		application:   application,
		parser:        nil,
		linker:        nil,
		events:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.scriptBuilder, app.application)
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
	if app.parser == nil {
		return nil, errors.New("the parser is mandatory in order to build an Interpreter instance")
	}

	if app.linker == nil {
		return nil, errors.New("the linker is mandatory in order to build an Interpreter instance")
	}

	script, err := app.scriptBuilder.Create().WithEvents(app.events).WithLinker(app.linker).WithParser(app.parser).Now()
	if err != nil {
		return nil, err
	}

	return createInterpreter(app.application, script), nil
}
