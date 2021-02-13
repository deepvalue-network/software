package disks

import (
	"github.com/deepvalue-network/software/blockchain/domain/blocks"
	mined_blocks "github.com/deepvalue-network/software/blockchain/domain/blocks/mined"
	"github.com/deepvalue-network/software/libs/events"
	"github.com/deepvalue-network/software/libs/files/domain/files"
)

type serviceBlockMined struct {
	eventManager         events.Manager
	minedBlockRepository mined_blocks.Repository
	blockService         blocks.Service
	fileService          files.Service
	pointerFileService   files.Service
}

func createServiceBlockMined(
	eventManager events.Manager,
	minedBlockRepository mined_blocks.Repository,
	blockService blocks.Service,
	fileService files.Service,
	pointerFileService files.Service,
) mined_blocks.Service {
	out := serviceBlockMined{
		eventManager:         eventManager,
		minedBlockRepository: minedBlockRepository,
		blockService:         blockService,
		fileService:          fileService,
		pointerFileService:   pointerFileService,
	}

	return &out
}

// Insert inserts a block
func (app *serviceBlockMined) Insert(minedBlock mined_blocks.Block) error {
	return app.eventManager.Trigger(EventBlockMinedInsert, minedBlock, func() error {
		err := app.blockService.Insert(minedBlock.Block())
		if err != nil {
			return err
		}

		// save the mined block:
		minedBlockHashStr := minedBlock.Hash().String()
		err = app.fileService.Insert(minedBlockHashStr, minedBlock)
		if err != nil {
			return err
		}

		// save the pointers:
		err = app.pointerFileService.Insert(minedBlock.Block().Tree().Head().String(), minedBlockHashStr)
		if err != nil {
			return err
		}

		return nil
	})
}

// Delete deletes a block
func (app *serviceBlockMined) Delete(minedBlock mined_blocks.Block) error {
	return app.eventManager.Trigger(EventBlockMinedDelete, minedBlock, func() error {
		return app.fileService.Delete(minedBlock.Hash().String())
	})
}

// DeleteByBlock deletes a mined block by underlying block
func (app *serviceBlockMined) DeleteByBlock(block blocks.Block) error {
	minedBlock, err := app.minedBlockRepository.RetrieveByBlockHash(block.Tree().Head())
	if err != nil {
		return err
	}

	return app.Delete(minedBlock)
}
