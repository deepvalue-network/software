package languages

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/definitions"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

type adapter struct {
	definitionAdapter  definitions.Adapter
	applicationAdapter applications.Adapter
	builder            Builder
}

func createAdapter(
	definitionAdapter definitions.Adapter,
	applicationAdapter applications.Adapter,
	builder Builder,
) Adapter {
	out := adapter{
		definitionAdapter:  definitionAdapter,
		applicationAdapter: applicationAdapter,
		builder:            builder,
	}

	return &out
}

// ToLanguage converts a parsed language to language instance
func (app *adapter) ToLanguage(parsed parsers.Language) (Language, error) {
	builder := app.builder.Create()
	if parsed.IsApplication() {
		parsedApplication := parsed.Application()
		app, err := app.applicationAdapter.ToApplication(parsedApplication)
		if err != nil {
			return nil, err
		}

		builder.WithApplication(app)
	}

	if parsed.IsDefinition() {
		parsedDef := parsed.Definition()
		def, err := app.definitionAdapter.ToDefinition(parsedDef)
		if err != nil {
			return nil, err
		}

		builder.WithDefinition(def)
	}

	return builder.Now()
}
