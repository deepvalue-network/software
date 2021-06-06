package machines

import (
	"errors"
	"fmt"

	"github.com/deepvalue-network/software/pangolin/domain/interpreters/stackframes"
	application_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction/condition"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction/registry"
	var_value "github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction/variable/value"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction/variable/value/computable"
	label_instructions "github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/labels/label/instructions"
	label_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/labels/label/instructions/instruction"
)

type instruction struct {
	computableValueBuilder computable.Builder
	labelFn                CallLabelByNameFn
	currentStackFrameName  string
	stackframes            map[string]stackframes.StackFrame
}

func createInstruction(
	computableValueBuilder computable.Builder,
	labelFn CallLabelByNameFn,
	stackFrame stackframes.StackFrame,
) Instruction {
	out := instruction{
		computableValueBuilder: computableValueBuilder,
		labelFn:                labelFn,
		currentStackFrameName:  currentStackFrameName,
		stackframes: map[string]stackframes.StackFrame{
			currentStackFrameName: stackFrame,
		},
	}

	return &out
}

// Receive receives an instruction
func (app *instruction) Receive(ins application_instruction.Instruction) error {
	if ins.IsStackframe() {
		stkFrame := ins.Stackframe()
		if stkFrame.IsPush() {
			app.stackframes[app.currentStackFrameName].Push()
			return nil
		}

		if stkFrame.IsPop() {
			return app.stackframes[app.currentStackFrameName].Pop()
		}

		if stkFrame.IsIndex() {
			indexVariable := stkFrame.Index()
			stkFrameIndex := app.stackframes[app.currentStackFrameName].Index()
			value, err := app.computableValueBuilder.Create().WithInt64(int64(stkFrameIndex)).Now()
			if err != nil {
				return err
			}

			return app.stackframes[app.currentStackFrameName].Current().UpdateValue(indexVariable, value)
		}

		if stkFrame.IsSkip() {
			skip := stkFrame.Skip()
			if skip.IsVariable() {
				variable := skip.Variable()
				indexVariable, err := app.stackframes[app.currentStackFrameName].Current().Fetch(variable)
				if err != nil {
					return err
				}

				if !indexVariable.IsIntSixtyFour() {
					str := fmt.Sprintf("the skip variable (%s) was expected to contain an int64 value", variable)
					return errors.New(str)
				}

				ptrIndexVal := indexVariable.IntSixtyFour()
				return app.stackframes[app.currentStackFrameName].Skip(int(*ptrIndexVal))
			}
		}

		if stkFrame.IsSave() {
			save := stkFrame.Save()
			to := save.To()
			from := app.currentStackFrameName
			if save.HasFrom() {
				from = save.From()
			}

			if _, ok := app.stackframes[to]; !ok {
				str := fmt.Sprintf("the to stackFrame (%s) does not exists", to)
				return errors.New(str)
			}

			if fromStackFrame, ok := app.stackframes[from]; ok {
				app.stackframes[to] = fromStackFrame
			}

			str := fmt.Sprintf("the from stackFrame (%s) does not exists", from)
			return errors.New(str)
		}

		if stkFrame.IsSwitch() {
			variableName := stkFrame.Switch()
			if _, ok := app.stackframes[variableName]; !ok {
				str := fmt.Sprintf("the stackFrame (%s) does not exists", variableName)
				return errors.New(str)
			}

			app.currentStackFrameName = variableName
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
		return app.stackframes[app.currentStackFrameName].Current().Standard(first, second, result, operation)
	}

	if ins.IsRemaining() {
		rem := ins.Remaining()
		first := rem.First()
		second := rem.Second()
		result := rem.Result()
		remaining := rem.Remaining()
		operation := rem.Operation()
		return app.stackframes[app.currentStackFrameName].Current().Remaining(first, second, result, remaining, operation)
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
		return app.stackframes[app.currentStackFrameName].Current().Insert(vr)
	}

	if ins.IsSave() {
		vr := ins.Save()
		return app.stackframes[app.currentStackFrameName].Current().Update(vr)
	}

	if ins.IsDelete() {
		name := ins.Delete()
		if _, ok := app.stackframes[name]; ok {
			if app.currentStackFrameName == name {
				str := fmt.Sprintf("the stackFrame (%s) cannot be deleted because it is the current stackFrame", name)
				return errors.New(str)
			}

			delete(app.stackframes, name)
		}

		return app.stackframes[app.currentStackFrameName].Current().Delete(name)
	}

	if ins.IsCall() {
		panic(errors.New("finish call in instruction (interpreter)"))
	}

	if ins.IsModule() {
		panic(errors.New("finish module in instruction (interpreter)"))
	}

	if ins.IsExit() {
		exit := ins.Exit()
		if exit.HasCondition() {
			condition := exit.Condition()
			val, err := app.stackframes[app.currentStackFrameName].Current().Fetch(condition)
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
				app.stackframes[app.currentStackFrameName].Current().Stop()
			}

			return nil
		}

		app.stackframes[app.currentStackFrameName].Current().Stop()
		return nil
	}

	if ins.IsRegistry() {
		registry := ins.Registry()
		if registry.IsFetch() {
			fetch := registry.Fetch()
			from := fetch.From()
			to := fetch.To()
			if fetch.HasIndex() {
				index := fetch.Index()
				indexInt, err := app.registryIndexToIndex(index)
				if err != nil {
					return err
				}

				currentIndex := app.stackframes[app.currentStackFrameName].Index()
				err = app.stackframes[app.currentStackFrameName].Skip(indexInt)
				if err != nil {
					return err
				}

				fromVal, err := app.stackframes[app.currentStackFrameName].Registry().Fetch(from)
				if err != nil {
					return err
				}

				err = app.stackframes[app.currentStackFrameName].Current().UpdateValue(to, fromVal)
				if err != nil {
					return err
				}

				err = app.stackframes[app.currentStackFrameName].Skip(currentIndex)
				if err != nil {
					return err
				}

				return nil
			}

			fromVal, err := app.stackframes[app.currentStackFrameName].Registry().Fetch(from)
			if err != nil {
				return err
			}

			err = app.stackframes[app.currentStackFrameName].Current().UpdateValue(to, fromVal)
			if err != nil {
				return err
			}

			return nil
		}

		if registry.IsRegister() {
			register := registry.Register()
			variable := register.Variable()
			if register.HasIndex() {
				index := register.Index()
				indexInt, err := app.registryIndexToIndex(index)
				if err != nil {
					return err
				}

				currentIndex := app.stackframes[app.currentStackFrameName].Index()
				err = app.stackframes[app.currentStackFrameName].Skip(indexInt)
				if err != nil {
					return err
				}

				fromVal, err := app.stackframes[app.currentStackFrameName].Current().Fetch(variable)
				if err != nil {
					return err
				}

				err = app.stackframes[app.currentStackFrameName].Registry().Insert(variable, fromVal)
				if err != nil {
					return err
				}

				err = app.stackframes[app.currentStackFrameName].Skip(currentIndex)
				if err != nil {
					return err
				}

				return nil
			}

			fromVal, err := app.stackframes[app.currentStackFrameName].Current().Fetch(variable)
			if err != nil {
				return err
			}

			err = app.stackframes[app.currentStackFrameName].Registry().Insert(variable, fromVal)
			if err != nil {
				return err
			}

			return nil
		}

		if registry.IsUnregister() {
			variable := registry.Unregister()
			err := app.stackframes[app.currentStackFrameName].Registry().Delete(variable)
			if err != nil {
				return err
			}

			return nil
		}
	}

	return errors.New("the instruction is invalid")
}

