package parsers

import (
	"errors"
)

type instructionBuilder struct {
	variable   Variable
	operation  Operation
	print      Print
	stackFrame StackFrame
	jmp        Jump
	exit       Exit
	call       Call
	reg        Registry
	swtch      Switch
	save       Save
}

func createInstructionBuilder() InstructionBuilder {
	out := instructionBuilder{
		variable:   nil,
		operation:  nil,
		print:      nil,
		stackFrame: nil,
		jmp:        nil,
		exit:       nil,
		call:       nil,
		reg:        nil,
		swtch:      nil,
		save:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *instructionBuilder) Create() InstructionBuilder {
	return createInstructionBuilder()
}

// WithVariable adds a variable to the builder
func (app *instructionBuilder) WithVariable(variable Variable) InstructionBuilder {
	app.variable = variable
	return app
}

// WithOperation adds an operation to the builder
func (app *instructionBuilder) WithOperation(operation Operation) InstructionBuilder {
	app.operation = operation
	return app
}

// WithPrint adds a print to the builder
func (app *instructionBuilder) WithPrint(print Print) InstructionBuilder {
	app.print = print
	return app
}

// WithStackFrame adds a stackFrame to the builder
func (app *instructionBuilder) WithStackFrame(stackFrame StackFrame) InstructionBuilder {
	app.stackFrame = stackFrame
	return app
}

// WithJump adds a jump to the builder
func (app *instructionBuilder) WithJump(jmp Jump) InstructionBuilder {
	app.jmp = jmp
	return app
}

// WithExit adds an exit to the builder
func (app *instructionBuilder) WithExit(exit Exit) InstructionBuilder {
	app.exit = exit
	return app
}

// WithCall adds a call to the builder
func (app *instructionBuilder) WithCall(call Call) InstructionBuilder {
	app.call = call
	return app
}

// WithRegistry adds a registry to the builder
func (app *instructionBuilder) WithRegistry(registry Registry) InstructionBuilder {
	app.reg = registry
	return app
}

// WithSwitch adds a switch to the builder
func (app *instructionBuilder) WithSwitch(swtch Switch) InstructionBuilder {
	app.swtch = swtch
	return app
}

// WithSave adds a save to the builder
func (app *instructionBuilder) WithSave(save Save) InstructionBuilder {
	app.save = save
	return app
}

// Now builds an instruction instance
func (app *instructionBuilder) Now() (Instruction, error) {
	if app.variable != nil {
		return createInstructionWithVariable(app.variable), nil
	}

	if app.operation != nil {
		return createInstructionWithOperation(app.operation), nil
	}

	if app.print != nil {
		return createInstructionWithPrint(app.print), nil
	}

	if app.stackFrame != nil {
		return createInstructionWithStackFrame(app.stackFrame), nil
	}

	if app.jmp != nil {
		return createInstructionWithJump(app.jmp), nil
	}

	if app.exit != nil {
		return createInstructionWithExit(app.exit), nil
	}

	if app.call != nil {
		return createInstructionWithCall(app.call), nil
	}

	if app.reg != nil {
		return createInstructionWithRegistry(app.reg), nil
	}

	if app.swtch != nil {
		return createInstructionWithSwitch(app.swtch), nil
	}

	if app.save != nil {
		return createInstructionWithSave(app.save), nil
	}

	return nil, errors.New("the Instruction is invalid")
}
