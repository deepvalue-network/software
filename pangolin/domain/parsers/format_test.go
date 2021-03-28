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
	if format.Results() != "myResults" {
		t.Errorf("the results was expected to be %s, %s returned", format.Results(), "myResults")
		return
	}

	if format.Pattern() != "myPattern" {
		t.Errorf("the pattern was expected to be %s, %s returned", format.Pattern(), "myPattern")
		return
	}

	if format.First() != "firstVal" {
		t.Errorf("the first was expected to be %s, %s returned", format.First(), "firstVal")
		return
	}

	if format.Second() != "secondVal" {
		t.Errorf("the second was expected to be %s, %s returned", format.Second(), "secondVal")
		return
	}

}
