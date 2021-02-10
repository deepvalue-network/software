package disks

import (
	"github.com/deepvalue-network/software/blockchain/domain/blocks"
	"github.com/deepvalue-network/software/blockchain/domain/links"
	"github.com/deepvalue-network/software/libs/files/domain/files"
)

type serviceLink struct {
	blockService blocks.Service
	fileService  files.Service
}

func createServiceLink(
	blockService blocks.Service,
	fileService files.Service,
) links.Service {
	out := serviceLink{
		blockService: blockService,
		fileService:  fileService,
	}

	return &out
}

// Insert inserts a link
func (app *serviceLink) Insert(link links.Link) error {
	block := link.NextBlock()
	err := app.blockService.Insert(block)
	if err != nil {
		return err
	}

	return app.fileService.Insert(link.Hash().String(), link)
}

// Delete deletes a link
func (app *serviceLink) Delete(link links.Link) error {
	return app.fileService.Delete(link.Hash().String())
}
