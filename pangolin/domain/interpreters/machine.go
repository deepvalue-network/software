package interpreters

import (
	"errors"
	"fmt"

	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/condition"
	var_value "github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/variable/value"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/variable/value/computable"
	label_instructions "github.com/deepvalue-network/software/pangolin/domain/middle/applications/labels/label/instructions"
	label_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/applications/labels/label/instructions/instruction"
)

type machine struct {
	computableValueBuilder computable.Builder
	callLabelFn            CallLabelFunc
	fetchStackFrameFn      FetchStackFrameFunc
}

func createMachine(
	computableValueBuilder computable.Builder,
	callLabelFn CallLabelFunc,
	fetchStackFrameFn FetchStackFrameFunc,
) Machine {
	return createMachineInternally(computableValueBuilder, callLabelFn, fetchStackFrameFn)
}

func createMachineInternally(
	computableValueBuilder computable.Builder,
	callLabelFn CallLabelFunc,
	fetchStackFrameFn FetchStackFrameFunc,
) Machine {
	out := machine{
		computableValueBuilder: computableValueBuilder,
		callLabelFn:            callLabelFn,
		fetchStackFrameFn:      fetchStackFrameFn,
	}

	return &out
}

// Receive receives an instruction
func (app *machine) Receive(ins instruction.Instruction) error {
	if ins.IsStackframe() {
		stkFrame := ins.Stackframe()
		if stkFrame.IsPush() {
			app.fetchStackFrameFn().Push()
			return nil
		}

		if stkFrame.IsPop() {
			return app.fetchStackFrameFn().Pop()
		}

		if stkFrame.IsIndex() {
			indexVariable := stkFrame.Index()
			stkFrameIndex := app.fetchStackFrameFn().Index()
			value, err := app.computableValueBuilder.Create().WithInt64(int64(stkFrameIndex)).Now()
			if err != nil {
				return err
			}

			return app.fetchStackFrameFn().Current().UpdateValue(indexVariable, value)
		}

		if stkFrame.IsSkip() {
			skip := stkFrame.Skip()
			if skip.IsVariable() {
				variable := skip.Variable()
				indexVariable, err := app.fetchStackFrameFn().Current().Fetch(variable)
				if err != nil {
					return err
				}

				if !indexVariable.IsIntSixtyFour() {
					str := fmt.Sprintf("the skip variable (%s) was expected to contain an int64 value", variable)
					return errors.New(str)
				}

				ptrIndexVal := indexVariable.IntSixtyFour()
				return app.fetchStackFrameFn().Skip(int(*ptrIndexVal))
			}
		}
	}

	if ins.IsTransform() {
		transform := ins.Transform()
		input := transform.Input()
		result := transform.Result()
		operation := transform.Operation()
		return app.fetchStackFrameFn().Current().Transform(input, result, operation)
	}

	if ins.IsVariableName() {
		vrName := ins.VariableName()
		name := vrName.Variable()
		operation := vrName.Operation()
		if operation.IsMisc() {
			misc := operation.Misc()
			if misc.IsPush() {
				current := app.fetchStackFrameFn().Current()
				app.fetchStackFrameFn().Push()
				return app.fetchStackFrameFn().Current().PushTo(name, current)
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
		return app.fetchStackFrameFn().Current().Standard(first, second, result, operation)
	}

	if ins.IsRemaining() {
		rem := ins.Remaining()
		first := rem.First()
		second := rem.Second()
		result := rem.Result()
		remaining := rem.Remaining()
		operation := rem.Operation()
		return app.fetchStackFrameFn().Current().Remaining(first, second, result, remaining, operation)
	}

	if ins.IsTransform() {
		trsf := ins.Transform()
		input := trsf.Input()
		result := trsf.Result()
		operation := trsf.Operation()
		return app.fetchStackFrameFn().Current().Transform(input, result, operation)
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
		return app.fetchStackFrameFn().Current().Insert(vr)
	}

	if ins.IsSave() {
		vr := ins.Save()
		return app.fetchStackFrameFn().Current().Update(vr)
	}

	if ins.IsDelete() {
		name := ins.Delete()
		return app.fetchStackFrameFn().Current().Delete(name)
	}

	if ins.IsCall() {
		panic(errors.New("finish call in machine (interpreter)"))
	}

	if ins.IsExit() {
		exit := ins.Exit()
		if exit.HasCondition() {
			condition := exit.Condition()
			val, err := app.fetchStackFrameFn().Current().Fetch(condition)
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
				app.fetchStackFrameFn().Current().Stop()
			}

			return nil
		}

		app.fetchStackFrameFn().Current().Stop()
		return nil
	}

	return errors.New("the instruction is invalid")
}

// ReceiveLbl receives a label instruction
func (app *machine) ReceiveLbl(lblIns label_instruction.Instruction) (bool, error) {
	if lblIns.IsRet() {
		return true, nil
	}

	ins := lblIns.Instruction()
	err := app.Receive(ins)
	return false, err
}

func (app *machine) labelInstructions(lblIns label_instructions.Instructions) error {
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

func (app *machine) proposition(prop condition.Proposition) error {
	if prop.HasCondition() {
		cond := prop.Condition()
		com, err := app.fetchStackFrameFn().Current().Fetch(cond)
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
	err := app.callLabelFn(name)
	if err != nil {
		return err
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

	name := val.Variable()
	com, err := app.fetchStackFrameFn().Current().Fetch(name)
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
