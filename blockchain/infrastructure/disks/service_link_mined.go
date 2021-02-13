package disks

import (
	"github.com/deepvalue-network/software/blockchain/domain/links"
	link_mined "github.com/deepvalue-network/software/blockchain/domain/links/mined"
	"github.com/deepvalue-network/software/libs/events"
	"github.com/deepvalue-network/software/libs/files/domain/files"
)

type serviceLinkMined struct {
	eventManager           events.Manager
	minedLinkRepository    link_mined.Repository
	linkService            links.Service
	fileService            files.Service
	linkPointerFileService files.Service
	headPointerFileService files.Service
	headFileName           string
}

func createServiceLinkMined(
	eventManager events.Manager,
	minedLinkRepository link_mined.Repository,
	linkService links.Service,
	fileService files.Service,
	linkPointerFileService files.Service,
	headPointerFileService files.Service,
	headFileName string,
) link_mined.Service {
	out := serviceLinkMined{
		eventManager:           eventManager,
		minedLinkRepository:    minedLinkRepository,
		linkService:            linkService,
		fileService:            fileService,
		linkPointerFileService: linkPointerFileService,
		headPointerFileService: headPointerFileService,
		headFileName:           headFileName,
	}

	return &out
}

// Insert inserts a mined link
func (app *serviceLinkMined) Insert(minedLink link_mined.Link) error {
	return app.eventManager.Trigger(EventLinkMinedInsert, minedLink, func() error {
		link := minedLink.Link()
		err := app.linkService.Insert(link)
		if err != nil {
			return err
		}

		minedLinkHashStr := minedLink.Hash().String()
		err = app.fileService.Insert(minedLinkHashStr, minedLink)
		if err != nil {
			return err
		}

		// save the pointers:
		err = app.linkPointerFileService.Insert(link.Hash().String(), minedLinkHashStr)
		if err != nil {
			return err
		}

		err = app.headPointerFileService.Insert(app.headFileName, minedLinkHashStr)
		if err != nil {
			return err
		}

		return nil
	})
}

// Delete deletes a mined link
func (app *serviceLinkMined) Delete(minedLink link_mined.Link) error {
	return app.eventManager.Trigger(EventLinkMinedDelete, minedLink, func() error {
		return app.fileService.Delete(minedLink.Hash().String())
	})
}

// DeleteByLink deletes a mined link by link
func (app *serviceLinkMined) DeleteByLink(link links.Link) error {
	minedLink, err := app.minedLinkRepository.RetrieveByLinkHash(link.Hash())
	if err != nil {
		return err
	}

	return app.Delete(minedLink)
}
