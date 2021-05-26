package parsers

import (
	"testing"
)

func Test_registry_withRegister_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("registry", grammarFile)

	file := "./tests/codes/registry/register.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	reg := ins.(Registry)
	if !reg.IsRegister() {
		t.Errorf("the registry was expected to contain a Register instruction")
		return
	}

	if reg.IsUnregister() {
		t.Errorf("the registry was NOT expected to contain an Unregister instruction")
		return
	}

	if reg.IsFetch() {
		t.Errorf("the registry was NOT expected to contain a Fetchregistry instruction")
		return
	}
}

func Test_registry_withUnregister_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("registry", grammarFile)

	file := "./tests/codes/registry/unregister.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	reg := ins.(Registry)
	if reg.IsRegister() {
		t.Errorf("the registry was NOT expected to contain a Register instruction")
		return
	}

	if !reg.IsUnregister() {
		t.Errorf("the registry was expected to contain an Unregister instruction")
		return
	}

	if reg.IsFetch() {
		t.Errorf("the registry was NOT expected to contain a Fetchregistry instruction")
		return
	}
}

func Test_registry_withFetchRegistry_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("registry", grammarFile)

	file := "./tests/codes/registry/fetchregistry.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	reg := ins.(Registry)
	if reg.IsRegister() {
		t.Errorf("the registry was NOT expected to contain a Register instruction")
		return
	}

	if reg.IsUnregister() {
		t.Errorf("the registry was NOT expected to contain an Unregister instruction")
		return
	}

	if !reg.IsFetch() {
		t.Errorf("the registry was expected to contain a Fetchregistry instruction")
		return
	}
}
