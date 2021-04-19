package parsers

import (
	"testing"
)

func Test_languageValue_root_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("languageValue", grammarFile)

	file := "./tests/codes/languagevalue/root.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	val := ins.(LanguageValue)
	if !val.IsRoot() {
		t.Errorf("the Language was expected to be a root")
		return
	}

	root := val.Root()
	if root != "myPattern" {
		t.Errorf("the root was expected to be %s, %s returned", "myPattern", root)
		return
	}

	if val.IsTokens() {
		t.Errorf("the Language was NOT expected to be a tokens")
		return
	}

	if val.IsChannels() {
		t.Errorf("the Language was NOT expected to be a channels")
		return
	}

	if val.IsRules() {
		t.Errorf("the Language was NOT expected to be a rules")
		return
	}

	if val.IsLogic() {
		t.Errorf("the Language was NOT expected to be a logic")
		return
	}

	if val.IsInputVariable() {
		t.Errorf("the Language was NOT expected to be an inputVariable")
		return
	}

	if val.IsExtends() {
		t.Errorf("the Language was NOT expected to be an extends")
		return
	}
}

func Test_languageValue_tokens_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("languageValue", grammarFile)

	file := "./tests/codes/languagevalue/tokens.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	val := ins.(LanguageValue)
	if val.IsRoot() {
		t.Errorf("the Language was NOT expected to be a root")
		return
	}

	if !val.IsTokens() {
		t.Errorf("the Language was expected to be a tokens")
		return
	}

	if val.IsChannels() {
		t.Errorf("the Language was NOT expected to be a channels")
		return
	}

	if val.IsRules() {
		t.Errorf("the Language was NOT expected to be a rules")
		return
	}

	if val.IsLogic() {
		t.Errorf("the Language was NOT expected to be a logic")
		return
	}

	if val.IsInputVariable() {
		t.Errorf("the Language was NOT expected to be an inputVariable")
		return
	}

	if val.IsExtends() {
		t.Errorf("the Language was NOT expected to be an extends")
		return
	}
}

func Test_languageValue_channels_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("languageValue", grammarFile)

	file := "./tests/codes/languagevalue/channels.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	val := ins.(LanguageValue)
	if val.IsRoot() {
		t.Errorf("the Language was NOT expected to be a root")
		return
	}

	if val.IsTokens() {
		t.Errorf("the Language was NOT expected to be a tokens")
		return
	}

	if !val.IsChannels() {
		t.Errorf("the Language was expected to be a channels")
		return
	}

	if val.IsRules() {
		t.Errorf("the Language was NOT expected to be a rules")
		return
	}

	if val.IsLogic() {
		t.Errorf("the Language was NOT expected to be a logic")
		return
	}

	if val.IsInputVariable() {
		t.Errorf("the Language was NOT expected to be an inputVariable")
		return
	}

	if val.IsExtends() {
		t.Errorf("the Language was NOT expected to be an extends")
		return
	}
}

func Test_languageValue_rules_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("languageValue", grammarFile)

	file := "./tests/codes/languagevalue/rules.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	val := ins.(LanguageValue)
	if val.IsRoot() {
		t.Errorf("the Language was NOT expected to be a root")
		return
	}

	if val.IsTokens() {
		t.Errorf("the Language was NOT expected to be a tokens")
		return
	}

	if val.IsChannels() {
		t.Errorf("the Language was NOT expected to be a channels")
		return
	}

	if !val.IsRules() {
		t.Errorf("the Language was expected to be a rules")
		return
	}

	if val.IsLogic() {
		t.Errorf("the Language was NOT expected to be a logic")
		return
	}

	if val.IsInputVariable() {
		t.Errorf("the Language was NOT expected to be an inputVariable")
		return
	}

	if val.IsExtends() {
		t.Errorf("the Language was NOT expected to be an extends")
		return
	}
}

