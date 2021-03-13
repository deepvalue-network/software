package parsers

import (
	"testing"
)

func Test_format_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("format", grammarFile)

	file := "./tests/codes/format/simple.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	format := ins.(Format)
	if format.Results().String() != "myResults" {
		t.Errorf("the results was expected to be %s, %s returned", format.Results().String(), "myResults")
		return
	}

	if format.Pattern().String() != "myPattern" {
		t.Errorf("the pattern was expected to be %s, %s returned", format.Pattern().String(), "myPattern")
		return
	}

	if format.First().String() != "firstVal" {
		t.Errorf("the first was expected to be %s, %s returned", format.First().String(), "firstVal")
		return
	}

	if format.Second().String() != "secondVal" {
		t.Errorf("the second was expected to be %s, %s returned", format.Second().String(), "secondVal")
		return
	}

}
