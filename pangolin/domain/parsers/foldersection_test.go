package parsers

import (
	"testing"
)

func Test_folderSection_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("folderSection", grammarFile)

	file := "./tests/codes/foldersection/middle.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	section := ins.(FolderSection)
	if section.IsTail() {
		t.Errorf("the FolderSection was NOT expecting to be a tail")
		return
	}

	if section.Name().Name() != "myFolder" {
		t.Errorf("the name was expected to be %s, %s returned", "myFolder", section.Name().Name())
		return
	}
}

func Test_folderSection_isTail_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("folderSection", grammarFile)

	file := "./tests/codes/foldersection/tail.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	section := ins.(FolderSection)
	if !section.IsTail() {
		t.Errorf("the FolderSection was expecting to be a tail")
		return
	}

	if section.Name().Name() != "myTail" {
		t.Errorf("the name was expected to be %s, %s returned", "myTail", section.Name().Name())
		return
	}
}
