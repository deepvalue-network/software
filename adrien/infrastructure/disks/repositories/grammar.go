package repositories

import "github.com/deepvalue-network/software/adrien/domain/grammars"

type grammar struct {
}

func createGrammar() grammars.Repository {
	out := grammar{}
	return &out
}

// Retrieve retrieves a grammar from path, from disk
func (app *grammar) Retrieve(relativePath string) (grammars.Grammar, error) {
	return nil, nil
}