func Test_languageValue_logic_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("languageValue", grammarFile)

	file := "./tests/codes/languagevalue/logic.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	val := ins.(LanguageValue)
	if val.IsRoot() {
		t.Errorf("the Language was NOT expected to be a root")
		return
	}

	if val.IsTokens() {
		t.Errorf("the Language was NOT expected to be a tokens")
		return
	}

	if val.IsChannels() {
		t.Errorf("the Language was NOT expected to be a channels")
		return
	}

	if val.IsRules() {
		t.Errorf("the Language was NOT expected to be a rules")
		return
	}

	if !val.IsLogic() {
		t.Errorf("the Language was expected to be a logic")
		return
	}

	if val.IsInputVariable() {
		t.Errorf("the Language was NOT expected to be an inputVariable")
		return
	}

	if val.IsExtends() {
		t.Errorf("the Language was NOT expected to be an extends")
		return
	}
}

func Test_languageValue_inputVariable_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("languageValue", grammarFile)

	file := "./tests/codes/languagevalue/input.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	val := ins.(LanguageValue)
	if val.IsRoot() {
		t.Errorf("the Language was NOT expected to be a root")
		return
	}

	if val.IsTokens() {
		t.Errorf("the Language was NOT expected to be a tokens")
		return
	}

	if val.IsChannels() {
		t.Errorf("the Language was NOT expected to be a channels")
		return
	}

	if val.IsRules() {
		t.Errorf("the Language was NOT expected to be a rules")
		return
	}

	if val.IsLogic() {
		t.Errorf("the Language was NOT expected to be a logic")
		return
	}

	if !val.IsInputVariable() {
		t.Errorf("the Language was expected to be an inputVariable")
		return
	}

	if val.IsExtends() {
		t.Errorf("the Language was NOT expected to be an extends")
		return
	}
}

func Test_languageValue_extends_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("languageValue", grammarFile)

	file := "./tests/codes/languagevalue/extends.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	val := ins.(LanguageValue)
	if val.IsRoot() {
		t.Errorf("the Language was NOT expected to be a root")
		return
	}

	if val.IsTokens() {
		t.Errorf("the Language was NOT expected to be a tokens")
		return
	}

	if val.IsChannels() {
		t.Errorf("the Language was NOT expected to be a channels")
		return
	}

	if val.IsRules() {
		t.Errorf("the Language was NOT expected to be a rules")
		return
	}

	if val.IsLogic() {
		t.Errorf("the Language was NOT expected to be a logic")
		return
	}

	if val.IsInputVariable() {
		t.Errorf("the Language was NOT expected to be an inputVariable")
		return
	}

	if !val.IsExtends() {
		t.Errorf("the Language was expected to be an extends")
		return
	}

	extends := val.Extends()
	if len(extends) != 2 {
		t.Errorf("%d extends were expected, %d returned", 2, len(extends))
		return
	}
}

func Test_languageValue_matches_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("languageValue", grammarFile)

	file := "./tests/codes/languagevalue/matches.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	val := ins.(LanguageValue)
	if val.IsRoot() {
		t.Errorf("the Language was NOT expected to be a root")
		return
	}

	if val.IsTokens() {
		t.Errorf("the Language was NOT expected to be a tokens")
		return
	}

	if val.IsChannels() {
		t.Errorf("the Language was NOT expected to be a channels")
		return
	}

	if val.IsRules() {
		t.Errorf("the Language was NOT expected to be a rules")
		return
	}

	if val.IsLogic() {
		t.Errorf("the Language was NOT expected to be a logic")
		return
	}

	if val.IsInputVariable() {
		t.Errorf("the Language was NOT expected to be an inputVariable")
		return
	}

	if val.IsExtends() {
		t.Errorf("the Language was NOT expected to be an extends")
		return
	}

	if !val.IsPatternMatches() {
		t.Errorf("the Language was expected to be a patternMatches")
		return
	}

	matches := val.PatternMatches()
	if len(matches) != 3 {
		t.Errorf("%d extends were expected, %d returned", 3, len(matches))
		return
	}
}
