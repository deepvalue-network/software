package disks

import (
	"github.com/deepvalue-network/software/blockchain/domain/links"
	link_mined "github.com/deepvalue-network/software/blockchain/domain/links/mined"
	"github.com/deepvalue-network/software/libs/files/domain/files"
)

type serviceLinkMined struct {
	linkService links.Service
	fileService files.Service
}

func createServiceLinkMined(
	linkService links.Service,
	fileService files.Service,
) link_mined.Service {
	out := serviceLinkMined{
		linkService: linkService,
		fileService: fileService,
	}

	return &out
}

// Insert inserts a mined link
func (app *serviceLinkMined) Insert(minedLink link_mined.Link) error {
	link := minedLink.Link()
	err := app.linkService.Insert(link)
	if err != nil {
		return err
	}

	return app.fileService.Insert(minedLink.Hash().String(), minedLink)
}

// Delete deletes a mined link
func (app *serviceLinkMined) Delete(minedLink link_mined.Link) error {
	return app.fileService.Delete(minedLink.Hash().String())
}
