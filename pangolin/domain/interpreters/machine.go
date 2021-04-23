package interpreters

import (
	"errors"
	"fmt"

	"github.com/deepvalue-network/software/pangolin/domain/lexers"
	lexer_parser "github.com/deepvalue-network/software/pangolin/domain/lexers/parser"
	"github.com/deepvalue-network/software/pangolin/domain/middle"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/condition"
	label_instructions "github.com/deepvalue-network/software/pangolin/domain/middle/applications/labels/label/instructions"
	label_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/applications/labels/label/instructions/instruction"
	var_variable "github.com/deepvalue-network/software/pangolin/domain/middle/variables/variable"
	var_value "github.com/deepvalue-network/software/pangolin/domain/middle/variables/variable/value"
	"github.com/deepvalue-network/software/pangolin/domain/middle/variables/variable/value/computable"
)

type machine struct {
	variableBuilder        var_variable.Builder
	valueBuilder           var_value.Builder
	computableValueBuilder computable.Builder
	lexerParserApplication lexer_parser.Application
	lexerParserBuilder     lexer_parser.Builder
	stkFrame               StackFrame
	lbls                   map[string]label_instructions.Instructions
	patternMatches         map[string]middle.PatternMatch
	lexerAdapterBuilder    lexers.AdapterBuilder
	currentToken           lexers.NodeTree
}

func createMachine(
	variableBuilder var_variable.Builder,
	valueBuilder var_value.Builder,
	computableValueBuilder computable.Builder,
	lexerParserApplication lexer_parser.Application,
	lexerParserBuilder lexer_parser.Builder,
	stkFrame StackFrame,
	lbls map[string]label_instructions.Instructions,
) Machine {
	return createMachineInternally(variableBuilder, valueBuilder, computableValueBuilder, lexerParserApplication, lexerParserBuilder, stkFrame, lbls, nil, nil)
}

func createMachineWithPatternMatches(
	variableBuilder var_variable.Builder,
	valueBuilder var_value.Builder,
	computableValueBuilder computable.Builder,
	lexerParserApplication lexer_parser.Application,
	lexerParserBuilder lexer_parser.Builder,
	stkFrame StackFrame,
	lbls map[string]label_instructions.Instructions,
	patternMatches map[string]middle.PatternMatch,
	lexerAdapterBuilder lexers.AdapterBuilder,
) Machine {
	return createMachineInternally(variableBuilder, valueBuilder, computableValueBuilder, lexerParserApplication, lexerParserBuilder, stkFrame, lbls, patternMatches, lexerAdapterBuilder)
}

func createMachineInternally(
	variableBuilder var_variable.Builder,
	valueBuilder var_value.Builder,
	computableValueBuilder computable.Builder,
	lexerParserApplication lexer_parser.Application,
	lexerParserBuilder lexer_parser.Builder,
	stkFrame StackFrame,
	lbls map[string]label_instructions.Instructions,
	patternMatches map[string]middle.PatternMatch,
	lexerAdapterBuilder lexers.AdapterBuilder,
) Machine {
	out := machine{
		variableBuilder:        variableBuilder,
		valueBuilder:           valueBuilder,
		computableValueBuilder: computableValueBuilder,
		lexerParserApplication: lexerParserApplication,
		lexerParserBuilder:     lexerParserBuilder,
		stkFrame:               stkFrame,
		lbls:                   lbls,
		patternMatches:         patternMatches,
		lexerAdapterBuilder:    lexerAdapterBuilder,
		currentToken:           nil,
	}

	return &out
}

// StackFrame returns the machine's StackFrame
func (app *machine) StackFrame() StackFrame {
	return app.stkFrame
}

