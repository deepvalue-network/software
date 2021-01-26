package parsers

import (
	"testing"
)

func Test_labelInstruction_instruction_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("labelInstruction", grammarFile)

	file := "./tests/codes/labelinstruction/instruction.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	instruction := ins.(LabelInstruction)
	if !instruction.IsInstruction() {
		t.Errorf("the LabelInstruction was not expected to contain an Instruction")
		return
	}
}

func Test_labelInstruction_ret_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("labelInstruction", grammarFile)

	file := "./tests/codes/labelinstruction/ret.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	instruction := ins.(LabelInstruction)
	if !instruction.IsRet() {
		t.Errorf("the LabelInstruction was not expected to contain a return")
		return
	}
}
