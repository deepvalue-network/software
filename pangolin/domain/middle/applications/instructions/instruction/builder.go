package instruction

import (
	"errors"

	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/call"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/condition"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/exit"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/remaining"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/stackframe"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/standard"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/transform"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/value"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/variablename"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/variable"
)

type builder struct {
	stackframe   stackframe.Stackframe
	transform    transform.Transform
	variableName variablename.VariableName
	condition    condition.Condition
	standard     standard.Standard
	remaining    remaining.Remaining
	value        value.Value
	insert       variable.Variable
	save         variable.Variable
	del          string
	call         call.Call
	exit         exit.Exit
}

func createBuilder() Builder {
	out := builder{
		stackframe:   nil,
		transform:    nil,
		variableName: nil,
		condition:    nil,
		standard:     nil,
		remaining:    nil,
		value:        nil,
		insert:       nil,
		save:         nil,
		del:          "",
		call:         nil,
		exit:         nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithStackframe adds a stackframe to the builder
func (app *builder) WithStackframe(stackframe stackframe.Stackframe) Builder {
	app.stackframe = stackframe
	return app
}

// WithTransform adds a transform to the builder
func (app *builder) WithTransform(transform transform.Transform) Builder {
	app.transform = transform
	return app
}

// WithVariableName adds a variableName to the builder
func (app *builder) WithVariableName(variableName variablename.VariableName) Builder {
	app.variableName = variableName
	return app
}

// WithCondition adds a condition to the builder
func (app *builder) WithCondition(condition condition.Condition) Builder {
	app.condition = condition
	return app
}

// WithStandard adds a standard to the builder
func (app *builder) WithStandard(standard standard.Standard) Builder {
	app.standard = standard
	return app
}

// WithRemaining adds a remaining to the builder
func (app *builder) WithRemaining(remaining remaining.Remaining) Builder {
	app.remaining = remaining
	return app
}

// WithValue adds a value to the builder
func (app *builder) WithValue(value value.Value) Builder {
	app.value = value
	return app
}

// WithInsert adds an insert to the builder
func (app *builder) WithInsert(insert variable.Variable) Builder {
	app.insert = insert
	return app
}

// WithSave adds a save to the builder
func (app *builder) WithSave(save variable.Variable) Builder {
	app.save = save
	return app
}

// WithDelete adds a delete to the builder
func (app *builder) WithDelete(del string) Builder {
	app.del = del
	return app
}

// WithCall adds a call to the builder
func (app *builder) WithCall(call call.Call) Builder {
	app.call = call
	return app
}

// WithExit adds an exit to the builder
func (app *builder) WithExit(exit exit.Exit) Builder {
	app.exit = exit
	return app
}

// Now builds a new Instruction instance
func (app *builder) Now() (Instruction, error) {
	if app.stackframe != nil {
		return createInstructionWithStackframe(app.stackframe), nil
	}

	if app.transform != nil {
		return createInstructionWithTransform(app.transform), nil
	}

	if app.variableName != nil {
		return createInstructionWithVariableName(app.variableName), nil
	}

	if app.condition != nil {
		return createInstructionWithCondition(app.condition), nil
	}

	if app.standard != nil {
		return createInstructionWithStandard(app.standard), nil
	}

	if app.remaining != nil {
		return createInstructionWithRemaining(app.remaining), nil
	}

	if app.value != nil {
		return createInstructionWithValue(app.value), nil
	}

	if app.insert != nil {
		return createInstructionWithInsert(app.insert), nil
	}

	if app.save != nil {
		return createInstructionWithSave(app.save), nil
	}

	if app.del != "" {
		return createInstructionWithDelete(app.del), nil
	}

	if app.call != nil {
		return createInstructionWithCall(app.call), nil
	}

	if app.exit != nil {
		return createInstructionWithExit(app.exit), nil
	}

	return nil, errors.New("the Instruction is invalid")
}
