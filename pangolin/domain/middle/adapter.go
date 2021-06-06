package middle

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

type adapter struct {
	testableAdapter testables.Adapter
	languageAdapter applications.Adapter
	builder         Builder
}

func createAdapter(
	testableAdapter testables.Adapter,
	languageAdapter applications.Adapter,
	builder Builder,
) Adapter {
	out := adapter{
		testableAdapter: testableAdapter,
		languageAdapter: languageAdapter,
		builder:         builder,
	}

	return &out
}

// ToProgram converts a parsed program to program instance
func (app *adapter) ToProgram(parsed parsers.Program) (Program, error) {
	builder := app.builder.Create()
	if parsed.IsTestable() {
		parsedTestable := parsed.Testable()
		testable, err := app.testableAdapter.ToTestable(parsedTestable)
		if err != nil {
			return nil, err
		}

		builder.WithTestable(testable)
	}

	if parsed.IsLanguage() {
		parsedLanguage := parsed.Language()
		language, err := app.languageAdapter.ToApplication(parsedLanguage)
		if err != nil {
			return nil, err
		}

		builder.WithLanguage(language)
	}

	return builder.Now()
}
