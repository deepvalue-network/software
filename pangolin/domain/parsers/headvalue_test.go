package parsers

import (
	"testing"
)

func Test_headValue_withName_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("headValue", grammarFile)

	file := "./tests/codes/headvalue/name.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	val := ins.(HeadValue)
	if !val.IsName() {
		t.Errorf("the HeadValue was expected to be a name")
		return
	}

	if val.Name() != "my_name" {
		t.Errorf("the name was expected to be %s, %s returned", "my_name", val.Name())
		return
	}

	if val.IsVersion() {
		t.Errorf("the HeadValue was NOT expected to be a version")
		return
	}

	if val.IsImport() {
		t.Errorf("the HeadValue was NOT expected to be an import")
		return
	}
}

func Test_headValue_withVersion_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("headValue", grammarFile)

	file := "./tests/codes/headvalue/version.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	val := ins.(HeadValue)
	if val.IsName() {
		t.Errorf("the HeadValue was NOT expected to be a name")
		return
	}

	if !val.IsVersion() {
		t.Errorf("the HeadValue was expected to be a version")
		return
	}

	if val.Version() != "2020.03.04" {
		t.Errorf("the version was expected to be %s, %s returned", "2020.03.04", val.Version())
		return
	}

	if val.IsImport() {
		t.Errorf("the HeadValue was NOT expected to be an import")
		return
	}
}

func Test_headValue_withImport_Single_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("headValue", grammarFile)

	file := "./tests/codes/headvalue/import_single.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	val := ins.(HeadValue)
	if val.IsName() {
		t.Errorf("the HeadValue was NOT expected to be a name")
		return
	}

	if val.IsVersion() {
		t.Errorf("the HeadValue was NOT expected to be a version")
		return
	}

	if !val.IsImport() {
		t.Errorf("the HeadValue was expected to be an import")
		return
	}

	imp := val.Import()
	if len(imp) != 1 {
		t.Errorf("%d imports were expected, %d returned", 1, len(imp))
		return
	}
}

func Test_headValue_withImport_Multiple_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("headValue", grammarFile)

	file := "./tests/codes/headvalue/import_multiple.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	val := ins.(HeadValue)
	if val.IsName() {
		t.Errorf("the HeadValue was NOT expected to be a name")
		return
	}

	if val.IsVersion() {
		t.Errorf("the HeadValue was NOT expected to be a version")
		return
	}

	if !val.IsImport() {
		t.Errorf("the HeadValue was expected to be an import")
		return
	}

	imp := val.Import()
	if len(imp) != 2 {
		t.Errorf("%d imports were expected, %d returned", 2, len(imp))
		return
	}
}
