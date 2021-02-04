package disks

import (
	"errors"
	"fmt"

	blocks_mined "github.com/steve-care-software/products/blockchain/domain/blocks/mined"
	"github.com/steve-care-software/products/libs/files/domain/files"
	"github.com/steve-care-software/products/libs/hash"
)

type repositoryBlockMined struct {
	fileRepository files.Repository
}

func createRepositoryBlockMined(
	fileRepository files.Repository,
) blocks_mined.Repository {
	out := repositoryBlockMined{
		fileRepository: fileRepository,
	}

	return &out
}

// List lists the hashes of head blocks
func (app *repositoryBlockMined) List() ([]hash.Hash, error) {
	return app.fileRepository.List()
}

// Retrieve retrieves a block by hash
func (app *repositoryBlockMined) Retrieve(blockHash hash.Hash) (blocks_mined.Block, error) {
	dehydrated, err := app.fileRepository.Retrieve(blockHash.String())
	if err != nil {
		return nil, err
	}

	if ins, ok := dehydrated.(blocks_mined.Block); ok {
		return ins, nil
	}

	str := fmt.Sprintf("the mined block (head hash: %s) could not be dehydrated into a mined block instance", blockHash.String())
	return nil, errors.New(str)
}
