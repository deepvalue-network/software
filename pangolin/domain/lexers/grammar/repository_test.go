package grammar

import (
	"log"
	"testing"
)

func TestRetrieve_Success(t *testing.T) {
	// create the repository:
	repository, _ := NewRepositoryBuilder().Create().Now()

	// read the grammar:
	grammar, err := repository.RetrieveFromFile("", "grammar", "tests/valid/grammar.json")
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	if grammar == nil {
		t.Errorf("the grammar was expected to be valid, nil returned")
		return
	}

	log.Printf("\n %v \n", grammar)
}
