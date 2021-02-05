package repositories

import (
	mined_block "github.com/deepvalue-network/software/blockchain/domain/blocks/mined"
	"github.com/deepvalue-network/software/libs/hash"
)

type minedBlock struct {
	mineBlockRepository mined_block.Repository
}

func createMinedBlock(
	mineBlockRepository mined_block.Repository,
) MinedBlock {
	out := minedBlock{
		mineBlockRepository: mineBlockRepository,
	}

	return &out
}

// List returns the list of mined block hashes
func (app *minedBlock) List() ([]hash.Hash, error) {
	return app.mineBlockRepository.List()
}

// Retrieve retrieves a mined block by hash
func (app *minedBlock) Retrieve(hash hash.Hash) (mined_block.Block, error) {
	return app.mineBlockRepository.Retrieve(hash)
}
