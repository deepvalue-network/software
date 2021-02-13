package disks

import (
	"errors"
	"fmt"

	"github.com/deepvalue-network/software/blockchain/domain/links"
	"github.com/deepvalue-network/software/libs/files/domain/files"
	"github.com/deepvalue-network/software/libs/hash"
)

type repositoryLink struct {
	hashAdapter                    hash.Adapter
	fileRepository                 files.Repository
	blockPointerFileRepository     files.Repository
	minedLinkPointerFileRepository files.Repository
}

func createRepositoryLink(
	hashAdapter hash.Adapter,
	fileRepository files.Repository,
	blockPointerFileRepository files.Repository,
	minedLinkPointerFileRepository files.Repository,
) links.Repository {
	out := repositoryLink{
		hashAdapter:                    hashAdapter,
		fileRepository:                 fileRepository,
		blockPointerFileRepository:     blockPointerFileRepository,
		minedLinkPointerFileRepository: minedLinkPointerFileRepository,
	}

	return &out
}

// List returns the list
func (app *repositoryLink) List() ([]hash.Hash, error) {
	return app.fileRepository.List()
}

// Retrieve retrieves the link by hash
func (app *repositoryLink) Retrieve(linkHash hash.Hash) (links.Link, error) {
	dehydrated, err := app.fileRepository.Retrieve(linkHash.String())
	if err != nil {
		return nil, err
	}

	if ins, ok := dehydrated.(links.Link); ok {
		return ins, nil
	}

	str := fmt.Sprintf("the link ( hash: %s) could not be dehydrated into a link instance", linkHash.String())
	return nil, errors.New(str)
}

// RetrieveByBlockHash retrieves a link by block hash
func (app *repositoryLink) RetrieveByBlockHash(blockHash hash.Hash) (links.Link, error) {
	ptrData, err := app.blockPointerFileRepository.Retrieve(blockHash.String())
	if err != nil {
		return nil, err
	}

	linkHash, err := app.hashAdapter.FromString(string(ptrData.([]byte)))
	if err != nil {
		return nil, err
	}

	return app.Retrieve(*linkHash)
}

// RetrieveByMinedLinkHash retrieves a link by mined link hash
func (app *repositoryLink) RetrieveByMinedLinkHash(minedLinkHash hash.Hash) (links.Link, error) {
	ptrData, err := app.minedLinkPointerFileRepository.Retrieve(minedLinkHash.String())
	if err != nil {
		return nil, err
	}

	linkHash, err := app.hashAdapter.FromString(string(ptrData.([]byte)))
	if err != nil {
		return nil, err
	}

	return app.Retrieve(*linkHash)
}
