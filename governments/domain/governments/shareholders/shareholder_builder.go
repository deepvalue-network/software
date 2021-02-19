package shareholders

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/deepvalue-network/software/blockchain/domain/chains"
	"github.com/deepvalue-network/software/libs/hash"
)

type shareHolderBuilder struct {
	hashAdapter             hash.Adapter
	minHashesInShareHolders uint
	chain                   chains.Chain
	keys                    []hash.Hash
	power                   uint
	createdOn               *time.Time
}

func createShareHolderBuilder(
	hashAdapter hash.Adapter,
	minHashesInShareHolders uint,
) ShareHolderBuilder {
	out := shareHolderBuilder{
		hashAdapter:             hashAdapter,
		minHashesInShareHolders: minHashesInShareHolders,
		chain:                   nil,
		keys:                    nil,
		power:                   0,
		createdOn:               nil,
	}

	return &out
}

// Create initializes the builder
func (app *shareHolderBuilder) Create() ShareHolderBuilder {
	return createShareHolderBuilder(app.hashAdapter, app.minHashesInShareHolders)
}

// WithChain adds a chain to the builder
func (app *shareHolderBuilder) WithChain(chain chains.Chain) ShareHolderBuilder {
	app.chain = chain
	return app
}

// WithKeys add keys to the builder
func (app *shareHolderBuilder) WithKeys(keys []hash.Hash) ShareHolderBuilder {
	app.keys = keys
	return app
}

// WithPower add power to the builder
func (app *shareHolderBuilder) WithPower(power uint) ShareHolderBuilder {
	app.power = power
	return app
}

// CreatedOn adds creation time to the builder
func (app *shareHolderBuilder) CreatedOn(createdOn time.Time) ShareHolderBuilder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new ShareHolder instance
func (app *shareHolderBuilder) Now() (ShareHolder, error) {
	if app.chain == nil {
		return nil, errors.New("the chain is mandatory in order to build a ShareHolder instance")
	}

	if app.power <= 0 {
		return nil, errors.New("the power must be greater than zero (0) in order to build a ShareHolder instance")
	}

	if app.keys == nil {
		app.keys = []hash.Hash{}
	}

	amount := len(app.keys)
	if amount < int(app.minHashesInShareHolders) {
		str := fmt.Sprintf("there must be at least %d key hashes in order to build a ShareHolder instance, %d provided", app.minHashesInShareHolders, amount)
		return nil, errors.New(str)
	}

	if app.createdOn == nil {
		createdOn := time.Now().UTC()
		app.createdOn = &createdOn
	}

	data := [][]byte{
		app.chain.ID().Bytes(),
		[]byte(strconv.Itoa(int(app.power))),
	}

	for _, oneKey := range app.keys {
		data = append(data, oneKey.Bytes())
	}

	hash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createShareHolder(*hash, app.chain, app.keys, app.power, *app.createdOn), nil
}
