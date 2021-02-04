package disks

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/products/blockchain/domain/blocks"
	"github.com/steve-care-software/products/libs/files/domain/files"
	"github.com/steve-care-software/products/libs/hash"
)

type repositoryBlock struct {
	fileRepository files.Repository
}

func createRepositoryBlock(
	fileRepository files.Repository,
) blocks.Repository {
	out := repositoryBlock{
		fileRepository: fileRepository,
	}

	return &out
}

// List lists the hashes of head blocks
func (app *repositoryBlock) List() ([]hash.Hash, error) {
	return app.fileRepository.List()
}

// Retrieve retrieves a block by hash
func (app *repositoryBlock) Retrieve(blockHash hash.Hash) (blocks.Block, error) {
	dehydrated, err := app.fileRepository.Retrieve(blockHash.String())
	if err != nil {
		return nil, err
	}

	if ins, ok := dehydrated.(blocks.Block); ok {
		return ins, nil
	}

	str := fmt.Sprintf("the block (head hash: %s) could not be dehydrated into a block instance", blockHash.String())
	return nil, errors.New(str)
}
