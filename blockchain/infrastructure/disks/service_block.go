package disks

import (
	"github.com/deepvalue-network/software/blockchain/domain/blocks"
	"github.com/deepvalue-network/software/libs/events"
	"github.com/deepvalue-network/software/libs/files/domain/files"
)

type serviceBlock struct {
	eventManager events.Manager
	fileService  files.Service
}

func createServiceBlock(
	eventManager events.Manager,
	fileService files.Service,
) blocks.Service {
	out := serviceBlock{
		eventManager: eventManager,
		fileService:  fileService,
	}

	return &out
}

// Insert inserts a block
func (app *serviceBlock) Insert(block blocks.Block) error {
	return app.eventManager.Trigger(EventBlockInsert, block, func() error {
		return app.fileService.Insert(block.Tree().Head().String(), block)
	})
}

// Delete deletes a block
func (app *serviceBlock) Delete(block blocks.Block) error {
	return app.eventManager.Trigger(EventBlockDelete, block, func() error {
		return app.fileService.Delete(block.Tree().Head().String())
	})
}
