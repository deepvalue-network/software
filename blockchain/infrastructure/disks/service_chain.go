package disks

import (
	mined_blocks "github.com/deepvalue-network/software/blockchain/domain/blocks/mined"
	"github.com/deepvalue-network/software/blockchain/domain/chains"
	mined_links "github.com/deepvalue-network/software/blockchain/domain/links/mined"
	"github.com/deepvalue-network/software/libs/files/domain/files"
)

type serviceChain struct {
	validator    chains.Validator
	blockService mined_blocks.Service
	linkService  mined_links.Service
	fileService  files.Service
}

func createServiceChain(
	validator chains.Validator,
	blockService mined_blocks.Service,
	linkService mined_links.Service,
	fileService files.Service,
) chains.Service {
	out := serviceChain{
		validator:    validator,
		blockService: blockService,
		linkService:  linkService,
		fileService:  fileService,
	}

	return &out
}

// Insert inserts a chain
func (app *serviceChain) Insert(chain chains.Chain) error {
	return nil
}

// Update updates a chain
func (app *serviceChain) Update(original chains.Chain, updated chains.Chain) error {
	return nil
}

// Delete deletes a chain
func (app *serviceChain) Delete(chain chains.Chain) error {
	return nil
}
