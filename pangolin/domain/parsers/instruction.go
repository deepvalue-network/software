package parsers

type instruction struct {
	variable   Variable
	operation  Operation
	print      Print
	stackFrame StackFrame
	jmp        Jump
	match      Match
	exit       Exit
	call       Call
	token      Token
	trigger    Trigger
	format     Format
}

func createInstructionWithVariable(variable Variable) Instruction {
	return createInstructionInternally(variable, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
}

func createInstructionWithOperation(operation Operation) Instruction {
	return createInstructionInternally(nil, operation, nil, nil, nil, nil, nil, nil, nil, nil, nil)
}

func createInstructionWithPrint(print Print) Instruction {
	return createInstructionInternally(nil, nil, print, nil, nil, nil, nil, nil, nil, nil, nil)
}

func createInstructionWithStackFrame(stackFrame StackFrame) Instruction {
	return createInstructionInternally(nil, nil, nil, stackFrame, nil, nil, nil, nil, nil, nil, nil)
}

func createInstructionWithJump(jmp Jump) Instruction {
	return createInstructionInternally(nil, nil, nil, nil, jmp, nil, nil, nil, nil, nil, nil)
}

func createInstructionWithMatch(match Match) Instruction {
	return createInstructionInternally(nil, nil, nil, nil, nil, match, nil, nil, nil, nil, nil)
}

func createInstructionWithExit(exit Exit) Instruction {
	return createInstructionInternally(nil, nil, nil, nil, nil, nil, exit, nil, nil, nil, nil)
}

func createInstructionWithCall(call Call) Instruction {
	return createInstructionInternally(nil, nil, nil, nil, nil, nil, nil, call, nil, nil, nil)
}

func createInstructionWithToken(token Token) Instruction {
	return createInstructionInternally(nil, nil, nil, nil, nil, nil, nil, nil, token, nil, nil)
}

func createInstructionWithTrigger(trigger Trigger) Instruction {
	return createInstructionInternally(nil, nil, nil, nil, nil, nil, nil, nil, nil, trigger, nil)
}

func createInstructionWithFormat(format Format) Instruction {
	return createInstructionInternally(nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, format)
}

func createInstructionInternally(
	variable Variable,
	operation Operation,
	print Print,
	stackFrame StackFrame,
	jmp Jump,
	match Match,
	exit Exit,
	call Call,
	token Token,
	trigger Trigger,
	format Format,
) Instruction {
	out := instruction{
		variable:   variable,
		operation:  operation,
		print:      print,
		stackFrame: stackFrame,
		jmp:        jmp,
		match:      match,
		exit:       exit,
		call:       call,
		token:      token,
		trigger:    trigger,
		format:     format,
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

// IsMatch retruns true if the instruction is a match, false otherwise
func (obj *instruction) IsMatch() bool {
	return obj.match != nil
}

// Match returns the match, if any
func (obj *instruction) Match() Match {
	return obj.match
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

// IsToken returns true if there is a token, false otherwise
func (obj *instruction) IsToken() bool {
	return obj.token != nil
}

// Token returns the token, if any
func (obj *instruction) Token() Token {
	return obj.token
}

// IsTrigger returns true if there is a trigger, false otherwise
func (obj *instruction) IsTrigger() bool {
	return obj.trigger != nil
}

// Trigger returns the trigger, if any
func (obj *instruction) Trigger() Trigger {
	return obj.trigger
}

// IsFormat returns true if there is a format, false otherwise
func (obj *instruction) IsFormat() bool {
	return obj.format != nil
}

// Format returns the format, if any
func (obj *instruction) Format() Format {
	return obj.format
}
