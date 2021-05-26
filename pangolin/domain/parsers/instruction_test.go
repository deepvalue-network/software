package parsers

import (
	"testing"
)

func Test_instruction_operation_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("instruction", grammarFile)

	file := "./tests/codes/instruction/operation.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	instruction := ins.(Instruction)
	if !instruction.IsOperation() {
		t.Errorf("the Instruction was expected to contain an Operation")
		return
	}
}

func Test_instruction_print_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("instruction", grammarFile)

	file := "./tests/codes/instruction/print.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	instruction := ins.(Instruction)
	if !instruction.IsPrint() {
		t.Errorf("the Instruction was expected to contain a Print")
		return
	}
}

func Test_instruction_stackFrame_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("instruction", grammarFile)

	file := "./tests/codes/instruction/stackframe.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	instruction := ins.(Instruction)
	if !instruction.IsStackFrame() {
		t.Errorf("the Instruction was expected to contain a Stackframe")
		return
	}
}

func Test_instruction_variable_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("instruction", grammarFile)

	file := "./tests/codes/instruction/variable.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	instruction := ins.(Instruction)
	if !instruction.IsVariable() {
		t.Errorf("the Instruction was expected to contain a Variable")
		return
	}
}

func Test_instruction_jump_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("instruction", grammarFile)

	file := "./tests/codes/instruction/jump.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	instruction := ins.(Instruction)
	if !instruction.IsJump() {
		t.Errorf("the Instruction was expected to contain a Jump")
		return
	}
}

func Test_instruction_exit_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("instruction", grammarFile)

	file := "./tests/codes/instruction/exit.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	instruction := ins.(Instruction)
	if !instruction.IsExit() {
		t.Errorf("the Instruction was expected to contain an exit")
		return
	}
}

func Test_instruction_call_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("instruction", grammarFile)

	file := "./tests/codes/instruction/call.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	instruction := ins.(Instruction)
	if !instruction.IsCall() {
		t.Errorf("the Instruction was expected to contain a call")
		return
	}
}

func Test_instruction_registry_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("instruction", grammarFile)

	file := "./tests/codes/instruction/registry.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	instruction := ins.(Instruction)
	if !instruction.IsRegistry() {
		t.Errorf("the Instruction was expected to contain a registry")
		return
	}
}
