package transfers

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/deepvalue-network/software/libs/hash"
)

type contentBuilder struct {
	hashAdapter          hash.Adapter
	minPubKeysPerRingSig uint
	origin               *hash.Hash
	amount               *hash.Hash
	owner                []hash.Hash
	createdOn            *time.Time
}

func createContentBuilder(
	hashAdapter hash.Adapter,
	minPubKeysPerRingSig uint,
) ContentBuilder {
	out := contentBuilder{
		hashAdapter:          hashAdapter,
		minPubKeysPerRingSig: minPubKeysPerRingSig,
		origin:               nil,
		amount:               nil,
		owner:                nil,
		createdOn:            nil,
	}

	return &out
}

// Create initializes the builder
func (app *contentBuilder) Create() ContentBuilder {
	return createContentBuilder(app.hashAdapter, app.minPubKeysPerRingSig)
}

// WithOrigin adds an origin to the builder
func (app *contentBuilder) WithOrigin(origin hash.Hash) ContentBuilder {
	app.origin = &origin
	return app
}

// WithAmount adds an amount to the builder
func (app *contentBuilder) WithAmount(amount hash.Hash) ContentBuilder {
	app.amount = &amount
	return app
}

// WithOwner adds an owner to the builder
func (app *contentBuilder) WithOwner(owner []hash.Hash) ContentBuilder {
	app.owner = owner
	return app
}

// CreatedOn adds a creation time to the builder
func (app *contentBuilder) CreatedOn(createdOn time.Time) ContentBuilder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new Content instance
func (app *contentBuilder) Now() (Content, error) {
	if app.origin == nil {
		return nil, errors.New("the origin is mandatory in order to build a transfer Content instance")
	}

	if app.amount == nil {
		return nil, errors.New("the amount is mandatory in order to build a transfer Content instance")
	}

	if app.createdOn == nil {
		createdOn := time.Now().UTC()
		app.createdOn = &createdOn
	}

	if app.owner == nil {
		app.owner = []hash.Hash{}
	}

	amount := len(app.owner)
	if amount < int(app.minPubKeysPerRingSig) {
		str := fmt.Sprintf("the owner must contain at least %d pubkey hashes, %d provided", app.minPubKeysPerRingSig, amount)
		return nil, errors.New(str)
	}

	data := [][]byte{
		app.origin.Bytes(),
		app.amount.Bytes(),
		[]byte(strconv.Itoa(app.createdOn.Second())),
	}

	for _, oneHash := range app.owner {
		data = append(data, oneHash.Bytes())
	}

	hash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createContent(*hash, *app.origin, *app.amount, app.owner, *app.createdOn), nil

}
