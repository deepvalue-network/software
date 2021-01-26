package rodan

import (
	"errors"
	"path/filepath"
	"testing"
)

func TestPangolin_executeTests_Success(t *testing.T) {
	// generate then save language:
	langPath := "language.pangolin"
	dirPath, err := filepath.Abs(".")
	if err != nil {
		panic(err)
	}

	parserLinker, interpreterBuilder, err := create(dirPath)
	if err != nil {
		panic(err)
	}

	program, err := parserLinker.File(langPath)
	if err != nil {
		panic(err)
	}

	interpreter, err := interpreterBuilder.Create().WithProgram(program).Now()
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
