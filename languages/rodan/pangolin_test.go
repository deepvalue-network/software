package rodan

import (
	"errors"
	"path/filepath"
	"testing"

	"github.com/deepvalue-network/software/pangolin/bundles"
)

func TestPangolin_executeTests_Success(t *testing.T) {
	name := "pangolin"
	grammarFile := "../../pangolin/domain/parsers/grammar/grammar.json"
	langPath := "language.pangolin"
	dirPath, err := filepath.Abs(".")
	if err != nil {
		panic(err)
	}

	interpreter, err := bundles.NewInterpreter(dirPath, langPath, grammarFile, name)
	if err != nil {
		panic(err)
	}

	if !interpreter.IsLanguage() {
		panic(errors.New("the interpreter was expected to be a language interpreter"))
	}

	err = interpreter.Language().TestsAll()
	if err != nil {
		panic(err)
	}
}
