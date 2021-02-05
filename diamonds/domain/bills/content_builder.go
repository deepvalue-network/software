package bills

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/steve-care-software/products/diamonds/domain/genesis/spends/views"
	"github.com/steve-care-software/products/libs/cryptography/pk/signature"
	"github.com/steve-care-software/products/libs/hash"
)

type contentBuilder struct {
	hashAdapter        hash.Adapter
	minPubKeysAmount   uint
	originGenesis      views.Genesis
	originBill         ViewBill
	amount             uint64
	hashedAmount       *hash.Hash
	seed               string
	hashedSeed         *hash.Hash
	pubKeysOwner       []signature.PublicKey
	hashedPubKeysOwner []hash.Hash
	createdOn          *time.Time
}

func createContentBuilder(hashAdapter hash.Adapter, minPubKeysAmount uint) ContentBuilder {
	out := contentBuilder{
		hashAdapter:        hashAdapter,
		minPubKeysAmount:   minPubKeysAmount,
		originGenesis:      nil,
		originBill:         nil,
		amount:             0,
		hashedAmount:       nil,
		seed:               "",
		hashedSeed:         nil,
		pubKeysOwner:       nil,
		hashedPubKeysOwner: nil,
		createdOn:          nil,
	}

	return &out
}

// Create initializes the builder
func (app *contentBuilder) Create() ContentBuilder {
	return createContentBuilder(app.hashAdapter, app.minPubKeysAmount)
}

// WithOriginGenesis adds an origin genesis to the builder
func (app *contentBuilder) WithOriginGenesis(originGenesis views.Genesis) ContentBuilder {
	app.originGenesis = originGenesis
	return app
}

// WithOriginBill adds an origin bill to the builder
func (app *contentBuilder) WithOriginBill(originBill ViewBill) ContentBuilder {
	app.originBill = originBill
	return app
}

// WithAmount adds an amount to the builder
func (app *contentBuilder) WithAmount(amount uint64) ContentBuilder {
	app.amount = amount
	return app
}

// WithHashedAmount adds an hashedAmount to the builder
func (app *contentBuilder) WithHashedAmount(hashedAmount hash.Hash) ContentBuilder {
	app.hashedAmount = &hashedAmount
	return app
}

// WithSeed adds a seed to the builder
func (app *contentBuilder) WithSeed(seed string) ContentBuilder {
	app.seed = seed
	return app
}

// WithHashedSeed adds an hashed seed to the builder
func (app *contentBuilder) WithHashedSeed(hashedSeed hash.Hash) ContentBuilder {
	app.hashedSeed = &hashedSeed
	return app
}

// WithPubKeysOwner add pubKeys owner to the builder
func (app *contentBuilder) WithPubKeysOwner(pubKeysOwner []signature.PublicKey) ContentBuilder {
	app.pubKeysOwner = pubKeysOwner
	return app
}

// WithHashedSeed adds an hashed seed to the builder
func (app *contentBuilder) WithHashedPubKeysOwner(hashedPubKeysOwner []hash.Hash) ContentBuilder {
	app.hashedPubKeysOwner = hashedPubKeysOwner
	return app
}

// CreatedOn adds a creation time to the builder
func (app *contentBuilder) CreatedOn(createdOn time.Time) ContentBuilder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new Bill instance
func (app *contentBuilder) Now() (Content, error) {
	var origin Origin
	if app.originGenesis != nil {
		origin = createOriginWithGenesis(app.originGenesis)
	}

	if app.originBill != nil {
		origin = createOriginWithBill(app.originBill)
	}

	if origin == nil {
		return nil, errors.New("the origin (view genesis/bill) is mandatory in order to build a bill Content instance")
	}

	if app.amount <= 0 {
		hsh, err := app.hashAdapter.FromBytes([]byte(strconv.Itoa(int(app.amount))))
		if err != nil {
			return nil, err
		}

		app.hashedAmount = hsh
	}

	if app.seed != "" {
		hsh, err := app.hashAdapter.FromBytes([]byte(app.seed))
		if err != nil {
			return nil, err
		}

		app.hashedSeed = hsh
	}

	if app.hashedAmount == nil {
		return nil, errors.New("the hashed amount is mandatory in order to build a bill Content instance")
	}

	if app.hashedSeed == nil {
		return nil, errors.New("the hashed seed is mandatory in order to build a bill Content instance")
	}

	if app.pubKeysOwner != nil {
		if len(app.pubKeysOwner) < int(app.minPubKeysAmount) {
			str := fmt.Sprintf("there must be at least %d public keys in the owner's list, %d provided", app.minPubKeysAmount, len(app.pubKeysOwner))
			return nil, errors.New(str)
		}

		hashes := []hash.Hash{}
		for _, onePubKey := range app.pubKeysOwner {
			hsh, err := app.hashAdapter.FromBytes([]byte(onePubKey.String()))
			if err != nil {
				return nil, err
			}

			hashes = append(hashes, *hsh)
		}

		app.hashedPubKeysOwner = hashes
	}

	if app.hashedPubKeysOwner == nil {
		return nil, errors.New("the hashed pubKeys owner are mandatory in order to build a bill Content instance")
	}

	if app.createdOn == nil {
		createdOn := time.Now().UTC()
		app.createdOn = &createdOn
	}

	data := [][]byte{
		origin.Hash().Bytes(),
		app.hashedAmount.Bytes(),
		app.hashedSeed.Bytes(),
		[]byte(strconv.Itoa(int(app.createdOn.UnixNano()))),
	}

	for _, onePubKeyHash := range app.hashedPubKeysOwner {
		data = append(data, onePubKeyHash.Bytes())
	}

	hsh, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createContent(*hsh, origin, *app.hashedAmount, *app.hashedSeed, app.hashedPubKeysOwner, *app.createdOn), nil
}
