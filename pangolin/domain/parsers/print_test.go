package parsers

import (
	"testing"
)

func Test_print_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("print", grammarFile)

	file := "./tests/codes/print/print.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	prt := ins.(Print)
	value := prt.Value()
	if !value.IsVariable() {
		t.Errorf("the print's value was expected to be a variable")
		return
	}

	variable := value.Variable().Local()
	if variable != "myVariable" {
		t.Errorf("the print's variable was expected to be %s, %s returned", "myVariable", variable)
		return
	}
}
