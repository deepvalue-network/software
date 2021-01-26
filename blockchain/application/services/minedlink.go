package services

import (
	"github.com/steve-care-software/products/blockchain/application/repositories"
	mined_link "github.com/steve-care-software/products/blockchain/domain/links/mined"
	"github.com/steve-care-software/products/libs/hash"
)

type minedLink struct {
	minedLinkService    mined_link.Service
	minedLinkBuilder    mined_link.Builder
	linkRepository      repositories.Link
	minedLinkRepository repositories.MinedLink
	minerApp            Miner
}

func createMinedLink(
	minedLinkService mined_link.Service,
	minedLinkBuilder mined_link.Builder,
	linkRepository repositories.Link,
	minedLinkRepository repositories.MinedLink,
	minerApp Miner,
) MinedLink {
	out := minedLink{
		minedLinkService:    minedLinkService,
		minedLinkBuilder:    minedLinkBuilder,
		linkRepository:      linkRepository,
		minedLinkRepository: minedLinkRepository,
		minerApp:            minerApp,
	}

	return &out
}

// Mine mines a link by hash
func (app *minedLink) Mine(miningValue uint8, difficulty uint, linkHash hash.Hash) (mined_link.Link, error) {
	link, err := app.linkRepository.Retrieve(linkHash)
	if err != nil {
		return nil, err
	}

	hash := link.Hash()
	results, _, err := app.minerApp.Mine(miningValue, difficulty, hash)
	if err != nil {
		return nil, err
	}

	minedLink, err := app.minedLinkBuilder.Create().WithLink(link).WithResults(results).Now()
	if err != nil {
		return nil, err
	}

	err = app.minedLinkService.Insert(minedLink)
	if err != nil {
		return nil, err
	}

	return minedLink, nil
}

// MineList mines all the remaining link that needs to be mined
func (app *minedLink) MineList(miningValue uint8, difficulty uint) ([]mined_link.Link, error) {
	hashes, err := app.linkRepository.List()
	if err != nil {
		return nil, err
	}

	out := []mined_link.Link{}
	for _, oneHash := range hashes {
		minedLink, err := app.Mine(miningValue, difficulty, oneHash)
		if err != nil {
			return nil, err
		}

		out = append(out, minedLink)
	}

	return out, nil
}

// Delete deletes a mined link by hash
func (app *minedLink) Delete(hash hash.Hash) error {
	link, err := app.minedLinkRepository.Retrieve(hash)
	if err != nil {
		return err
	}

	return app.minedLinkService.Delete(link)
}
