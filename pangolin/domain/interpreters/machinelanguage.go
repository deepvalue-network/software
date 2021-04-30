package interpreters

import (
	"errors"
	"fmt"

	"github.com/deepvalue-network/software/pangolin/domain/lexers"
	lexer_parser "github.com/deepvalue-network/software/pangolin/domain/lexers/parser"
	var_variable "github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/variable"
	var_value "github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/variable/value"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/variable/value/computable"
	language_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions/instruction"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions/instruction/commands"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions/instruction/match"
	label_instructions "github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/labels/label/instructions"
	label_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/labels/label/instructions/instruction"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/labels/label/instructions/instruction/token"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/definitions"
)

type machineLanguage struct {
	variableBuilder        var_variable.Builder
	valueBuilder           var_value.Builder
	computableValueBuilder computable.Builder
	lexerParserApplication lexer_parser.Application
	lexerParserBuilder     lexer_parser.Builder
	lexerAdapterBuilder    lexers.AdapterBuilder
	patternMatches         map[string]definitions.PatternMatch
	lbls                   map[string]label_instructions.Instructions
	machine                Machine
	fetchStackFrameFn      FetchStackFrameFunc
	currentToken           lexers.NodeTree
}

func createMachineLanguage(
	variableBuilder var_variable.Builder,
	valueBuilder var_value.Builder,
	computableValueBuilder computable.Builder,
	lexerParserApplication lexer_parser.Application,
	lexerParserBuilder lexer_parser.Builder,
	lexerAdapterBuilder lexers.AdapterBuilder,
	patternMatches map[string]definitions.PatternMatch,
	lbls map[string]label_instructions.Instructions,
	machine Machine,
	fetchStackFrameFn FetchStackFrameFunc,
) MachineLanguage {
	out := machineLanguage{
		variableBuilder:        variableBuilder,
		valueBuilder:           valueBuilder,
		computableValueBuilder: computableValueBuilder,
		lexerParserApplication: lexerParserApplication,
		lexerParserBuilder:     lexerParserBuilder,
		lexerAdapterBuilder:    lexerAdapterBuilder,
		patternMatches:         patternMatches,
		lbls:                   lbls,
		machine:                machine,
		fetchStackFrameFn:      fetchStackFrameFn,
		currentToken:           nil,
	}

	return &out
}

// Receive receives a language instruction
func (app *machineLanguage) Receive(langIns language_instruction.Instruction) error {
	if langIns.IsInstruction() {
		ins := langIns.Instruction()
		err := app.machine.Receive(ins)
		if err != nil {
			return err
		}
	}

	if langIns.IsCommand() {
		command := langIns.Command()
		err := app.Command(command)
		if err != nil {
			return err
		}

	}

	if langIns.IsMatch() {
		match := langIns.Match()
		err := app.Match(match)
		if err != nil {
			return err
		}
	}

	return nil
}

// ReceiveLbl receives a label instruction
func (app *machineLanguage) ReceiveLbl(lblIns label_instruction.Instruction) (bool, error) {
	if lblIns.IsLabel() {
		subLblIns := lblIns.Label()
		return app.machine.ReceiveLbl(subLblIns)
	}

	if lblIns.IsLanguage() {
		langIns := lblIns.Language()
		err := app.Receive(langIns)
		return false, err
	}

	tokIns := lblIns.Token()
	err := app.token(tokIns)
	return false, err
}

// Command executes a command to the machine language
func (app *machineLanguage) Command(command commands.Command) error {
	panic(errors.New("finish command in machine language"))
}

// Match executes a match to the machine language
func (app *machineLanguage) Match(match match.Match) error {
	if app.lexerAdapterBuilder == nil {
		return errors.New("the lexerAdapter builder is mandatory in order to execute a Match instruction in the machine")
	}

	inputName := match.Input()
	input, err := app.fetchStackFrameFn().Current().Fetch(inputName)
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
			if _, ok := app.lbls[enter]; !ok {
				str := fmt.Sprintf("the label %s was set as an onEnter label on pattern: %s, but the label is not declared", enter, onePatternMatch.Pattern())
				return errors.New(str)
			}

			evt.OnEnter = func(tree lexers.NodeTree) (interface{}, error) {
				err := app.treeLabelInstructions(app.lbls[enter], tree)
				if err != nil {
					return nil, err
				}

				return nil, nil
			}
		}

		if onePatternMatch.HasExitLabel() {
			exit := onePatternMatch.ExitLabel()
			if _, ok := app.lbls[exit]; !ok {
				str := fmt.Sprintf("the label %s was set as an onExit label on pattern: %s, but the label is not declared", exit, onePatternMatch.Pattern())
				return errors.New(str)
			}

			evt.OnExit = func(tree lexers.NodeTree) (interface{}, error) {
				err := app.treeLabelInstructions(app.lbls[exit], tree)
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

func (app *machineLanguage) treeLabelInstructions(lblIns label_instructions.Instructions, tree lexers.NodeTree) error {
	app.currentToken = tree
	return app.labelInstructions(lblIns)
}

func (app *machineLanguage) labelInstructions(lblIns label_instructions.Instructions) error {
	lblAll := lblIns.All()
	for _, oneLblIns := range lblAll {
		stop, err := app.ReceiveLbl(oneLblIns)
		if err != nil {
			return err
		}

		if stop {
			return nil
		}
	}

	return nil
}

func (app *machineLanguage) token(token token.Token) error {
	if token.IsCodeMatch() {
		codeMatch := token.CodeMatch()
		retName := codeMatch.Return()
		sectionName := codeMatch.SectionName()
		patternNames := codeMatch.Patterns()
		section, code := app.currentToken.BestMatchFromNames(patternNames)

		// section:
		computableSection, err := app.computableValueBuilder.Create().WithString(section).Now()
		if err != nil {
			return err
		}

		err = app.fetchStackFrameFn().Current().UpdateValue(sectionName, computableSection)
		if err != nil {
			return err
		}

		// code:
		computableCode, err := app.computableValueBuilder.Create().WithString(code).Now()
		if err != nil {
			return err
		}

		err = app.fetchStackFrameFn().Current().UpdateValue(retName, computableCode)
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
		code := app.currentToken.Code()

		// code:
		computableCode, err := app.computableValueBuilder.Create().WithString(code).Now()
		if err != nil {
			return err
		}

		err = app.fetchStackFrameFn().Current().UpdateValue(retName, computableCode)
		if err != nil {
			return err
		}

		return nil
	}

	if hasPattern && !hasAmount {
		// fetch code from token's name:
		pattern := code.Pattern()
		code := app.currentToken.CodeFromName(pattern)

		// code:
		computableCode, err := app.computableValueBuilder.Create().WithString(code).Now()
		if err != nil {
			return err
		}

		err = app.fetchStackFrameFn().Current().UpdateValue(retName, computableCode)
		if err != nil {
			return err
		}

		return nil
	}

	if hasPattern && hasAmount {
		// fetch code from token's name:
		pattern := code.Pattern()
		codes := app.currentToken.CodesFromName(pattern)
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

			err = app.fetchStackFrameFn().Current().Insert(variable)
			if err != nil {
				return err
			}

			// push:
			app.fetchStackFrameFn().Push()
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

		err = app.fetchStackFrameFn().Current().Insert(variable)
		if err != nil {
			return err
		}

		return nil
	}

	return errors.New("the token instruction is invalid")
}
