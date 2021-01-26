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

func Test_testInstruction_withStart_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("testInstruction", grammarFile)

	file := "./tests/codes/testinstruction/start.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	testIns := ins.(TestInstruction)
	if !testIns.IsStart() {
		t.Errorf("the testInstruction was expected to contain a Start instruction")
		return
	}
}

func Test_testInstruction_withStop_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("testInstruction", grammarFile)

	file := "./tests/codes/testinstruction/stop.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	testIns := ins.(TestInstruction)
	if !testIns.IsStop() {
		t.Errorf("the testInstruction was expected to contain a Stop instruction")
		return
	}
}
