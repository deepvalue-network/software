package services

import (
	"github.com/deepvalue-network/software/blockchain/application/repositories"
	"github.com/deepvalue-network/software/blockchain/domain/blocks"
	"github.com/deepvalue-network/software/libs/hash"
)

type block struct {
	blockBuilder    blocks.Builder
	blockRepository repositories.Block
	blockService    blocks.Service
}

func createBlock(
	blockBuilder blocks.Builder,
	blockRepository repositories.Block,
	blockService blocks.Service,
) Block {
	out := block{
		blockBuilder:    blockBuilder,
		blockRepository: blockRepository,
		blockService:    blockService,
	}

	return &out
}

// Create creates a block from hashes
func (app *block) Create(hashes []hash.Hash) (blocks.Block, error) {
	block, err := app.blockBuilder.Create().WithHashes(hashes).Now()
	if err != nil {
		return nil, err
	}

	err = app.blockService.Insert(block)
	if err != nil {
		return nil, err
	}

	return block, nil
}

// Delete deletes a block by hash
func (app *block) Delete(hash hash.Hash) error {
	block, err := app.blockRepository.Retrieve(hash)
	if err != nil {
		return err
	}

	return app.blockService.Delete(block)
}
