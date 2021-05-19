package machines

import (
	"errors"
	"fmt"

	"github.com/deepvalue-network/software/pangolin/domain/interpreters/stackframes"
	application_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/condition"
	var_value "github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/variable/value"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/variable/value/computable"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/labels"
	label_instructions "github.com/deepvalue-network/software/pangolin/domain/middle/applications/labels/label/instructions"
	label_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/applications/labels/label/instructions/instruction"
)

type instruction struct {
	computableValueBuilder computable.Builder
	labels                 labels.Labels
	stackFrame             stackframes.StackFrame
}

func createInstruction(
	computableValueBuilder computable.Builder,
	labels labels.Labels,
	stackFrame stackframes.StackFrame,
) Instruction {
	out := instruction{
		computableValueBuilder: computableValueBuilder,
		labels:                 labels,
		stackFrame:             stackFrame,
	}

	return &out
}

// Receive receives an instruction
func (app *instruction) Receive(ins application_instruction.Instruction) error {
	if ins.IsStackframe() {
		stkFrame := ins.Stackframe()
		if stkFrame.IsPush() {
			app.stackFrame.Push()
			return nil
		}

		if stkFrame.IsPop() {
			return app.stackFrame.Pop()
		}

		if stkFrame.IsIndex() {
			indexVariable := stkFrame.Index()
			stkFrameIndex := app.stackFrame.Index()
			value, err := app.computableValueBuilder.Create().WithInt64(int64(stkFrameIndex)).Now()
			if err != nil {
				return err
			}

			return app.stackFrame.Current().UpdateValue(indexVariable, value)
		}

		if stkFrame.IsSkip() {
			skip := stkFrame.Skip()
			if skip.IsVariable() {
				variable := skip.Variable()
				indexVariable, err := app.stackFrame.Current().Fetch(variable)
				if err != nil {
					return err
				}

				if !indexVariable.IsIntSixtyFour() {
					str := fmt.Sprintf("the skip variable (%s) was expected to contain an int64 value", variable)
					return errors.New(str)
				}

				ptrIndexVal := indexVariable.IntSixtyFour()
				return app.stackFrame.Skip(int(*ptrIndexVal))
			}
		}
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
		return app.stackFrame.Current().Standard(first, second, result, operation)
	}

	if ins.IsRemaining() {
		rem := ins.Remaining()
		first := rem.First()
		second := rem.Second()
		result := rem.Result()
		remaining := rem.Remaining()
		operation := rem.Operation()
		return app.stackFrame.Current().Remaining(first, second, result, remaining, operation)
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
		return app.stackFrame.Current().Insert(vr)
	}

	if ins.IsSave() {
		vr := ins.Save()
		return app.stackFrame.Current().Update(vr)
	}

	if ins.IsDelete() {
		name := ins.Delete()
		return app.stackFrame.Current().Delete(name)
	}

	if ins.IsCall() {
		panic(errors.New("finish call in instruction (interpreter)"))
	}

	if ins.IsExit() {
		exit := ins.Exit()
		if exit.HasCondition() {
			condition := exit.Condition()
			val, err := app.stackFrame.Current().Fetch(condition)
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
				app.stackFrame.Current().Stop()
			}

			return nil
		}

		app.stackFrame.Current().Stop()
		return nil
	}

	return errors.New("the instruction is invalid")
}

// ReceiveLbl receives a label instruction
func (app *instruction) ReceiveLbl(lblIns label_instruction.Instruction) (bool, error) {
	if lblIns.IsRet() {
		return true, nil
	}

	ins := lblIns.Instruction()
	err := app.Receive(ins)
	return false, err
}

func (app *instruction) labelInstructions(lblIns label_instructions.Instructions) error {
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

func (app *instruction) proposition(prop condition.Proposition) error {
	if prop.HasCondition() {
		cond := prop.Condition()
		com, err := app.stackFrame.Current().Fetch(cond)
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
	lbl, err := app.labels.Fetch(name)
	if err != nil {
		return err
	}

	lblInsList := lbl.Instructions().All()
	for _, oneLblIns := range lblInsList {
		stop, err := app.ReceiveLbl(oneLblIns)
		if err != nil {
			return err
		}

		if stop {
			break
		}
	}

	return nil
}

func (app *instruction) print(val var_value.Value) error {
	if val.IsComputable() {
		com := val.Computable()
		str := com.StringRepresentation()
		fmt.Println(str)
		return nil
	}

	name := val.Variable()
	com, err := app.stackFrame.Current().Fetch(name)
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
