package middle

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages"
	"github.com/deepvalue-network/software/pangolin/domain/middle/scripts"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

type adapter struct {
	applicationAdapter applications.Adapter
	languageAdapter    languages.Adapter
	scriptAdapter      scripts.Adapter
	builder            Builder
}

func createAdapter(
	applicationAdapter applications.Adapter,
	languageAdapter languages.Adapter,
	scriptAdapter scripts.Adapter,
	builder Builder,
) Adapter {
	out := adapter{
		applicationAdapter: applicationAdapter,
		languageAdapter:    languageAdapter,
		scriptAdapter:      scriptAdapter,
		builder:            builder,
	}

	return &out
}

// ToProgram converts a parsed program to program instance
func (app *adapter) ToProgram(parsed parsers.Program) (Program, error) {
	builder := app.builder.Create()
	if parsed.IsApplication() {
		parsedApplication := parsed.Application()
		app, err := app.applicationAdapter.ToApplication(parsedApplication)
		if err != nil {
			return nil, err
		}

		builder.WithApplication(app)
	}

	if parsed.IsLanguage() {
		parsedLang := parsed.Language()
		lang, err := app.languageAdapter.ToLanguage(parsedLang)
		if err != nil {
			return nil, err
		}

		builder.WithLanguage(lang)
	}

	if parsed.IsScript() {
		parsedScript := parsed.Script()
		script, err := app.scriptAdapter.ToScript(parsedScript)
		if err != nil {
			return nil, err
		}

		builder.WithScript(script)
	}

	return builder.Now()
}
