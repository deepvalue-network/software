package repositories

import (
	mined_link "github.com/steve-care-software/products/blockchain/domain/links/mined"
	"github.com/steve-care-software/products/libs/hash"
)

type minedLink struct {
	minedLinkRepository mined_link.Repository
}

func createMinedLink(
	minedLinkRepository mined_link.Repository,
) MinedLink {
	out := minedLink{
		minedLinkRepository: minedLinkRepository,
	}

	return &out
}

// List returns the list of mined link hashes
func (app *minedLink) List() ([]hash.Hash, error) {
	return app.minedLinkRepository.List()
}

// Head returns the head mined link
func (app *minedLink) Head() (mined_link.Link, error) {
	return app.minedLinkRepository.Head()
}

// Retrieve retrieves a mined link by hash
func (app *minedLink) Retrieve(hash hash.Hash) (mined_link.Link, error) {
	return app.minedLinkRepository.Retrieve(hash)
}
