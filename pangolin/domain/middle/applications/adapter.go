package applications

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/labels"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/tests"
	"github.com/deepvalue-network/software/pangolin/domain/middle/heads"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

type adapter struct {
	testsAdapter        tests.Adapter
	labelsAdapter       labels.Adapter
	instructionsAdapter instructions.Adapter
	headAdapter         heads.Adapter
	builder             Builder
}

func createAdapter(
	testsAdapter tests.Adapter,
	labelsAdapter labels.Adapter,
	instructionsAdapter instructions.Adapter,
	headAdapter heads.Adapter,
	builder Builder,
) Adapter {
	out := adapter{
		testsAdapter:        testsAdapter,
		labelsAdapter:       labelsAdapter,
		instructionsAdapter: instructionsAdapter,
		headAdapter:         headAdapter,
		builder:             builder,
	}

	return &out
}

// ToApplication converts a parsed application to an application instance
func (app *adapter) ToApplication(parsed parsers.Application) (Application, error) {
	mainIns := parsed.Main().Instructions()
	instructions, err := app.instructionsAdapter.ToInstructions(mainIns)
	if err != nil {
		return nil, err
	}

	parsedHead := parsed.Head()
	head, err := app.headAdapter.ToHead(parsedHead)
	if err != nil {
		return nil, err
	}

	builder := app.builder.Create().WithMain(instructions).WithHead(head)
	if parsed.HasTest() {
		parsedTest := parsed.Test()
		tests, err := app.testsAdapter.ToTests(parsedTest)
		if err != nil {
			return nil, err
		}

		builder.WithTests(tests)
	}

	if parsed.HasLabel() {
		section := parsed.Label()
		labels, err := app.labelsAdapter.ToLabels(section)
		if err != nil {
			return nil, err
		}

		builder.WithLabels(labels)
	}

	return builder.Now()
}
