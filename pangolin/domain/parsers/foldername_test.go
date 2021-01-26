package parsers

import (
	"testing"
)

func Test_folderName_withCurrent_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("folderName", grammarFile)

	file := "./tests/codes/foldername/current.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	folder := ins.(FolderName)
	if !folder.IsCurrent() {
		t.Errorf("the folderName was expecting to be a current")
		return
	}

	if folder.IsPrevious() {
		t.Errorf("the folderName was NOT expecting to be a previous")
		return
	}

	if folder.IsName() {
		t.Errorf("the folderName was NOT expecting to be a name")
		return
	}
}

func Test_folderName_withPrevious_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("folderName", grammarFile)

	file := "./tests/codes/foldername/previous.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	folder := ins.(FolderName)
	if folder.IsCurrent() {
		t.Errorf("the folderName was NOT expecting to be a current")
		return
	}

	if !folder.IsPrevious() {
		t.Errorf("the folderName was expecting to be a previous")
		return
	}

	if folder.IsName() {
		t.Errorf("the folderName was NOT expecting to be a name")
		return
	}
}

func Test_folderName_withName_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("folderName", grammarFile)

	file := "./tests/codes/foldername/name.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	folder := ins.(FolderName)
	if folder.IsCurrent() {
		t.Errorf("the folderName was NOT expecting to be a current")
		return
	}

	if folder.IsPrevious() {
		t.Errorf("the folderName was NOT expecting to be a previous")
		return
	}

	if !folder.IsName() {
		t.Errorf("the folderName was expecting to be a name")
		return
	}

	if folder.Name() != "myFolderName" {
		t.Errorf("the name was expected to be %s, %s returned", "myFolderName", folder.Name())
		return
	}
}
