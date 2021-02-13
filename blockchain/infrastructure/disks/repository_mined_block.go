package disks

import (
	"errors"
	"fmt"

	blocks_mined "github.com/deepvalue-network/software/blockchain/domain/blocks/mined"
	"github.com/deepvalue-network/software/libs/files/domain/files"
	"github.com/deepvalue-network/software/libs/hash"
)

type repositoryBlockMined struct {
	hashAdapter           hash.Adapter
	fileRepository        files.Repository
	pointerFileRepository files.Repository
}

func createRepositoryBlockMined(
	hashAdapter hash.Adapter,
	fileRepository files.Repository,
	pointerFileRepository files.Repository,
) blocks_mined.Repository {
	out := repositoryBlockMined{
		hashAdapter:           hashAdapter,
		fileRepository:        fileRepository,
		pointerFileRepository: pointerFileRepository,
	}

	return &out
}

// List lists the hashes of head blocks
func (app *repositoryBlockMined) List() ([]hash.Hash, error) {
	return app.fileRepository.List()
}

// Retrieve retrieves a mined block by hash
func (app *repositoryBlockMined) Retrieve(minedBlockHash hash.Hash) (blocks_mined.Block, error) {
	dehydrated, err := app.fileRepository.Retrieve(minedBlockHash.String())
	if err != nil {
		return nil, err
	}

	if ins, ok := dehydrated.(blocks_mined.Block); ok {
		return ins, nil
	}

	str := fmt.Sprintf("the mined block (head hash: %s) could not be dehydrated into a mined block instance", minedBlockHash.String())
	return nil, errors.New(str)
}

// RetrieveByBlockHash retrieves a mined block by block hash
func (app *repositoryBlockMined) RetrieveByBlockHash(blockHash hash.Hash) (blocks_mined.Block, error) {
	ptrData, err := app.pointerFileRepository.Retrieve(blockHash.String())
	if err != nil {
		return nil, err
	}

	minedBlockHash, err := app.hashAdapter.FromString(string(ptrData.([]byte)))
	if err != nil {
		return nil, err
	}

	return app.Retrieve(*minedBlockHash)
}
