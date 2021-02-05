package views

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/deepvalue-network/software/diamonds/domain/genesis/spends"
	"github.com/deepvalue-network/software/libs/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	gen         spends.Genesis
	seed        string
	amount      uint64
}

func createBuilder(hashAdapter hash.Adapter) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		gen:         nil,
		seed:        "",
		amount:      0,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithGenesis adds a genesis to the builder
func (app *builder) WithGenesis(gen spends.Genesis) Builder {
	app.gen = gen
	return app
}

// WithSeed adds a seed to the builder
func (app *builder) WithSeed(seed string) Builder {
	app.seed = seed
	return app
}

// WithAmount adds an amount to the builder
func (app *builder) WithAmount(amount uint64) Builder {
	app.amount = amount
	return app
}

// Now builds a new Genesis instance
func (app *builder) Now() (Genesis, error) {
	if app.gen == nil {
		return nil, errors.New("the spent genesis is mandatory in order to build a view Genesis instance")
	}

	if app.seed == "" {
		return nil, errors.New("the seed is mandatory in order to build a view Genesis instance")
	}

	// make sure the seed works with the given spent genesis:
	seedHash, err := app.hashAdapter.FromBytes([]byte(app.seed))
	if err != nil {
		return nil, err
	}

	if !app.gen.Seed().Compare(*seedHash) {
		str := fmt.Sprintf(
			"the given seed (seed: %s), when hashed (hash: %s) does not match the hashed seed (hash: %s inside the given spent Genesis (hash: %s)",
			app.seed,
			seedHash.String(),
			app.gen.Seed().String(),
			app.gen.Hash().String(),
		)
		return nil, errors.New(str)
	}

	// make sure the hashed amount fits:
	hashedAmount, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.seed),
		[]byte(strconv.Itoa(int(app.amount))),
	})

	if err != nil {
		return nil, err
	}

	if !app.gen.Amount().Compare(*hashedAmount) {
		str := fmt.Sprintf(
			"the given amount (amount: %d), when hashed by combing it with the seed (seed: %s, hash: %s) does not match the hashed amount (hash: %s inside the given spent Genesis (hash: %s)",
			app.amount,
			app.seed,
			hashedAmount.String(),
			app.gen.Amount().String(),
			app.gen.Hash().String(),
		)
		return nil, errors.New(str)
	}

	hsh, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.gen.Hash().Bytes(),
		[]byte(app.seed),
	})

	if err != nil {
		return nil, err
	}

	return createGenesis(*hsh, app.gen, app.seed, app.amount), nil
}
