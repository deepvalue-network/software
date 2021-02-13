package mined

import (
	"errors"
	"fmt"
	"strings"

	"github.com/deepvalue-network/software/blockchain/domain/genesis"
	"github.com/deepvalue-network/software/libs/hash"
)

type validator struct {
	hashAdapter hash.Adapter
}

func createValidator(
	hashAdapter hash.Adapter,
) Validator {
	out := validator{
		hashAdapter: hashAdapter,
	}

	return &out
}

// Execute executes a validator
func (app *validator) Execute(gen genesis.Genesis, block Block) error {
	// calculate difficulty:
	baseDiff := gen.BlockBaseDifficulty()
	incrPerHashDiff := gen.BlockIncreasePerHashDifficulty()
	hashes := block.Block().Hashes()
	diff := int(calculateDifficulty(baseDiff, incrPerHashDiff, len(hashes)))

	// create the difficulty prefix:
	prefix := ""
	miningValue := gen.MiningValue()
	for i := 0; i < diff; i++ {
		prefix = fmt.Sprintf("%s%d", prefix, miningValue)
	}

	// hash the results:
	results := block.Results()
	blockhash := block.Block().Tree().Head()
	resultsHash, err := minerHash(results, blockhash, app.hashAdapter)
	if err != nil {
		return err
	}

	// make sure the results contains the right diffiulty:
	if !strings.HasPrefix(resultsHash.String(), prefix) {
		str := fmt.Sprintf("the prefix (%s) was expected on the result hash (hash: %s)", prefix, resultsHash.String())
		return errors.New(str)
	}

	return nil
}
