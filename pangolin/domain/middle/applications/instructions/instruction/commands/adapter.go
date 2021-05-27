package commands

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/commands/heads"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/commands/labels"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/commands/languages"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/commands/mains"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/commands/scripts"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/commands/tests"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

type adapter struct {
	headAdapter     heads.Adapter
	labelAdapter    labels.Adapter
	languageAdapter languages.Adapter
	mainAdapter     mains.Adapter
	scriptAdapter   scripts.Adapter
	testAdapter     tests.Adapter
	builder         Builder
}

func createAdapter(
	headAdapter heads.Adapter,
	labelAdapter labels.Adapter,
	languageAdapter languages.Adapter,
	mainAdapter mains.Adapter,
	scriptAdapter scripts.Adapter,
	testAdapter tests.Adapter,
	builder Builder,
) Adapter {
	out := adapter{
		headAdapter:     headAdapter,
		labelAdapter:    labelAdapter,
		languageAdapter: languageAdapter,
		mainAdapter:     mainAdapter,
		scriptAdapter:   scriptAdapter,
		testAdapter:     testAdapter,
		builder:         builder,
	}

	return &out
}

// ToCommand converts a parsed command to a command instance
func (app *adapter) ToCommand(parsed parsers.Command) (Command, error) {
	builder := app.builder.Create()
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

	if parsed.IsHead() {
		parsedHead := parsed.Head()
		head, err := app.headAdapter.ToHead(parsedHead)
		if err != nil {
			return nil, err
		}

		builder.WithHead(head)
	}

	if parsed.IsMain() {
		parsedMain := parsed.Main()
		main, err := app.mainAdapter.ToMain(parsedMain)
		if err != nil {
			return nil, err
		}

		builder.WithMain(main)
	}

	if parsed.IsLabel() {
		parsedLabel := parsed.Label()
		label, err := app.labelAdapter.ToLabel(parsedLabel)
		if err != nil {
			return nil, err
		}

		builder.WithLabel(label)
	}

	if parsed.IsTest() {
		parsedTest := parsed.Test()
		test, err := app.testAdapter.ToTest(parsedTest)
		if err != nil {
			return nil, err
		}

		builder.WithTest(test)
	}

	return builder.Now()
}
