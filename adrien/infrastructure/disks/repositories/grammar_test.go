package repositories

import (
	"testing"
)

func TestGrammar_Success(t *testing.T) {
	basePath := "./test_files"
	grammarFileRepository := NewGrammarFile(basePath)
	grammarRepository := NewGrammar(grammarFileRepository)

	relPath := "./grammar.json"
	grammar, err := grammarRepository.Retrieve(relPath)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if grammar.HasChannels() {
		t.Errorf("the grammar was NOT expecting channels")
		return
	}
}

func TestGrammar_withChannels_Success(t *testing.T) {
	basePath := "./test_files"
	grammarFileRepository := NewGrammarFile(basePath)
	grammarRepository := NewGrammar(grammarFileRepository)

	relPath := "./grammar_with_channels.json"
	grammar, err := grammarRepository.Retrieve(relPath)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !grammar.HasChannels() {
		t.Errorf("the grammar was expecting channels")
		return
	}
}
