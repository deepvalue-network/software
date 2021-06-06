package instruction

import (
	"errors"

	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction/call"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction/condition"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction/exit"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction/module"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction/registry"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction/remaining"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction/stackframe"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction/standard"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction/value"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction/variable"
)

type builder struct {
	stackframe stackframe.Stackframe
	condition  condition.Condition
	standard   standard.Standard
	remaining  remaining.Remaining
	value      value.Value
	insert     variable.Variable
	save       variable.Variable
	del        string
	call       call.Call
	module     module.Module
	exit       exit.Exit
	reg        registry.Registry
}

func createBuilder() Builder {
	out := builder{
		stackframe: nil,
		condition:  nil,
		standard:   nil,
		remaining:  nil,
		value:      nil,
		insert:     nil,
		save:       nil,
		del:        "",
		call:       nil,
		module:     nil,
		exit:       nil,
		reg:        nil,
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

// WithModule adds a module to the builder
func (app *builder) WithModule(module module.Module) Builder {
	app.module = module
	return app
}

// WithExit adds an exit to the builder
func (app *builder) WithExit(exit exit.Exit) Builder {
	app.exit = exit
	return app
}

// WithRegistry adds a registry to the builder
func (app *builder) WithRegistry(reg registry.Registry) Builder {
	app.reg = reg
	return app
}

// Now builds a new Instruction instance
func (app *builder) Now() (Instruction, error) {
	if app.stackframe != nil {
		return createInstructionWithStackframe(app.stackframe), nil
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

	if app.module != nil {
		return createInstructionWithModule(app.module), nil
	}

	if app.exit != nil {
		return createInstructionWithExit(app.exit), nil
	}

	if app.reg != nil {
		return createInstructionWithRegistry(app.reg), nil
	}

	return nil, errors.New("the Instruction is invalid")
}