// Receive receives an instruction
func (app *machine) Receive(ins instruction.Instruction) error {
	if ins.IsStackframe() {
		stkFrame := ins.Stackframe()
		if stkFrame.IsPush() {
			app.stkFrame.Push()
			return nil
		}

		if stkFrame.IsPop() {
			return app.stkFrame.Pop()
		}

		if stkFrame.IsIndex() {
			indexVariable := stkFrame.Index()
			stkFrameIndex := app.stkFrame.Index()
			value, err := app.computableValueBuilder.Create().WithInt64(int64(stkFrameIndex)).Now()
			if err != nil {
				return err
			}

			return app.stkFrame.Current().UpdateValue(indexVariable, value)
		}

		if stkFrame.IsSkip() {
			skip := stkFrame.Skip()
			if skip.IsVariable() {
				variable := skip.Variable()
				indexVariable, err := app.stkFrame.Current().Fetch(variable)
				if err != nil {
					return err
				}

				if !indexVariable.IsIntSixtyFour() {
					str := fmt.Sprintf("the skip variable (%s) was expected to contain an int64 value", variable)
					return errors.New(str)
				}

				ptrIndexVal := indexVariable.IntSixtyFour()
				return app.stkFrame.Skip(int(*ptrIndexVal))
			}
		}
	}

	if ins.IsTransform() {
		transform := ins.Transform()
		input := transform.Input()
		result := transform.Result()
		operation := transform.Operation()
		return app.stkFrame.Current().Transform(input, result, operation)
	}

	if ins.IsVariableName() {
		vrName := ins.VariableName()
		name := vrName.Variable()
		operation := vrName.Operation()
		if operation.IsMisc() {
			misc := operation.Misc()
			if misc.IsPush() {
				current := app.stkFrame.Current()
				app.stkFrame.Push()
				return app.stkFrame.Current().PushTo(name, current)
			}

			return errors.New("the misc's variableName is invalid")
		}
		return errors.New("the variableName is invalid")
	}

	if ins.IsCondition() {
		condition := ins.Condition()
		proposition := condition.Proposition()
		operation := condition.Operation()
		if operation.IsJump() {
			return app.proposition(proposition)
		}

		return errors.New("the condition is invalid")
	}

	if ins.IsStandard() {
		standard := ins.Standard()
		operation := standard.Operation()
		result := standard.Result()
		first := standard.First()
		second := standard.Second()
		return app.stkFrame.Current().Standard(first, second, result, operation)
	}

	if ins.IsRemaining() {
		rem := ins.Remaining()
		first := rem.First()
		second := rem.Second()
		result := rem.Result()
		remaining := rem.Remaining()
		operation := rem.Operation()
		return app.stkFrame.Current().Remaining(first, second, result, remaining, operation)
	}

	if ins.IsTransform() {
		trsf := ins.Transform()
		input := trsf.Input()
		result := trsf.Result()
		operation := trsf.Operation()
		return app.stkFrame.Current().Transform(input, result, operation)
	}

	if ins.IsValue() {
		val := ins.Value()
		varValue := val.Value()
		operation := val.Operation()
		if operation.IsPrint() {
			return app.print(varValue)
		}

		return errors.New("the value operation is invalid")
	}

	if ins.IsInsert() {
		vr := ins.Insert()
		return app.stkFrame.Current().Insert(vr)
	}

	if ins.IsSave() {
		vr := ins.Save()
		return app.stkFrame.Current().Update(vr)
	}

	if ins.IsDelete() {
		name := ins.Delete()
		return app.stkFrame.Current().Delete(name)
	}

	if ins.IsMatch() {
		if app.lexerAdapterBuilder == nil {
			return errors.New("the lexerAdapter builder is mandatory in order to execute a Match instruction in the machine")
		}

		match := ins.Match()
		inputName := match.Input()
		input, err := app.stkFrame.Current().Fetch(inputName)
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

	if ins.IsCall() {
		panic(errors.New("finish call in machine (interpreter)"))
	}

	if ins.IsExit() {
		exit := ins.Exit()
		if exit.HasCondition() {
			condition := exit.Condition()
			val, err := app.StackFrame().Current().Fetch(condition)
			if err != nil {
				return err
			}

			if val == nil {
				return nil
			}

			if !val.IsBool() {
				str := fmt.Sprintf("the condition inside the exit instruction was expected to be a boolean")
				return errors.New(str)
			}

			bl := val.Bool()
			if *bl {
				app.StackFrame().Current().Stop()
			}

			return nil
		}

		app.StackFrame().Current().Stop()
		return nil
	}

	if ins.IsToken() {
		token := ins.Token()
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

			err = app.StackFrame().Current().UpdateValue(sectionName, computableSection)
			if err != nil {
				return err
			}

			// code:
			computableCode, err := app.computableValueBuilder.Create().WithString(code).Now()
			if err != nil {
				return err
			}

			err = app.StackFrame().Current().UpdateValue(retName, computableCode)
			if err != nil {
				return err
			}

			return nil
		}

		if token.IsCode() {
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

				err = app.StackFrame().Current().UpdateValue(retName, computableCode)
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

				err = app.StackFrame().Current().UpdateValue(retName, computableCode)
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

					err = app.StackFrame().Current().Insert(variable)
					if err != nil {
						return err
					}

					// push:
					app.stkFrame.Push()
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

				err = app.StackFrame().Current().Insert(variable)
				if err != nil {
					return err
				}

				return nil
			}

			return errors.New("the token instruction is invalid")
		}

	}

	return errors.New("the instruction is invalid")
}

func (app *machine) treeLabelInstructions(lblIns label_instructions.Instructions, tree lexers.NodeTree) error {
	app.currentToken = tree
	return app.labelInstructions(lblIns)
}

func (app *machine) labelInstructions(lblIns label_instructions.Instructions) error {
	lblAll := lblIns.All()
	for _, oneLblIns := range lblAll {
		stop, err := app.labelInstruction(oneLblIns)
		if err != nil {
			return err
		}

		if stop {
			return nil
		}
	}

	return nil
}

func (app *machine) labelInstruction(lblIns label_instruction.Instruction) (bool, error) {
	if lblIns.IsRet() {
		return true, nil
	}

	ins := lblIns.Instruction()
	err := app.Receive(ins)
	return false, err
}

func (app *machine) proposition(prop condition.Proposition) error {
	if prop.HasCondition() {
		cond := prop.Condition()
		com, err := app.stkFrame.Current().Fetch(cond)
		if err != nil {
			return err
		}

		if com == nil {
			return nil
		}

		if !com.IsBool() {
			return errors.New("the condition expected a boolean value")
		}

		// skip:
		bl := com.Bool()
		if !*bl {
			return nil
		}
	}

	name := prop.Name()
	if ins, ok := app.lbls[name]; ok {
		err := app.labelInstructions(ins)
		if err != nil {
			return err
		}
	}

	str := fmt.Sprintf("the condition's proposition contains a name (%s) that is not a valid label", name)
	return errors.New(str)
}

func (app *machine) print(val var_value.Value) error {
	if val.IsComputable() {
		com := val.Computable()
		str := com.StringRepresentation()
		fmt.Println(str)
		return nil
	}

	if val.IsGlobalVariable() {
		name := val.GlobalVariable()
		com, err := app.stkFrame.Current().Fetch(name)
		if err != nil {
			return err
		}

		if com == nil {
			return nil
		}

		str := com.StringRepresentation()
		fmt.Println(str)
		return nil
	}

	name := val.LocalVariable()
	com, err := app.stkFrame.Current().Fetch(name)
	if err != nil {
		return err
	}

	if com == nil {
		return nil
	}

	str := com.StringRepresentation()
	fmt.Println(str)
	return nil
}
