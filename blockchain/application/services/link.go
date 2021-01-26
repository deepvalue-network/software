package services

import (
	"github.com/steve-care-software/products/blockchain/application/repositories"
	"github.com/steve-care-software/products/blockchain/domain/links"
	"github.com/steve-care-software/products/libs/hash"
)

type link struct {
	linkService         links.Service
	linkBuilder         links.Builder
	blockRepository     repositories.Block
	linkRepository      repositories.Link
	minedLinkRepository repositories.MinedLink
}

func createLink(
	linkService links.Service,
	linkBuilder links.Builder,
	blockRepository repositories.Block,
	linkRepository repositories.Link,
	minedLinkRepository repositories.MinedLink,
) Link {
	out := link{
		linkService:         linkService,
		linkBuilder:         linkBuilder,
		blockRepository:     blockRepository,
		linkRepository:      linkRepository,
		minedLinkRepository: minedLinkRepository,
	}

	return &out
}

// Create creates a link using the previous mined link and the next block
func (app *link) Create(prevMinedLinkHash hash.Hash, nextBlockHash hash.Hash) (links.Link, error) {
	prevMinedLink, err := app.minedLinkRepository.Retrieve(prevMinedLinkHash)
	if err != nil {
		return nil, err
	}

	nextBlock, err := app.blockRepository.Retrieve(nextBlockHash)
	if err != nil {
		return nil, err
	}

	hash := prevMinedLink.Hash()
	index := prevMinedLink.Link().Index() + 1
	link, err := app.linkBuilder.Create().WithIndex(index).WithPreviousMinedLink(hash).WithNextBlock(nextBlock).Now()
	if err != nil {
		return nil, err
	}

	err = app.linkService.Insert(link)
	if err != nil {
		return nil, err
	}

	return link, nil
}

// Delete deletes a link by hash
func (app *link) Delete(hash hash.Hash) error {
	link, err := app.linkRepository.Retrieve(hash)
	if err != nil {
		return err
	}

	return app.linkService.Delete(link)
}
