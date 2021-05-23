package parsers

import (
	"testing"
)

func Test_scriptValue_withName_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("scriptValue", grammarFile)

	file := "./tests/codes/scriptvalue/name.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	value := ins.(ScriptValue)
	if !value.IsName() {
		t.Errorf("the scriptValue was expecting to be a name")
		return
	}

	if value.Name() != "my_name" {
		t.Errorf("the name was expected to be %s, %s returned", "my_name", value.Name())
		return
	}

	if value.IsVersion() {
		t.Errorf("the scriptValue was NOT expecting a version")
		return
	}

	if value.IsLanguage() {
		t.Errorf("the scriptValue was NOT expecting a language")
		return
	}

	if value.IsScript() {
		t.Errorf("the scriptValue was NOT expecting a script")
		return
	}

	if value.IsOutput() {
		t.Errorf("the scriptValue was NOT expecting an output")
		return
	}

	if value.IsScriptTests() {
		t.Errorf("the scriptValue was NOT expecting as ScriptTests")
		return
	}
}

func Test_scriptValue_withVersion_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("scriptValue", grammarFile)

	file := "./tests/codes/scriptvalue/version.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	value := ins.(ScriptValue)
	if value.IsName() {
		t.Errorf("the scriptValue was NOT expecting to be a name")
		return
	}

	if !value.IsVersion() {
		t.Errorf("the scriptValue was expecting a version")
		return
	}

	if value.Version() != "2019.03.11" {
		t.Errorf("the name was expected to be %s, %s returned", "2019.03.11", value.Name())
		return
	}

	if value.IsLanguage() {
		t.Errorf("the scriptValue was NOT expecting a language")
		return
	}

	if value.IsScript() {
		t.Errorf("the scriptValue was NOT expecting a script")
		return
	}

	if value.IsOutput() {
		t.Errorf("the scriptValue was NOT expecting an output")
		return
	}

	if value.IsScriptTests() {
		t.Errorf("the scriptValue was NOT expecting as ScriptTests")
		return
	}
}

func Test_scriptValue_withLanguage_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("scriptValue", grammarFile)

	file := "./tests/codes/scriptvalue/language.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	value := ins.(ScriptValue)
	if value.IsName() {
		t.Errorf("the scriptValue was NOT expecting to be a name")
		return
	}

	if value.IsVersion() {
		t.Errorf("the scriptValue was NOT expecting a version")
		return
	}

	if !value.IsLanguage() {
		t.Errorf("the scriptValue was expecting a language")
		return
	}

	if value.Language().String() != "./my/path/lang.rod" {
		t.Errorf("the name was expected to be %s, %s returned", "./my/path/lang.rod", value.Language().String())
		return
	}

	if value.IsScript() {
		t.Errorf("the scriptValue was NOT expecting a script")
		return
	}

	if value.IsOutput() {
		t.Errorf("the scriptValue was NOT expecting an output")
		return
	}

	if value.IsScriptTests() {
		t.Errorf("the scriptValue was NOT expecting as ScriptTests")
		return
	}
}

func Test_scriptValue_withScript_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("scriptValue", grammarFile)

	file := "./tests/codes/scriptvalue/script.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	value := ins.(ScriptValue)
	if value.IsName() {
		t.Errorf("the scriptValue was NOT expecting to be a name")
		return
	}

	if value.IsVersion() {
		t.Errorf("the scriptValue was NOT expecting a version")
		return
	}

	if value.IsLanguage() {
		t.Errorf("the scriptValue was NOT expecting a language")
		return
	}

	if !value.IsScript() {
		t.Errorf("the scriptValue was expecting a script")
		return
	}

	if value.Script().String() != "./my/path/script.rod" {
		t.Errorf("the name was expected to be %s, %s returned", "./my/path/script.rod", value.Script().String())
		return
	}

	if value.IsOutput() {
		t.Errorf("the scriptValue was NOT expecting an output")
		return
	}

	if value.IsScriptTests() {
		t.Errorf("the scriptValue was NOT expecting as ScriptTests")
		return
	}
}

func Test_scriptValue_withOutput_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("scriptValue", grammarFile)

	file := "./tests/codes/scriptvalue/output.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	value := ins.(ScriptValue)
	if value.IsName() {
		t.Errorf("the scriptValue was NOT expecting to be a name")
		return
	}

	if value.IsVersion() {
		t.Errorf("the scriptValue was NOT expecting a version")
		return
	}

	if value.IsLanguage() {
		t.Errorf("the scriptValue was NOT expecting a language")
		return
	}

	if value.IsScript() {
		t.Errorf("the scriptValue was NOT expecting a script")
		return
	}

	if !value.IsOutput() {
		t.Errorf("the scriptValue was expecting an output variable")
		return
	}

	if value.Output() != "myOutput" {
		t.Errorf("the name was expected to be %s, %s returned", "myOutput", value.Output())
		return
	}

	if value.IsScriptTests() {
		t.Errorf("the scriptValue was NOT expecting as ScriptTests")
		return
	}
}

func Test_scriptValue_withScriptTests_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("scriptValue", grammarFile)

	file := "./tests/codes/scriptvalue/scripttests.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	value := ins.(ScriptValue)
	if value.IsName() {
		t.Errorf("the scriptValue was NOT expecting to be a name")
		return
	}

	if value.IsVersion() {
		t.Errorf("the scriptValue was NOT expecting a version")
		return
	}

	if value.IsLanguage() {
		t.Errorf("the scriptValue was NOT expecting a language")
		return
	}

	if value.IsScript() {
		t.Errorf("the scriptValue was NOT expecting a script")
		return
	}

	if value.IsOutput() {
		t.Errorf("the scriptValue was NOT expecting an output variable")
		return
	}

	if !value.IsScriptTests() {
		t.Errorf("the scriptValue was expecting as ScriptTests")
		return
	}
}
