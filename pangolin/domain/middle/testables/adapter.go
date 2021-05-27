package testables

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/languages/definitions"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

type adapter struct {
	executableAdapter executables.Adapter
	languageAdapter   definitions.Adapter
	builder           Builder
}

func createAdapter(
	executableAdapter executables.Adapter,
	languageAdapter definitions.Adapter,
	builder Builder,
) Adapter {
	out := adapter{
		executableAdapter: executableAdapter,
		languageAdapter:   languageAdapter,
		builder:           builder,
	}

	return &out
}

// ToTestable converts a parsed Testable to a Testable instance
func (app *adapter) ToTestable(parsed parsers.Testable) (Testable, error) {
	builder := app.builder.Create()
	if parsed.IsExecutable() {
		parsedExecutable := parsed.Executable()
		executable, err := app.executableAdapter.ToExecutable(parsedExecutable)
		if err != nil {
			return nil, err
		}

		builder.WithExecutable(executable)
	}

	if parsed.IsLanguage() {
		parsedLanguage := parsed.Language()
		language, err := app.languageAdapter.ToDefinition(parsedLanguage)
		if err != nil {
			return nil, err
		}

		builder.WithLanguage(language)
	}

	return builder.Now()
}
