package services

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/steve-care-software/products/blockchain/domain/genesis"
	"github.com/steve-care-software/products/libs/hash"
)

type miner struct {
	hashAdapter hash.Adapter
}

func createMiner(
	hashAdapter hash.Adapter,
) *miner {
	out := miner{
		hashAdapter: hashAdapter,
	}

	return &out
}

// Test excutes a mining test using the given difficulty
func (app *miner) Test(difficulty uint) (string, *time.Duration, error) {
	data := strconv.Itoa(time.Now().UTC().Nanosecond())
	hsh, err := app.hashAdapter.FromBytes([]byte(data))

	if err != nil {
		return "", nil, err
	}

	return app.Mine(genesis.DefaultMiningValue, difficulty, *hsh)
}

// Mine executes mining using the given hash
func (app *miner) Mine(miningValue uint8, difficulty uint, hash hash.Hash) (string, *time.Duration, error) {
	// fetch the current time:
	beginsOn := time.Now().UTC()

	// create the requested prefix:
	requestedPrefix, err := app.prefix(miningValue, difficulty)
	if err != nil {
		return "", nil, err
	}

	// execute the mining:
	results, err := app.mineRecursively(
		requestedPrefix,
		hash,
		"",
	)

	if err != nil {
		return "", nil, nil
	}

	endsOn := time.Now().UTC()
	elapsed := endsOn.Sub(beginsOn)
	return results, &elapsed, nil
}

func (app *miner) mineRecursively(
	requestedPrefix string,
	hsh hash.Hash,
	baseStr string,
) (string, error) {
	str := ""
	for i := uint(0); i <= maxMiningValue; i++ {
		str = fmt.Sprintf("%s%s", baseStr, []byte(strconv.Itoa(int(i))))
		res, err := app.hashAdapter.FromMultiBytes([][]byte{
			[]byte(str),
			hsh.Bytes(),
		})

		if err != nil {
			return "", err
		}

		if strings.HasPrefix(res.String(), requestedPrefix) {
			return str, nil
		}
	}

	for i := 0; i < maxMiningTries; i++ {
		results, err := app.mineRecursively(requestedPrefix, hsh, str)
		if err != nil {
			continue
		}

		return results, nil
	}

	return "", errors.New("the mining was impossible")
}

func (app *miner) prefix(miningValue uint8, difficulty uint) (string, error) {
	output := ""
	for i := 0; i < int(difficulty); i++ {
		output = fmt.Sprintf("%s%d", output, miningValue)
	}

	return output, nil
}
