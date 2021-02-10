package disks

import (
	"errors"
	"fmt"

	link_mined "github.com/deepvalue-network/software/blockchain/domain/links/mined"
	"github.com/deepvalue-network/software/libs/files/domain/files"
	"github.com/deepvalue-network/software/libs/hash"
)

type repositoryLinkMined struct {
	fileRepository files.Repository
}

func createRepositoryLinkMined(
	fileRepository files.Repository,
) link_mined.Repository {
	out := repositoryLinkMined{
		fileRepository: fileRepository,
	}

	return &out
}

// Head returns the head link
func (app *repositoryLinkMined) Head() (link_mined.Link, error) {
	return nil, nil
}

// List returns the list of mined links
func (app *repositoryLinkMined) List() ([]hash.Hash, error) {
	return app.fileRepository.List()
}

// Retrieve retrieves a mined link by hash
func (app *repositoryLinkMined) Retrieve(linkHash hash.Hash) (link_mined.Link, error) {
	dehydrated, err := app.fileRepository.Retrieve(linkHash.String())
	if err != nil {
		return nil, err
	}

	if ins, ok := dehydrated.(link_mined.Link); ok {
		return ins, nil
	}

	str := fmt.Sprintf("the mined link ( hash: %s) could not be dehydrated into a mined link instance", linkHash.String())
	return nil, errors.New(str)
}
