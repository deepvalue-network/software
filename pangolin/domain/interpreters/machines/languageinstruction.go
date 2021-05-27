package machines

import (
	"errors"

	"github.com/deepvalue-network/software/pangolin/domain/interpreters/composers"
	"github.com/deepvalue-network/software/pangolin/domain/interpreters/stackframes"
	var_variable "github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction/variable"
	var_value "github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction/variable/value"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction/variable/value/computable"
	language_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction"
	language_label_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/applications/labels/label/instructions/instruction"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/labels/label/instructions/instruction/token"
)

type languageInstruction struct {
	variableBuilder        var_variable.Builder
	valueBuilder           var_value.Builder
	computableValueBuilder computable.Builder
	langCommonIns          LanguageInstructionCommon
	insApp                 Instruction
	stackFrame             stackframes.StackFrame
	languageState          LanguageState
	composerApp            composers.Composer
}

func createLanguageInstruction(
	variableBuilder var_variable.Builder,
	valueBuilder var_value.Builder,
	computableValueBuilder computable.Builder,
	langCommonIns LanguageInstructionCommon,
	insApp Instruction,
	stackFrame stackframes.StackFrame,
	languageState LanguageState,
	composerApp composers.Composer,
) LanguageInstruction {
	out := languageInstruction{
		variableBuilder:        variableBuilder,
		valueBuilder:           valueBuilder,
		computableValueBuilder: computableValueBuilder,
		langCommonIns:          langCommonIns,
		insApp:                 insApp,
		stackFrame:             stackFrame,
		languageState:          languageState,
		composerApp:            composerApp,
	}

	return &out
}

// Receive receives an instruction
func (app *languageInstruction) Receive(langIns language_instruction.Instruction) error {
	if langIns.IsInstruction() {
		ins := langIns.Instruction()
		return app.langCommonIns.Receive(ins)
	}

	if langIns.IsCommand() {
		command := langIns.Command()
		return app.composerApp.Receive(command)
	}

	return errors.New("the language instruction is invalid")
}

// ReceiveLbl receives a label instruction
func (app *languageInstruction) ReceiveLbl(lblIns language_label_instruction.Instruction) (bool, error) {
	if lblIns.IsLabel() {
		label := lblIns.Label()
		return app.insApp.ReceiveLbl(label)
	}

	if lblIns.IsLanguage() {
		language := lblIns.Language()
		err := app.Receive(language)
		return false, err
	}

	if lblIns.IsToken() {
		token := lblIns.Token()
		err := app.token(token)
		return false, err
	}

	return false, errors.New("the label instruction is invalid")
}

func (app *languageInstruction) token(token token.Token) error {
	if token.IsCodeMatch() {
		codeMatch := token.CodeMatch()
		retName := codeMatch.Return()
		sectionName := codeMatch.SectionName()
		patternNames := codeMatch.Patterns()
		section, code := app.languageState.CurrentToken().BestMatchFromNames(patternNames)

		// section:
		computableSection, err := app.computableValueBuilder.Create().WithString(section).Now()
		if err != nil {
			return err
		}

		err = app.stackFrame.Current().UpdateValue(sectionName, computableSection)
		if err != nil {
			return err
		}

		// code:
		computableCode, err := app.computableValueBuilder.Create().WithString(code).Now()
		if err != nil {
			return err
		}

		err = app.stackFrame.Current().UpdateValue(retName, computableCode)
		if err != nil {
			return err
		}

		return nil
	}

	code := token.Code()
	retName := code.Return()
	hasPattern := code.HasPattern()
	hasAmount := code.HasAmount()
	if !hasPattern && !hasAmount {
		// fetch code from token:
		code := app.languageState.CurrentToken().Code()

		// code:
		computableCode, err := app.computableValueBuilder.Create().WithString(code).Now()
		if err != nil {
			return err
		}

		err = app.stackFrame.Current().UpdateValue(retName, computableCode)
		if err != nil {
			return err
		}

		return nil
	}

	if hasPattern && !hasAmount {
		// fetch code from token's name:
		pattern := code.Pattern()
		code := app.languageState.CurrentToken().CodeFromName(pattern)

		// code:
		computableCode, err := app.computableValueBuilder.Create().WithString(code).Now()
		if err != nil {
			return err
		}

		err = app.stackFrame.Current().UpdateValue(retName, computableCode)
		if err != nil {
			return err
		}

		return nil
	}

	if hasPattern && hasAmount {
		// fetch code from token's name:
		pattern := code.Pattern()
		codes := app.languageState.CurrentToken().CodesFromName(pattern)
		amount := code.Amount()

		// add the codes:
		for _, oneCode := range codes {
			computableCode, err := app.computableValueBuilder.Create().WithString(oneCode).Now()
			if err != nil {
				return err
			}

			value, err := app.valueBuilder.WithComputable(computableCode).Now()
			if err != nil {
				return err
			}

			variable, err := app.variableBuilder.Create().WithName(retName).WithValue(value).Now()
			if err != nil {
				return err
			}

			err = app.stackFrame.Current().Insert(variable)
			if err != nil {
				return err
			}

			// push:
			app.stackFrame.Push()
		}

		// amount:
		computableAmount, err := app.computableValueBuilder.Create().WithInt64(int64(len(codes))).Now()
		if err != nil {
			return err
		}

		value, err := app.valueBuilder.WithComputable(computableAmount).Now()
		if err != nil {
			return err
		}

		variable, err := app.variableBuilder.Create().WithName(amount).WithValue(value).Now()
		if err != nil {
			return err
		}

		err = app.stackFrame.Current().Insert(variable)
		if err != nil {
			return err
		}

		return nil
	}

	return errors.New("the token instruction is invalid")
}
