package applications

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/heads"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/labels"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/tests"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

type adapter struct {
	headAdapter         heads.Adapter
	labelsAdapter       labels.Adapter
	instructionsAdapter instructions.Adapter
	testsAdapter        tests.Adapter
	builder             Builder
}

func createAdapter(
	headAdapter heads.Adapter,
	labelsAdapter labels.Adapter,
	instructionsAdapter instructions.Adapter,
	testsAdapter tests.Adapter,
	builder Builder,
) Adapter {
	out := adapter{
		headAdapter:         headAdapter,
		labelsAdapter:       labelsAdapter,
		instructionsAdapter: instructionsAdapter,
		testsAdapter:        testsAdapter,
		builder:             builder,
	}

	return &out
}

// ToApplication converts a parsed language application to application instance
func (app *adapter) ToApplication(parsed parsers.LanguageApplication) (Application, error) {
	parsedHead := parsed.Head()
	head, err := app.headAdapter.ToHead(parsedHead)
	if err != nil {
		return nil, err
	}

	parsedLabels := parsed.Labels()
	labels, err := app.labelsAdapter.ToLabels(parsedLabels)
	if err != nil {
		return nil, err
	}

	parsedMain := parsed.Main()
	main, err := app.instructionsAdapter.ToInstructions(parsedMain)
	if err != nil {
		return nil, err
	}

	builder := app.builder.Create().WithHead(head).WithLabels(labels).WithMain(main)
	if parsed.HasTests() {
		parsedTests := parsed.Tests()
		tests, err := app.testsAdapter.ToTests(parsedTests)
		if err != nil {
			return nil, err
		}

		builder.WithTests(tests)
	}

	return builder.Now()
}
