package disks

import (
	"github.com/steve-care-software/products/blockchain/domain/blocks"
	"github.com/steve-care-software/products/libs/files/domain/files"
)

type serviceBlock struct {
	fileService files.Service
}

func createServiceBlock(
	fileService files.Service,
) blocks.Service {
	out := serviceBlock{
		fileService: fileService,
	}

	return &out
}

// Insert inserts a block
func (app *serviceBlock) Insert(block blocks.Block) error {
	return app.fileService.Insert(block.Tree().Head().String(), block)
}

// Delete deletes a block
func (app *serviceBlock) Delete(block blocks.Block) error {
	return app.fileService.Delete(block.Tree().Head().String())
}
