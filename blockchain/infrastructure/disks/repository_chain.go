package disks

import (
	"errors"
	"fmt"

	"github.com/deepvalue-network/software/blockchain/domain/chains"
	"github.com/deepvalue-network/software/libs/files/domain/files"
	uuid "github.com/satori/go.uuid"
)

type repositoryChain struct {
	fileRepository files.Repository
}

func createRepositoryChain(
	fileRepository files.Repository,
) chains.Repository {
	out := repositoryChain{
		fileRepository: fileRepository,
	}

	return &out
}

// List lists the ids of the chain
func (app *repositoryChain) List() ([]*uuid.UUID, error) {
	return app.fileRepository.ListIDs()
}

// Retrieve retrieves the chain by id
func (app *repositoryChain) Retrieve(chainID *uuid.UUID) (chains.Chain, error) {
	dehydrated, err := app.fileRepository.Retrieve(chainID.String())
	if err != nil {
		return nil, err
	}

	if ins, ok := dehydrated.(chains.Chain); ok {
		return ins, nil
	}

	str := fmt.Sprintf("the chain ( ID: %s) could not be dehydrated into a chain instance", chainID.String())
	return nil, errors.New(str)
}
