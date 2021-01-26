package repositories

import (
	uuid "github.com/satori/go.uuid"
	"github.com/steve-care-software/products/blockchain/domain/chains"
)

type chain struct {
	chainRepository chains.Repository
}

func createChain(
	chainRepository chains.Repository,
) Chain {
	out := chain{
		chainRepository: chainRepository,
	}
	return &out
}

// List list the chain ids
func (app *chain) List() ([]*uuid.UUID, error) {
	return app.chainRepository.List()
}

// List list the chain ids
func (app *chain) Retrieve(id *uuid.UUID) (chains.Chain, error) {
	return app.chainRepository.Retrieve(id)
}
