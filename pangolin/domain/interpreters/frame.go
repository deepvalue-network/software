package interpreters

import (
	"errors"
	"fmt"

	"github.com/deepvalue-network/software/pangolin/domain/middle/instructions/instruction/remaining"
	"github.com/deepvalue-network/software/pangolin/domain/middle/instructions/instruction/standard"
	"github.com/deepvalue-network/software/pangolin/domain/middle/instructions/instruction/transform"
	var_variable "github.com/deepvalue-network/software/pangolin/domain/middle/variables/variable"
	"github.com/deepvalue-network/software/pangolin/domain/middle/variables/variable/value/computable"
)

type frame struct {
	computer  Computer
	builder   computable.Builder
	variables map[string]computable.Value
	constants map[string]computable.Value
}

func createFrame(
	computer Computer,
	builder computable.Builder,
	variables map[string]computable.Value,
	constants map[string]computable.Value,
) Frame {
	out := frame{
		computer:  computer,
		builder:   builder,
		variables: variables,
		constants: constants,
	}

	return &out
}

// Standard executes a standard operation on the frame
func (app *frame) Standard(first string, second string, result string, operation standard.Operation) error {
	if _, ok := app.variables[first]; !ok {
		str := fmt.Sprintf("standard: the first variable (%s) is not defined", first)
		return errors.New(str)
	}

	if _, ok := app.variables[second]; !ok {
		str := fmt.Sprintf("standard: the second variable (%s) is not defined", second)
		return errors.New(str)
	}

	if _, ok := app.variables[result]; !ok {
		str := fmt.Sprintf("standard: the result variable (%s) is not defined", result)
		return errors.New(str)
	}

	if operation.IsArythmetic() {
		arythmetic := operation.Arythmetic()
		if arythmetic.IsAdd() {
			res, err := app.computer.Add(app.variables[first], app.variables[second])
			if err != nil {
				return err
			}

			app.variables[result] = res
		}

		if arythmetic.IsSub() {
			res, err := app.computer.Substract(app.variables[first], app.variables[second])
			if err != nil {
				return err
			}

			app.variables[result] = res
		}

		if arythmetic.IsMul() {
			res, err := app.computer.Multiply(app.variables[first], app.variables[second])
			if err != nil {
				return err
			}

			app.variables[result] = res
		}
	}

	if operation.IsRelational() {
		relational := operation.Relational()
		if relational.IsLessThan() {
			res, err := app.computer.IsLessThan(app.variables[first], app.variables[second])
			if err != nil {
				return err
			}

			app.variables[result] = res
		}

		if relational.IsEqual() {
			res, err := app.computer.IsEqual(app.variables[first], app.variables[second])
			if err != nil {
				return err
			}

			app.variables[result] = res
		}

		if relational.IsNotEqual() {
			res, err := app.computer.IsNotEqual(app.variables[first], app.variables[second])
			if err != nil {
				return err
			}

			app.variables[result] = res
		}
	}

	if operation.IsLogical() {
		logical := operation.Logical()
		if logical.IsAnd() {
			res, err := app.computer.And(app.variables[first], app.variables[second])
			if err != nil {
				return err
			}

			app.variables[result] = res
		}

		if logical.IsOr() {
			res, err := app.computer.Or(app.variables[first], app.variables[second])
			if err != nil {
				return err
			}

			app.variables[result] = res
		}
	}

	if operation.IsMisc() {
		misc := operation.Misc()
		if misc.IsConcatenation() {
			res, err := app.computer.Concat(app.variables[first], app.variables[second])
			if err != nil {
				return err
			}

			app.variables[result] = res
		}

		if misc.IsFrameAssignment() {
			fmt.Printf("finish frame assignment in frame")
		}
	}

	return nil
}

// Remaining executes a remaining operation on the frame
func (app *frame) Remaining(first string, second string, result string, remaining string, operation remaining.Operation) error {
	if _, ok := app.variables[first]; !ok {
		str := fmt.Sprintf("remaining: the first variable (%s) is not defined", first)
		return errors.New(str)
	}

	if _, ok := app.variables[second]; !ok {
		str := fmt.Sprintf("remaining: the second variable (%s) is not defined", second)
		return errors.New(str)
	}

	if _, ok := app.variables[result]; !ok {
		str := fmt.Sprintf("remaining: the result variable (%s) is not defined", result)
		return errors.New(str)
	}

	if _, ok := app.variables[remaining]; !ok {
		str := fmt.Sprintf("remaining: the result variable (%s) is not defined", remaining)
		return errors.New(str)
	}

	if operation.IsArythmetic() {
		arythmetic := operation.Arythmetic()
		if arythmetic.IsDiv() {
			res, rem, err := app.computer.Divide(app.variables[first], app.variables[second])
			if err != nil {
				return err
			}

			app.variables[result] = res
			app.variables[remaining] = rem
		}
	}

	return nil
}

// Transform executes a transform operation on the frame
func (app *frame) Transform(input string, result string, operation transform.Operation) error {
	if _, ok := app.variables[input]; !ok {
		str := fmt.Sprintf("transform: the input variable (%s) is not defined", input)
		return errors.New(str)
	}

	if _, ok := app.variables[result]; !ok {
		str := fmt.Sprintf("transform: the result variable (%s) is not defined", result)
		return errors.New(str)
	}

	if operation.IsMisc() {
		misc := operation.Misc()
		if misc.IsPop() {

		}
	}

	return nil
}

// PushTo pushes a frame to a variable's stack
func (app *frame) PushTo(name string, frame Frame) error {
	return nil
}

// Insert inserts a new variable on the frame
func (app *frame) Insert(vr var_variable.Variable) error {
	name := vr.Name()
	if _, ok := app.variables[name]; ok {
		str := fmt.Sprintf("variable: the name variable (%s) is already defined", name)
		return errors.New(str)
	}

	return app.save(vr)
}

// Update updates an existing variable on the frame
func (app *frame) Update(vr var_variable.Variable) error {
	name := vr.Name()
	if _, ok := app.variables[name]; !ok {
		str := fmt.Sprintf("variable: the name variable (%s) is not defined", name)
		return errors.New(str)
	}

	return app.save(vr)
}

func (app *frame) save(vr var_variable.Variable) error {
	name := vr.Name()
	if _, ok := app.variables[name]; !ok {
		app.variables[name] = nil
	}

	val := vr.Value()
	if val.IsComputable() {
		app.variables[name] = val.Computable()
		return nil
	}

	fmt.Printf("\n finish global and local variable updates in frame's interpreter")
	return nil
}

// UpdateValue updates a value by name
func (app *frame) UpdateValue(name string, val computable.Value) error {
	if _, ok := app.variables[name]; !ok {
		str := fmt.Sprintf("variable: the name variable (%s) is not defined", name)
		return errors.New(str)
	}

	app.variables[name] = val
	return nil
}

// Delete deletes a variable from the frame
func (app *frame) Delete(name string) error {
	if _, ok := app.variables[name]; !ok {
		str := fmt.Sprintf("variable: the name variable (%s) is not defined", name)
		return errors.New(str)
	}

	delete(app.variables, name)
	return nil
}

// Fetch fetches a variable by name
func (app *frame) Fetch(name string) (computable.Value, error) {
	if val, ok := app.variables[name]; ok {
		return val, nil
	}

	str := fmt.Sprintf("the variable (name: %s) is not defined", name)
	return nil, errors.New(str)
}
