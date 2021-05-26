package parsers

type instruction struct {
	variable   Variable
	operation  Operation
	print      Print
	stackFrame StackFrame
	jmp        Jump
	exit       Exit
	call       Call
	reg        Registry
}

func createInstructionWithVariable(variable Variable) Instruction {
	return createInstructionInternally(variable, nil, nil, nil, nil, nil, nil, nil)
}

func createInstructionWithOperation(operation Operation) Instruction {
	return createInstructionInternally(nil, operation, nil, nil, nil, nil, nil, nil)
}

func createInstructionWithPrint(print Print) Instruction {
	return createInstructionInternally(nil, nil, print, nil, nil, nil, nil, nil)
}

func createInstructionWithStackFrame(stackFrame StackFrame) Instruction {
	return createInstructionInternally(nil, nil, nil, stackFrame, nil, nil, nil, nil)
}

func createInstructionWithJump(jmp Jump) Instruction {
	return createInstructionInternally(nil, nil, nil, nil, jmp, nil, nil, nil)
}

func createInstructionWithExit(exit Exit) Instruction {
	return createInstructionInternally(nil, nil, nil, nil, nil, exit, nil, nil)
}

func createInstructionWithCall(call Call) Instruction {
	return createInstructionInternally(nil, nil, nil, nil, nil, nil, call, nil)
}

func createInstructionWithRegistry(reg Registry) Instruction {
	return createInstructionInternally(nil, nil, nil, nil, nil, nil, nil, reg)
}

func createInstructionInternally(
	variable Variable,
	operation Operation,
	print Print,
	stackFrame StackFrame,
	jmp Jump,
	exit Exit,
	call Call,
	reg Registry,
) Instruction {
	out := instruction{
		variable:   variable,
		operation:  operation,
		print:      print,
		stackFrame: stackFrame,
		jmp:        jmp,
		exit:       exit,
		call:       call,
		reg:        reg,
	}

	return &out
}

// IsVariable retruns true if the instructicallon is a variable, false otherwise
func (obj *instruction) IsVariable() bool {
	return obj.variable != nil
}

// Variable returns the variable, if any
func (obj *instruction) Variable() Variable {
	return obj.variable
}

// IsOperation retruns true if the instruction is an operation, false otherwise
func (obj *instruction) IsOperation() bool {
	return obj.operation != nil
}

// Operation returns the operation, if any
func (obj *instruction) Operation() Operation {
	return obj.operation
}

// IsPrint retruns true if the instruction is a print, false otherwise
func (obj *instruction) IsPrint() bool {
	return obj.print != nil
}

// Print returns the print, if any
func (obj *instruction) Print() Print {
	return obj.print
}

// IsStackFrame retruns true if the instruction is a stackFrame, false otherwise
func (obj *instruction) IsStackFrame() bool {
	return obj.stackFrame != nil
}

// StackFrame returns the stackFrame, if any
func (obj *instruction) StackFrame() StackFrame {
	return obj.stackFrame
}

// IsJump retruns true if the instruction is a jump, false otherwise
func (obj *instruction) IsJump() bool {
	return obj.jmp != nil
}

// Jump returns the jump, if any
func (obj *instruction) Jump() Jump {
	return obj.jmp
}

// IsExit retruns true if the instruction is an exit, false otherwise
func (obj *instruction) IsExit() bool {
	return obj.exit != nil
}

// Exit returns the exit, if any
func (obj *instruction) Exit() Exit {
	return obj.exit
}

// IsCall retruns true if the instruction is a call, false otherwise
func (obj *instruction) IsCall() bool {
	return obj.call != nil
}

// Call returns the call, if any
func (obj *instruction) Call() Call {
	return obj.call
}

// IsRegistry retruns true if the instruction is a registry, false otherwise
func (obj *instruction) IsRegistry() bool {
	return obj.reg != nil
}

// Registry returns the registry, if any
func (obj *instruction) Registry() Registry {
	return obj.reg
}