func (app *instruction) registryIndexToIndex(index registry.Index) (int, error) {
	if index.IsInt() {
		val := index.Int()
		return int(val), nil
	}

	variable := index.Variable()
	vr, err := app.stackframes[app.currentStackFrameName].Current().Fetch(variable)
	if err != nil {
		return 0, err
	}

	if vr.IsIntHeight() {
		ptr := vr.IntHeight()
		return int(*ptr), nil
	}

	if vr.IsIntSixteen() {
		ptr := vr.IntSixteen()
		return int(*ptr), nil
	}

	if vr.IsIntThirtyTwo() {
		ptr := vr.IntThirtyTwo()
		return int(*ptr), nil
	}

	if vr.IsIntSixtyFour() {
		ptr := vr.IntSixtyFour()
		return int(*ptr), nil
	}

	str := fmt.Sprintf("the registry index's variable (%s) was expected to contain an int", variable)
	return 0, errors.New(str)
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
		com, err := app.stackframes[app.currentStackFrameName].Current().Fetch(cond)
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
	return app.labelFn(name)
}

func (app *instruction) print(val var_value.Value) error {
	if val.IsComputable() {
		com := val.Computable()
		str := com.StringRepresentation()
		fmt.Println(str)
		return nil
	}

	name := val.Variable()
	com, err := app.stackframes[app.currentStackFrameName].Current().Fetch(name)
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
