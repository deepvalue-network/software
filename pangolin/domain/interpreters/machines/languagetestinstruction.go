package machines

import (
	"github.com/deepvalue-network/software/pangolin/domain/interpreters/composers"
	"github.com/deepvalue-network/software/pangolin/domain/interpreters/stackframes"
	language_test_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/tests/test/instructions/instruction"
)

type languageTestInstruction struct {
	frameBuilder          stackframes.FrameBuilder
	insLangCommonApp      LanguageInstructionCommon
	testInsApp            TestInstruction
	composerApp           composers.Composer
	stackFrame            stackframes.StackFrame
	interpreterCallBackFn InterpretCallBackFn
}

func createLanguageTestInstruction(
	frameBuilder stackframes.FrameBuilder,
	insLangCommonApp LanguageInstructionCommon,
	testInsApp TestInstruction,
	composerApp composers.Composer,
	stackFrame stackframes.StackFrame,
	interpreterCallBackFn InterpretCallBackFn,
) LanguageTestInstruction {
	out := languageTestInstruction{
		frameBuilder:          frameBuilder,
		insLangCommonApp:      insLangCommonApp,
		testInsApp:            testInsApp,
		composerApp:           composerApp,
		stackFrame:            stackFrame,
		interpreterCallBackFn: interpreterCallBackFn,
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

	if testIns.IsInterpret() {
		linkedApp, err := app.composerApp.Now()
		if err != nil {
			return false, err
		}

		retStackFrame, err := app.interpreterCallBackFn(linkedApp, app.stackFrame)
		if err != nil {
			return false, nil
		}

		variables := retStackFrame.Registry().All()
		frame := app.frameBuilder.Create().WithVariables(variables).Now()
		app.stackFrame.Add(frame)
	}

	testInstruction := testIns.Test()
	return app.testInsApp.Receive(testInstruction)
}
