package disks

import (
	"github.com/deepvalue-network/software/blockchain/domain/blocks"
	mined_blocks "github.com/deepvalue-network/software/blockchain/domain/blocks/mined"
	"github.com/deepvalue-network/software/libs/files/domain/files"
)

type serviceBlockMined struct {
	blockService blocks.Service
	fileService  files.Service
}

func createServiceBlockMined(
	blockService blocks.Service,
	fileService files.Service,
) mined_blocks.Service {
	out := serviceBlockMined{
		blockService: blockService,
		fileService:  fileService,
	}

	return &out
}

// Insert inserts a block
func (app *serviceBlockMined) Insert(block mined_blocks.Block) error {
	err := app.blockService.Insert(block.Block())
	if err != nil {
		return err
	}

	return app.fileService.Insert(block.Hash().String(), block)
}

// Delete deletes a block
func (app *serviceBlockMined) Delete(block mined_blocks.Block) error {
	return app.fileService.Delete(block.Hash().String())
}
