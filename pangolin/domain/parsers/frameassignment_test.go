package parsers

import (
	"testing"
)

func Test_frameAssignment_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("frameAssignment", grammarFile)

	file := "./tests/codes/frameassignment/all.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	ass := ins.(FrameAssignment)
	if ass.Standard() == nil {
		t.Errorf("the StandardOperation was expected to be valid")
		return
	}
}
