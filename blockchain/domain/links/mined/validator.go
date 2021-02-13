package mined

import (
	"errors"
	"fmt"
	"strings"

	"github.com/deepvalue-network/software/blockchain/domain/blocks"
	"github.com/deepvalue-network/software/blockchain/domain/genesis"
	"github.com/deepvalue-network/software/libs/hash"
)

type validator struct {
	hashAdapter         hash.Adapter
	minedLinkRepository Repository
}

func createValidator(
	hashAdapter hash.Adapter,
	minedLinkRepository Repository,
) Validator {
	out := validator{
		hashAdapter:         hashAdapter,
		minedLinkRepository: minedLinkRepository,
	}

	return &out
}

// Execute executes the validator
func (app *validator) Execute(gen genesis.Genesis, minedLink Link, root blocks.Block) (uint, uint, error) {
	// fetch the difficulty:
	diff := int(gen.LinkDifficulty())

	// create the difficulty prefix:
	prefix := ""
	miningValue := gen.MiningValue()
	for i := 0; i < diff; i++ {
		prefix = fmt.Sprintf("%s%d", prefix, miningValue)
	}

	// hash the results:
	results := minedLink.Results()
	linkHash := minedLink.Link().Hash()
	resultsHash, err := minerHash(results, linkHash, app.hashAdapter)
	if err != nil {
		return 0, 0, err
	}

	// make sure the results contains the right diffiulty:
	if !strings.HasPrefix(resultsHash.String(), prefix) {
		str := fmt.Sprintf("the prefix (%s) was expected on the result hash (hash: %s)", prefix, resultsHash.String())
		return 0, 0, errors.New(str)
	}

	// execute the previous mined link:
	prevMinedLinkHash := minedLink.Link().PrevMinedLink()
	prevMinedLink, err := app.minedLinkRepository.Retrieve(prevMinedLinkHash)
	if err != nil {
		return 0, 0, err
	}

	prevLinkTotalHashes, prevLinkHeight, err := app.Execute(gen, prevMinedLink, root)
	if err != nil {
		return 0, 0, err
	}

	amountHashes := uint(len(minedLink.Link().NextBlock().Hashes())) + prevLinkTotalHashes
	height := prevLinkHeight + 1
	return amountHashes, height, nil
}
