package instruction

import (
	test_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/applications/tests/test/instructions/instruction"
	standard_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions/instruction"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

type adapter struct {
	languageAdapter standard_instruction.Adapter
	testAdapter     test_instruction.Adapter
	builder         Builder
}

func createAdapter(
	languageAdapter standard_instruction.Adapter,
	testAdapter test_instruction.Adapter,
	builder Builder,
) Adapter {
	out := adapter{
		languageAdapter: languageAdapter,
		testAdapter:     testAdapter,
		builder:         builder,
	}

	return &out
}

// ToInstruction converts a parsed language test instruction to instruction
func (app *adapter) ToInstruction(parsed parsers.LanguageTestInstruction) (Instruction, error) {
	builder := app.builder.Create()
	if parsed.IsLanguageInstruction() {
		parsedLangIns := parsed.LanguageInstruction()
		langIns, err := app.languageAdapter.ToCommonInstruction(parsedLangIns)
		if err != nil {
			return nil, err
		}

		builder.WithLanguage(langIns)
	}

	if parsed.IsTestInstruction() {
		parsedTest := parsed.TestInstruction()
		test, err := app.testAdapter.ToInstruction(parsedTest)
		if err != nil {
			return nil, err
		}

		builder.WithTest(test)
	}

	return builder.Now()
}
