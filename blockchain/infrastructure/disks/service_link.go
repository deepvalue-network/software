package disks

import (
	"github.com/deepvalue-network/software/blockchain/domain/blocks"
	"github.com/deepvalue-network/software/blockchain/domain/links"
	"github.com/deepvalue-network/software/libs/events"
	"github.com/deepvalue-network/software/libs/files/domain/files"
	"github.com/deepvalue-network/software/libs/hash"
)

type serviceLink struct {
	eventManager                events.Manager
	linkRepository              links.Repository
	blockService                blocks.Service
	fileService                 files.Service
	blockPointerFileService     files.Service
	minedLinkPointerFileService files.Service
}

func createServiceLink(
	eventManager events.Manager,
	linkRepository links.Repository,
	blockService blocks.Service,
	fileService files.Service,
	blockPointerFileService files.Service,
	minedLinkPointerFileService files.Service,
) links.Service {
	out := serviceLink{
		eventManager:                eventManager,
		linkRepository:              linkRepository,
		blockService:                blockService,
		fileService:                 fileService,
		blockPointerFileService:     blockPointerFileService,
		minedLinkPointerFileService: minedLinkPointerFileService,
	}

	return &out
}

// Insert inserts a link
func (app *serviceLink) Insert(link links.Link) error {
	return app.eventManager.Trigger(EventLinkInsert, link, func() error {
		block := link.NextBlock()
		err := app.blockService.Insert(block)
		if err != nil {
			return err
		}

		// save the link:
		linkHashStr := link.Hash().String()
		err = app.fileService.Insert(linkHashStr, link)
		if err != nil {
			return err
		}

		// save the pointers:
		err = app.blockPointerFileService.Insert(block.Tree().Head().String(), linkHashStr)
		if err != nil {
			return err
		}

		err = app.minedLinkPointerFileService.Insert(link.PrevMinedLink().String(), linkHashStr)
		if err != nil {
			return err
		}

		return nil
	})
}

// Delete deletes a link
func (app *serviceLink) Delete(link links.Link) error {
	return app.eventManager.Trigger(EventLinkDelete, link, func() error {
		return app.fileService.Delete(link.Hash().String())
	})
}

// DeleteByBlock deletes a link by block
func (app *serviceLink) DeleteByBlock(block blocks.Block) error {
	link, err := app.linkRepository.RetrieveByBlockHash(block.Tree().Head())
	if err != nil {
		return err
	}

	return app.Delete(link)
}

// DeleteByMinedLinkHash deletes a link by mined link hash
func (app *serviceLink) DeleteByMinedLinkHash(minedLinkHash hash.Hash) error {
	link, err := app.linkRepository.RetrieveByMinedLinkHash(minedLinkHash)
	if err != nil {
		return err
	}

	return app.Delete(link)
}
