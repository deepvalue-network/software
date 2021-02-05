package spends

import (
	"errors"
	"strconv"
	"time"

	domain_genesis "github.com/deepvalue-network/software/diamonds/domain/genesis"
	"github.com/deepvalue-network/software/libs/hash"
)

type builder struct {
	hashAdapter  hash.Adapter
	amount       uint64
	hashedAmount *hash.Hash
	seed         string
	hashedSeed   *hash.Hash
	gen          domain_genesis.Genesis
	createdOn    *time.Time
}

func createBuilder(hashAdapter hash.Adapter) Builder {
	out := builder{
		hashAdapter:  hashAdapter,
		amount:       0,
		hashedAmount: nil,
		seed:         "",
		hashedSeed:   nil,
		gen:          nil,
		createdOn:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithAmount adds an amount to the builder
func (app *builder) WithAmount(amount uint64) Builder {
	app.amount = amount
	return app
}

// WithHashedAmount adds an hashed amount to the builder
func (app *builder) WithHashedAmount(hashedAmount hash.Hash) Builder {
	app.hashedAmount = &hashedAmount
	return app
}

// WithSeed adds a seed to the builder
func (app *builder) WithSeed(seed string) Builder {
	app.seed = seed
	return app
}

// WithHashedSeed adds an hashed seed to the builder
func (app *builder) WithHashedSeed(hashedSeed hash.Hash) Builder {
	app.hashedSeed = &hashedSeed
	return app
}

// WithGenesis adds a genesis to the builder
func (app *builder) WithGenesis(gen domain_genesis.Genesis) Builder {
	app.gen = gen
	return app
}

// CreatedOn adds a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new Genesis instance
func (app *builder) Now() (Genesis, error) {
	if app.seed != "" && app.amount > 0 {
		hashedSeed, err := app.hashAdapter.FromBytes([]byte(app.seed))
		if err != nil {
			return nil, err
		}

		hashedAmount, err := app.hashAdapter.FromMultiBytes([][]byte{
			[]byte(app.seed),
			[]byte(strconv.Itoa(int(app.amount))),
		})

		if err != nil {
			return nil, err
		}

		app.hashedSeed = hashedSeed
		app.hashedAmount = hashedAmount
	}

	if app.hashedSeed == nil {
		return nil, errors.New("the hashed seed is mandatory in order to build a spent Genesis instance")
	}

	if app.hashedAmount == nil {
		return nil, errors.New("the hashed amount is mandatory in order to build a spent Genesis instance")
	}

	if app.gen == nil {
		return nil, errors.New("the genesis is mandatory in order to build a spent Genesis instance")
	}

	if app.createdOn == nil {
		createdOn := time.Now().UTC()
		app.createdOn = &createdOn
	}

	hsh, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.hashedSeed.Bytes(),
		app.hashedAmount.Bytes(),
		app.gen.Hash().Bytes(),
		[]byte(strconv.Itoa(int(app.createdOn.UnixNano()))),
	})

	if err != nil {
		return nil, err
	}

	return createGenesis(*hsh, *app.hashedAmount, *app.hashedSeed, app.gen, *app.createdOn), nil
}
