package repositories

import (
	"github.com/steve-care-software/products/blockchain/domain/links"
	"github.com/steve-care-software/products/libs/hash"
)

type link struct {
	linkRepository links.Repository
}

func createLink(
	linkRepository links.Repository,
) Link {
	out := link{
		linkRepository: linkRepository,
	}

	return &out
}

// List returns the list of hashes
func (app *link) List() ([]hash.Hash, error) {
	return app.linkRepository.List()
}

// Retrieve retrieves a link by hash
func (app *link) Retrieve(hash hash.Hash) (links.Link, error) {
	return app.linkRepository.Retrieve(hash)
}
