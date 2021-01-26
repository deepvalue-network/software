package parsers

import (
	"testing"
)

func Test_language_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("language", grammarFile)

	file := "./tests/codes/language/simple.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	lang := ins.(Language)
	if lang.Tokens().String() != "./my/tokens.json" {
		t.Errorf("the tokens was expected to be %s, %s returned", "./my/tokens.json", lang.Tokens().String())
		return
	}

	if lang.Rules().String() != "./my/rules.json" {
		t.Errorf("the rules was expected to be %s, %s returned", "./my/rules.json", lang.Rules().String())
		return
	}

	if lang.Logic().String() != "./my/logic.json" {
		t.Errorf("the logic was expected to be %s, %s returned", "./my/logic.json", lang.Logic().String())
		return
	}

	if lang.Input() != "MyInputVar" {
		t.Errorf("the input variable was expected to be %s, %s returned", "MyInputVar", lang.Input())
		return
	}

	if lang.Output() != "MyOutputVar" {
		t.Errorf("the input variable was expected to be %s, %s returned", "MyOutputVar", lang.Output())
		return
	}

	if lang.HasChannels() {
		t.Errorf("the language was NOT expecting channels")
		return
	}

	if lang.HasExtends() {
		t.Errorf("the language was NOT expecting extends")
		return
	}
}

func Test_language_withExtends_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("language", grammarFile)

	file := "./tests/codes/language/with_extends.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	lang := ins.(Language)
	if lang.Tokens().String() != "./my/tokens.json" {
		t.Errorf("the tokens was expected to be %s, %s returned", "./my/tokens.json", lang.Tokens().String())
		return
	}

	if lang.Rules().String() != "./my/rules.json" {
		t.Errorf("the rules was expected to be %s, %s returned", "./my/rules.json", lang.Rules().String())
		return
	}

	if lang.Logic().String() != "./my/logic.json" {
		t.Errorf("the logic was expected to be %s, %s returned", "./my/logic.json", lang.Logic().String())
		return
	}

	if lang.Input() != "MyInputVar" {
		t.Errorf("the input variable was expected to be %s, %s returned", "MyInputVar", lang.Input())
		return
	}

	if lang.Output() != "MyOutputVar" {
		t.Errorf("the input variable was expected to be %s, %s returned", "MyOutputVar", lang.Output())
		return
	}

	if lang.HasChannels() {
		t.Errorf("the language was NOT expecting channels")
		return
	}

	if !lang.HasExtends() {
		t.Errorf("the language was expecting extends")
		return
	}

	extends := lang.Extends()
	if len(extends) != 1 {
		t.Errorf("%d extends were expected, %d returned", 1, len(extends))
		return
	}
}

func Test_language_withChannels_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("language", grammarFile)

	file := "./tests/codes/language/with_channels.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	lang := ins.(Language)
	if lang.Tokens().String() != "./my/tokens.json" {
		t.Errorf("the tokens was expected to be %s, %s returned", "./my/tokens.json", lang.Tokens().String())
		return
	}

	if lang.Rules().String() != "./my/rules.json" {
		t.Errorf("the rules was expected to be %s, %s returned", "./my/rules.json", lang.Rules().String())
		return
	}

	if lang.Logic().String() != "./my/logic.json" {
		t.Errorf("the logic was expected to be %s, %s returned", "./my/logic.json", lang.Logic().String())
		return
	}

	if lang.Input() != "MyInputVar" {
		t.Errorf("the input variable was expected to be %s, %s returned", "MyInputVar", lang.Input())
		return
	}

	if lang.Output() != "MyOutputVar" {
		t.Errorf("the input variable was expected to be %s, %s returned", "MyOutputVar", lang.Output())
		return
	}

	if !lang.HasChannels() {
		t.Errorf("the language was expecting a channels filepath")
		return
	}

	if lang.Channels().String() != "./my/channels.json" {
		t.Errorf("the channels was expected to be %s, %s returned", "./my/channels.json", lang.Channels().String())
		return
	}

	if lang.HasExtends() {
		t.Errorf("the language was NOT expecting extends")
		return
	}
}

func Test_language_withChannels_withExtends_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("language", grammarFile)

	file := "./tests/codes/language/with_channels_with_extends.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	lang := ins.(Language)
	if lang.Tokens().String() != "./my/tokens.json" {
		t.Errorf("the tokens was expected to be %s, %s returned", "./my/tokens.json", lang.Tokens().String())
		return
	}

	if lang.Rules().String() != "./my/rules.json" {
		t.Errorf("the rules was expected to be %s, %s returned", "./my/rules.json", lang.Rules().String())
		return
	}

	if lang.Logic().String() != "./my/logic.json" {
		t.Errorf("the logic was expected to be %s, %s returned", "./my/logic.json", lang.Logic().String())
		return
	}

	if lang.Input() != "MyInputVar" {
		t.Errorf("the input variable was expected to be %s, %s returned", "MyInputVar", lang.Input())
		return
	}

	if lang.Output() != "MyOutputVar" {
		t.Errorf("the input variable was expected to be %s, %s returned", "MyOutputVar", lang.Output())
		return
	}

	if !lang.HasChannels() {
		t.Errorf("the language was expecting a channels filepath")
		return
	}

	if lang.Channels().String() != "./my/channels.json" {
		t.Errorf("the channels was expected to be %s, %s returned", "./my/channels.json", lang.Channels().String())
		return
	}

	if !lang.HasExtends() {
		t.Errorf("the language was expecting extends")
		return
	}

	extends := lang.Extends()
	if len(extends) != 3 {
		t.Errorf("%d extends were expected, %d returned", 3, len(extends))
		return
	}
}
