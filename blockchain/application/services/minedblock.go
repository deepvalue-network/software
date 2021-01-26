package services

import (
	"github.com/steve-care-software/products/blockchain/application/repositories"
	mined_block "github.com/steve-care-software/products/blockchain/domain/blocks/mined"
	"github.com/steve-care-software/products/libs/hash"
)

type minedBlock struct {
	mineBlockBuilder    mined_block.Builder
	mineBlockService    mined_block.Service
	mineBlockRepository repositories.MinedBlock
	blockRepository     repositories.Block
	minerApp            Miner
}

func createMinedBlock(
	mineBlockBuilder mined_block.Builder,
	mineBlockService mined_block.Service,
	mineBlockRepository repositories.MinedBlock,
	blockRepository repositories.Block,
	minerApp Miner,
) MinedBlock {
	out := minedBlock{
		mineBlockBuilder:    mineBlockBuilder,
		mineBlockService:    mineBlockService,
		mineBlockRepository: mineBlockRepository,
		blockRepository:     blockRepository,
		minerApp:            minerApp,
	}

	return &out
}

// Mine mines a block by block hash
func (app *minedBlock) Mine(miningValue uint8, baseDifficulty uint, incrPerHash float64, blockHash hash.Hash) (mined_block.Block, error) {
	// retrieve the block:
	block, err := app.blockRepository.Retrieve(blockHash)
	if err != nil {
		return nil, err
	}

	// calculate the difficulty:
	hashes := block.Hashes()
	difficulty := app.calculateDifficulty(baseDifficulty, incrPerHash, len(hashes))

	// mine the block:
	hash := block.Tree().Head()
	results, _, err := app.minerApp.Mine(miningValue, difficulty, hash)
	if err != nil {
		return nil, err
	}

	minedBlock, err := app.mineBlockBuilder.Create().WithBlock(block).WithResults(results).Now()
	if err != nil {
		return nil, err
	}

	err = app.mineBlockService.Insert(minedBlock)
	if err != nil {
		return nil, err
	}

	return minedBlock, nil
}

// MineList mine the list of block that needs to be mined
func (app *minedBlock) MineList(miningValue uint8, baseDifficulty uint, incrPerHash float64) ([]mined_block.Block, error) {
	blockHashes, err := app.blockRepository.List()
	if err != nil {
		return nil, err
	}

	out := []mined_block.Block{}
	for _, oneBlockHash := range blockHashes {
		minedBlock, err := app.Mine(miningValue, baseDifficulty, incrPerHash, oneBlockHash)
		if err != nil {
			return nil, err
		}

		out = append(out, minedBlock)
	}

	return out, nil
}

// Delete deletes a mined block by hash
func (app *minedBlock) Delete(hash hash.Hash) error {
	minedBlock, err := app.mineBlockRepository.Retrieve(hash)
	if err != nil {
		return err
	}

	return app.mineBlockService.Delete(minedBlock)
}

func (app *minedBlock) calculateDifficulty(baseDifficulty uint, incrPerHash float64, amountHashes int) uint {
	sum := float64(0)
	base := float64(baseDifficulty)
	for i := 0; i < int(amountHashes); i++ {
		sum += incrPerHash
	}

	return uint(sum + base)
}
