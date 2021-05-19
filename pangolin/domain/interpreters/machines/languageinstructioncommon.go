package machines

import (
	"errors"

	"github.com/deepvalue-network/software/pangolin/domain/interpreters/stackframes"
	"github.com/deepvalue-network/software/pangolin/domain/lexers"
	lexer_parser "github.com/deepvalue-network/software/pangolin/domain/lexers/parser"
	var_variable "github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/variable"
	var_value "github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/variable/value"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/variable/value/computable"
	language_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions/instruction"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions/instruction/match"
	language_labels "github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/labels"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/definitions"
)

type languageInstructionCommon struct {
	variableBuilder        var_variable.Builder
	valueBuilder           var_value.Builder
	computableValueBuilder computable.Builder
	lexerParserApplication lexer_parser.Application
	lexerParserBuilder     lexer_parser.Builder
	lexerAdapterBuilder    lexers.AdapterBuilder
	patternMatches         map[string]definitions.PatternMatch
	insApp                 Instruction
	langInsApp             LanguageInstruction
	stackFrame             stackframes.StackFrame
	labels                 language_labels.Labels
	languageState          LanguageState
}

func createLanguageInstructionCommon(
	variableBuilder var_variable.Builder,
	valueBuilder var_value.Builder,
	computableValueBuilder computable.Builder,
	lexerParserApplication lexer_parser.Application,
	lexerParserBuilder lexer_parser.Builder,
	lexerAdapterBuilder lexers.AdapterBuilder,
	patternMatches map[string]definitions.PatternMatch,
	insApp Instruction,
	langInsApp LanguageInstruction,
	stackFrame stackframes.StackFrame,
	labels language_labels.Labels,
	languageState LanguageState,
) LanguageInstructionCommon {
	out := languageInstructionCommon{
		variableBuilder:        variableBuilder,
		valueBuilder:           valueBuilder,
		computableValueBuilder: computableValueBuilder,
		lexerParserApplication: lexerParserApplication,
		lexerParserBuilder:     lexerParserBuilder,
		lexerAdapterBuilder:    lexerAdapterBuilder,
		patternMatches:         patternMatches,
		insApp:                 insApp,
		langInsApp:             langInsApp,
		stackFrame:             stackFrame,
		labels:                 labels,
		languageState:          languageState,
	}

	return &out
}

// Receive receives an instruction
func (app *languageInstructionCommon) Receive(commonIns language_instruction.CommonInstruction) error {
	if commonIns.IsInstruction() {
		ins := commonIns.Instruction()
		return app.insApp.Receive(ins)
	}

	if commonIns.IsMatch() {
		match := commonIns.Match()
		return app.match(match)
	}

	return errors.New("the language instruction is invalid")
}

func (app *languageInstructionCommon) match(match match.Match) error {
	if app.lexerAdapterBuilder == nil {
		return errors.New("the lexerAdapter builder is mandatory in order to execute a Match instruction in the machine")
	}

	inputName := match.Input()
	input, err := app.stackFrame.Current().Fetch(inputName)
	if err != nil {
		return err
	}

	if input == nil {
		return nil
	}

	lexerAdapterBuilder := app.lexerAdapterBuilder
	if match.HasPattern() {
		root := match.Pattern()
		lexerAdapterBuilder.WithRoot(root)
	}

	lexerAdapter, err := lexerAdapterBuilder.Now()
	if err != nil {
		return err
	}

	if !input.IsString() {
		return errors.New("the input in the match was expecting a string")
	}

	script := input.StringRepresentation()
	lexer, err := lexerAdapter.ToLexer(script)
	if err != nil {
		return err
	}

	params := []lexer_parser.ToEventsParams{}
	for _, onePatternMatch := range app.patternMatches {
		evt := lexer_parser.ToEventsParams{
			Token: onePatternMatch.Pattern(),
		}

		if onePatternMatch.HasEnterLabel() {
			enter := onePatternMatch.EnterLabel()
			evt.OnEnter = func(tree lexers.NodeTree) (interface{}, error) {
				err := app.treeLabelInstructions(enter, tree)
				if err != nil {
					return nil, err
				}

				return nil, nil
			}
		}

		if onePatternMatch.HasExitLabel() {
			exit := onePatternMatch.ExitLabel()
			evt.OnExit = func(tree lexers.NodeTree) (interface{}, error) {
				err := app.treeLabelInstructions(exit, tree)
				if err != nil {
					return nil, err
				}

				return nil, nil
			}
		}

		params = append(params, evt)
	}

	lexerParser, err := app.lexerParserBuilder.Create().WithLexer(lexer).WithEventParams(params).Now()
	if err != nil {
		return err
	}

	_, err = app.lexerParserApplication.Execute(lexerParser)
	if err != nil {
		return err
	}

	return nil
}

func (app *languageInstructionCommon) treeLabelInstructions(labelName string, tree lexers.NodeTree) error {
	app.languageState.ChangeCurrentToken(tree)

	lbl, err := app.labels.Fetch(labelName)
	if err != nil {
		return err
	}

	insList := lbl.Instructions().All()
	for _, oneInstruction := range insList {
		stops, err := app.langInsApp.ReceiveLbl(oneInstruction)
		if err != nil {
			return err
		}

		if stops {
			break
		}
	}

	return nil
}
