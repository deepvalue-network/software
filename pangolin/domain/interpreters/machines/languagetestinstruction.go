package machines

import (
	language_test_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/tests/test/instructions/instruction"
)

type languageTestInstruction struct {
	insLangCommonApp LanguageInstructionCommon
	testInsApp       TestInstruction
}

func createLanguageTestInstruction(
	insLangCommonApp LanguageInstructionCommon,
	testInsApp TestInstruction,
) LanguageTestInstruction {
	out := languageTestInstruction{
		insLangCommonApp: insLangCommonApp,
		testInsApp:       testInsApp,
	}

	return &out
}

// Receive receives an instruction
func (app *languageTestInstruction) Receive(testIns language_test_instruction.Instruction) (bool, error) {
	if testIns.IsLanguage() {
		langIns := testIns.Language()
		return false, app.insLangCommonApp.Receive(
			langIns,
		)
	}

	testInstruction := testIns.Test()
	return app.testInsApp.Receive(testInstruction)
}
