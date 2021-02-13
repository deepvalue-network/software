package chains

import (
	"errors"
	"fmt"

	mined_block "github.com/deepvalue-network/software/blockchain/domain/blocks/mined"
	mined_link "github.com/deepvalue-network/software/blockchain/domain/links/mined"
)

type validator struct {
	minedBlockValidator mined_block.Validator
	minedLinkValidator  mined_link.Validator
	chainRepository     Repository
}

func createValidator(
	minedBlockValidator mined_block.Validator,
	minedLinkValidator mined_link.Validator,
	chainRepository Repository,
) Validator {
	out := validator{
		minedBlockValidator: minedBlockValidator,
		minedLinkValidator:  minedLinkValidator,
		chainRepository:     chainRepository,
	}

	return &out
}

// Validate validates a chain, returns an error if invalid, nil if valid
func (app *validator) Execute(chain Chain) error {
	genesis := chain.Genesis()
	rootMinedBlock := chain.Root()

	// retrieve the current chain of the same ID:
	chainID := chain.ID()
	retChain, err := app.chainRepository.Retrieve(chainID)
	if err != nil {
		// the chain does not exists and therefore is validated:
		return nil
	}

	if chain.Genesis().Hash().Compare(retChain.Genesis().Hash()) {
		str := fmt.Sprintf(
			"the given chain (ID: %s) contains a Genesis (hash: %s) that was updated from its previously stored version (genesis hash: %s)",
			chainID.String(),
			chain.Genesis().Hash().String(),
			retChain.Genesis().Hash(),
		)

		return errors.New(str)
	}

	// validate the root block:
	err = app.minedBlockValidator.Execute(genesis, rootMinedBlock)
	if err != nil {
		return err
	}

	// fetch the counted totalHashes and height:
	rootBlock := rootMinedBlock.Block()
	countedTotalHashes := uint(len(rootBlock.Hashes()))
	countedHeight := uint(1)

	// validate the mined link:
	if chain.HasHead() {
		head := chain.Head()
		linkTotalHashes, linkHeight, err := app.minedLinkValidator.Execute(genesis, head, rootBlock)
		if err != nil {
			return err
		}

		// add the link amounts to the counted amounts:
		countedTotalHashes += linkTotalHashes
		countedHeight += linkHeight
	}

	expectedTotalHashes := chain.TotalHashes()
	if countedTotalHashes != expectedTotalHashes {
		str := fmt.Sprintf("the chain (ID: %s) was expecting %d total hashes, %d were calculated on verification", chainID.String(), expectedTotalHashes, countedTotalHashes)
		return errors.New(str)
	}

	expectedHeight := chain.Height()
	if countedHeight != expectedHeight {
		str := fmt.Sprintf("the chain (ID: %s) was expecting an height of %d, an height of %d was calculated on verification", chainID.String(), expectedHeight, countedHeight)
		return errors.New(str)
	}

	// the chain is validated:
	return nil
}
