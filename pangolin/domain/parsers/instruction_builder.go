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
	match      Match
	exit       Exit
	call       Call
	token      Token
}

func createInstructionBuilder() InstructionBuilder {
	out := instructionBuilder{
		variable:   nil,
		operation:  nil,
		print:      nil,
		stackFrame: nil,
		jmp:        nil,
		match:      nil,
		exit:       nil,
		call:       nil,
		token:      nil,
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

// WithMatch adds a match to the builder
func (app *instructionBuilder) WithMatch(match Match) InstructionBuilder {
	app.match = match
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

// WithToken adds a token to the builder
func (app *instructionBuilder) WithToken(token Token) InstructionBuilder {
	app.token = token
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

	if app.match != nil {
		return createInstructionWithMatch(app.match), nil
	}

	if app.exit != nil {
		return createInstructionWithExit(app.exit), nil
	}

	if app.call != nil {
		return createInstructionWithCall(app.call), nil
	}

	if app.token != nil {
		return createInstructionWithToken(app.token), nil
	}

	return nil, errors.New("the Instruction is invalid")
}
