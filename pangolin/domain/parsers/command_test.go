package parsers

import (
	"testing"
)

func Test_commmand_withHead_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("command", grammarFile)

	file := "./tests/codes/command/head.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	cmd := ins.(Command)
	if !cmd.IsHead() {
		t.Errorf("the command was expected to be a HeadCommand")
		return
	}

	if cmd.IsLabel() {
		t.Errorf("the command was NOT expected to be a LabelCommand")
		return
	}

	if cmd.IsLanguage() {
		t.Errorf("the command was NOT expected to be a LanguageCommand")
		return
	}

	if cmd.IsMain() {
		t.Errorf("the command was NOT expected to be a MainCommand")
		return
	}

	if cmd.IsScript() {
		t.Errorf("the command was NOT expected to be a ScriptCommand")
		return
	}

	if cmd.IsTest() {
		t.Errorf("the command was NOT expected to be a TestCommand")
		return
	}
}

func Test_commmand_withLabel_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("command", grammarFile)

	file := "./tests/codes/command/label.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	cmd := ins.(Command)
	if cmd.IsHead() {
		t.Errorf("the command was NOT expected to be a HeadCommand")
		return
	}

	if !cmd.IsLabel() {
		t.Errorf("the command was expected to be a LabelCommand")
		return
	}

	if cmd.IsLanguage() {
		t.Errorf("the command was NOT expected to be a LanguageCommand")
		return
	}

	if cmd.IsMain() {
		t.Errorf("the command was NOT expected to be a MainCommand")
		return
	}

	if cmd.IsScript() {
		t.Errorf("the command was NOT expected to be a ScriptCommand")
		return
	}

	if cmd.IsTest() {
		t.Errorf("the command was NOT expected to be a TestCommand")
		return
	}
}

func Test_commmand_withLanguage_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("command", grammarFile)

	file := "./tests/codes/command/language.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	cmd := ins.(Command)
	if cmd.IsHead() {
		t.Errorf("the command was NOT expected to be a HeadCommand")
		return
	}

	if cmd.IsLabel() {
		t.Errorf("the command was NOT expected to be a LabelCommand")
		return
	}

	if !cmd.IsLanguage() {
		t.Errorf("the command was expected to be a LanguageCommand")
		return
	}

	if cmd.IsMain() {
		t.Errorf("the command was NOT expected to be a MainCommand")
		return
	}

	if cmd.IsScript() {
		t.Errorf("the command was NOT expected to be a ScriptCommand")
		return
	}

	if cmd.IsTest() {
		t.Errorf("the command was NOT expected to be a TestCommand")
		return
	}
}

func Test_commmand_withMain_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("command", grammarFile)

	file := "./tests/codes/command/main.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	cmd := ins.(Command)
	if cmd.IsHead() {
		t.Errorf("the command was NOT expected to be a HeadCommand")
		return
	}

	if cmd.IsLabel() {
		t.Errorf("the command was NOT expected to be a LabelCommand")
		return
	}

	if cmd.IsLanguage() {
		t.Errorf("the command was NOT expected to be a LanguageCommand")
		return
	}

	if !cmd.IsMain() {
		t.Errorf("the command was expected to be a MainCommand")
		return
	}

	if cmd.IsScript() {
		t.Errorf("the command was NOT expected to be a ScriptCommand")
		return
	}

	if cmd.IsTest() {
		t.Errorf("the command was NOT expected to be a TestCommand")
		return
	}
}

func Test_commmand_withScript_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("command", grammarFile)

	file := "./tests/codes/command/script.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	cmd := ins.(Command)
	if cmd.IsHead() {
		t.Errorf("the command was NOT expected to be a HeadCommand")
		return
	}

	if cmd.IsLabel() {
		t.Errorf("the command was NOT expected to be a LabelCommand")
		return
	}

	if cmd.IsLanguage() {
		t.Errorf("the command was NOT expected to be a LanguageCommand")
		return
	}

	if cmd.IsMain() {
		t.Errorf("the command was NOT expected to be a MainCommand")
		return
	}

	if !cmd.IsScript() {
		t.Errorf("the command was expected to be a ScriptCommand")
		return
	}

	if cmd.IsTest() {
		t.Errorf("the command was NOT expected to be a TestCommand")
		return
	}
}

func Test_commmand_withTest_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("command", grammarFile)

	file := "./tests/codes/command/test.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	cmd := ins.(Command)
	if cmd.IsHead() {
		t.Errorf("the command was NOT expected to be a HeadCommand")
		return
	}

	if cmd.IsLabel() {
		t.Errorf("the command was NOT expected to be a LabelCommand")
		return
	}

	if cmd.IsLanguage() {
		t.Errorf("the command was NOT expected to be a LanguageCommand")
		return
	}

	if cmd.IsMain() {
		t.Errorf("the command was NOT expected to be a MainCommand")
		return
	}

	if cmd.IsScript() {
		t.Errorf("the command was NOT expected to be a ScriptCommand")
		return
	}

	if !cmd.IsTest() {
		t.Errorf("the command was expected to be a TestCommand")
		return
	}
}
