package genesis

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/deepvalue-network/software/blockchain/domain/chains"
	"github.com/deepvalue-network/software/libs/hash"
)

type builder struct {
	hashAdapter        hash.Adapter
	minPubKeysInOwner  uint
	amount             uint64
	chain              chains.Chain
	hashedPubKeysOwner []hash.Hash
	createdOn          *time.Time
	activeOn           *time.Time
}

func createBuilder(
	hashAdapter hash.Adapter,
	minPubKeysInOwner uint,
) Builder {
	out := builder{
		hashAdapter:        hashAdapter,
		minPubKeysInOwner:  minPubKeysInOwner,
		amount:             0,
		chain:              nil,
		hashedPubKeysOwner: nil,
		createdOn:          nil,
		activeOn:           nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter, app.minPubKeysInOwner)
}

// WithAmount adds an amount to the builder
func (app *builder) WithAmount(amount uint64) Builder {
	app.amount = amount
	return app
}

// WithChain adds a chain to the builder
func (app *builder) WithChain(chain chains.Chain) Builder {
	app.chain = chain
	return app
}

// WithHashedPubKeysOwner adds an hashed pubkeys owner to the builder
func (app *builder) WithHashedPubKeysOwner(hashedPubKeysOwner []hash.Hash) Builder {
	app.hashedPubKeysOwner = hashedPubKeysOwner
	return app
}

// CreatedOn adds a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.createdOn = &createdOn
	return app
}

// ActiveOn adds an activeOn time to the builder
func (app *builder) ActiveOn(activeOn time.Time) Builder {
	app.activeOn = &activeOn
	return app
}

// Now builds a new Genesis instance
func (app *builder) Now() (Genesis, error) {
	if app.amount == 0 {
		return nil, errors.New("the amount is mandatory in order to build a Genesis instance")
	}

	if app.chain == nil {
		return nil, errors.New("the chain is mandatory in order to build a Genesis instance")
	}

	if app.hashedPubKeysOwner != nil && len(app.hashedPubKeysOwner) <= 0 {
		app.hashedPubKeysOwner = nil
	}

	if len(app.hashedPubKeysOwner) < int(app.minPubKeysInOwner) {
		str := fmt.Sprintf("there must be at least %d hashed public keys in the owner, %d provided", app.minPubKeysInOwner, len(app.hashedPubKeysOwner))
		return nil, errors.New(str)
	}

	if app.createdOn == nil {
		createdOn := time.Now().UTC()
		app.createdOn = &createdOn
	}

	if app.activeOn == nil {
		activeOn := time.Now().UTC()
		app.activeOn = &activeOn
	}

	if app.activeOn.Before(*app.createdOn) {
		str := fmt.Sprintf("the activation time (%s), cannot be before the creation time (%s)", app.activeOn.String(), app.createdOn.String())
		return nil, errors.New(str)
	}

	data := [][]byte{
		[]byte(strconv.Itoa(int(app.amount))),
		app.chain.ID().Bytes(),
		[]byte(strconv.Itoa(int(app.createdOn.UnixNano()))),
		[]byte(strconv.Itoa(int(app.activeOn.UnixNano()))),
	}

	for _, oneHash := range app.hashedPubKeysOwner {
		data = append(data, oneHash.Bytes())
	}

	hsh, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createGenesis(*hsh, app.amount, app.chain, app.hashedPubKeysOwner, *app.createdOn, *app.activeOn), nil

}
