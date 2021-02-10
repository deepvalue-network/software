package disks

import (
	"errors"
	"fmt"

	"github.com/deepvalue-network/software/blockchain/domain/links"
	"github.com/deepvalue-network/software/libs/files/domain/files"
	"github.com/deepvalue-network/software/libs/hash"
)

type repositoryLink struct {
	fileRepository files.Repository
}

func createRepositoryLink(
	fileRepository files.Repository,
) links.Repository {
	out := repositoryLink{
		fileRepository: fileRepository,
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
