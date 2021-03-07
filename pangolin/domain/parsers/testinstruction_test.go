package parsers

import (
	"testing"
)

func Test_testInstruction_withInstruction_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("testInstruction", grammarFile)

	file := "./tests/codes/testinstruction/instruction.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	testIns := ins.(TestInstruction)
	if !testIns.IsInstruction() {
		t.Errorf("the testInstruction was expected to contain an Instruction instance")
		return
	}
}

func Test_testInstruction_withReadFile_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("testInstruction", grammarFile)

	file := "./tests/codes/testinstruction/readfile.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	testIns := ins.(TestInstruction)
	if !testIns.IsReadFile() {
		t.Errorf("the testInstruction was expected to contain a Readfile instance")
		return
	}
}

func Test_testInstruction_withAssert_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("testInstruction", grammarFile)

	file := "./tests/codes/testinstruction/assert.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	testIns := ins.(TestInstruction)
	if !testIns.IsAssert() {
		t.Errorf("the testInstruction was expected to contain an Assert instruction")
		return
	}
}
