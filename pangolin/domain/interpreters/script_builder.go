package interpreters

import (
	"errors"

	"github.com/deepvalue-network/software/pangolin/domain/interpreters/stackframes"
	"github.com/deepvalue-network/software/pangolin/domain/lexers"
	"github.com/deepvalue-network/software/pangolin/domain/linkers"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/variable/value/computable"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

type scriptBuilder struct {
	stackframeBuilder    stackframes.Builder
	computableBuilder    computable.Builder
	programBuilder       linkers.ProgramBuilder
	linkerLnguageBuilder linkers.LanguageBuilder
	languageBuilder      LanguageBuilder
	application          Application
	parser               parsers.Parser
	linker               linkers.Linker
	events               []lexers.Event
}

func createScriptBuilder(
	stackframeBuilder stackframes.Builder,
	computableBuilder computable.Builder,
	programBuilder linkers.ProgramBuilder,
	linkerLnguageBuilder linkers.LanguageBuilder,
	languageBuilder LanguageBuilder,
	application Application,
) ScriptBuilder {
	out := scriptBuilder{
		stackframeBuilder:    stackframeBuilder,
		computableBuilder:    computableBuilder,
		programBuilder:       programBuilder,
		linkerLnguageBuilder: linkerLnguageBuilder,
		languageBuilder:      languageBuilder,
		application:          application,
		parser:               nil,
		linker:               nil,
		events:               nil,
	}

	return &out
}

// Create initializes the builder
func (app *scriptBuilder) Create() ScriptBuilder {
	return createScriptBuilder(
		app.stackframeBuilder,
		app.computableBuilder,
		app.programBuilder,
		app.linkerLnguageBuilder,
		app.languageBuilder,
		app.application,
	)
}

// WithParser adds a parser to the builder
func (app *scriptBuilder) WithParser(parser parsers.Parser) ScriptBuilder {
	app.parser = parser
	return app
}

// WithLinker adds a linker to the builder
func (app *scriptBuilder) WithLinker(linker linkers.Linker) ScriptBuilder {
	app.linker = linker
	return app
}

// WithEvents add events to the builder
func (app *scriptBuilder) WithEvents(events []lexers.Event) ScriptBuilder {
	app.events = events
	return app
}

// Now builds a new Script instance
func (app *scriptBuilder) Now() (Script, error) {
	if app.parser == nil {
		return nil, errors.New("the parser is mandatory in order to build a Script instance")
	}

	if app.linker == nil {
		return nil, errors.New("the linker is mandatory in order to build a Script instance")
	}

	language, err := app.languageBuilder.Create().WithLinker(app.linker).WithEvents(app.events).Now()
	if err != nil {
		return nil, err
	}

	return createScript(
		app.stackframeBuilder,
		app.computableBuilder,
		app.programBuilder,
		app.linkerLnguageBuilder,
		app.application,
		language,
		app.parser,
		app.linker,
	), nil
}
