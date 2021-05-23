package interpreters

import (
	"errors"

	"github.com/deepvalue-network/software/pangolin/domain/interpreters/composers"
	"github.com/deepvalue-network/software/pangolin/domain/interpreters/machines"
	"github.com/deepvalue-network/software/pangolin/domain/interpreters/stackframes"
	"github.com/deepvalue-network/software/pangolin/domain/lexers"
	"github.com/deepvalue-network/software/pangolin/domain/linkers"
)

type languageBuilder struct {
	lexerAdapterBuilder       lexers.AdapterBuilder
	composerBuilder           composers.Builder
	machineStateFactory       machines.LanguageStateFactory
	stackFrameBuilder         stackframes.Builder
	machineLangTestInsBuilder machines.LanguageTestInstructionBuilder
	machineLangInsBuilder     machines.LanguageInstructionBuilder
	linker                    linkers.Linker
	events                    []lexers.Event
}

func createLanguageBuilder(
	lexerAdapterBuilder lexers.AdapterBuilder,
	composerBuilder composers.Builder,
	machineStateFactory machines.LanguageStateFactory,
	stackFrameBuilder stackframes.Builder,
	machineLangTestInsBuilder machines.LanguageTestInstructionBuilder,
	machineLangInsBuilder machines.LanguageInstructionBuilder,
) LanguageBuilder {
	out := languageBuilder{
		lexerAdapterBuilder:       lexerAdapterBuilder,
		composerBuilder:           composerBuilder,
		machineStateFactory:       machineStateFactory,
		stackFrameBuilder:         stackFrameBuilder,
		machineLangTestInsBuilder: machineLangTestInsBuilder,
		machineLangInsBuilder:     machineLangInsBuilder,
		linker:                    nil,
		events:                    nil,
	}

	return &out
}

// Create initializes the builder
func (app *languageBuilder) Create() LanguageBuilder {
	return createLanguageBuilder(
		app.lexerAdapterBuilder,
		app.composerBuilder,
		app.machineStateFactory,
		app.stackFrameBuilder,
		app.machineLangTestInsBuilder,
		app.machineLangInsBuilder,
	)
}

// WithLinker adds a linker to the builder
func (app *languageBuilder) WithLinker(linker linkers.Linker) LanguageBuilder {
	app.linker = linker
	return app
}

// WithEvents add events to the builder
func (app *languageBuilder) WithEvents(events []lexers.Event) LanguageBuilder {
	app.events = events
	return app
}

// Now builds a new Language instance
func (app *languageBuilder) Now() (Language, error) {
	if app.linker == nil {
		return nil, errors.New("the linker is mandatory in order to build a Language instance")
	}

	return createLanguage(
		app.composerBuilder,
		app.machineStateFactory,
		app.stackFrameBuilder,
		app.machineLangTestInsBuilder,
		app.machineLangInsBuilder,
		app.linker,
		app.events,
	), nil
}
