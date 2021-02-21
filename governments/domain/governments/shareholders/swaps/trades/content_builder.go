package trades

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/deepvalue-network/software/governments/domain/governments/shareholders/swaps/requests"
	"github.com/deepvalue-network/software/governments/domain/governments/shareholders/transfers/views"
	"github.com/deepvalue-network/software/libs/hash"
)

type contentBuilder struct {
	hashAdapter hash.Adapter
	minPubKeys  uint
	request     requests.Request
	transfer    views.Transfer
	to          []hash.Hash
	expiresOn   *time.Time
	createdOn   *time.Time
}

func createContentBuilder(
	hashAdapter hash.Adapter,
	minPubKeys uint,
) ContentBuilder {
	out := contentBuilder{
		hashAdapter: hashAdapter,
		minPubKeys:  minPubKeys,
		request:     nil,
		transfer:    nil,
		to:          nil,
		expiresOn:   nil,
		createdOn:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *contentBuilder) Create() ContentBuilder {
	return createContentBuilder(app.hashAdapter, app.minPubKeys)
}

// WithRequest adds a request to the builder
func (app *contentBuilder) WithRequest(request requests.Request) ContentBuilder {
	app.request = request
	return app
}

// WithTransfer adds a transfer to the builder
func (app *contentBuilder) WithTransfer(transfer views.Transfer) ContentBuilder {
	app.transfer = transfer
	return app
}

// To adds to public hashes to the builder
func (app *contentBuilder) To(to []hash.Hash) ContentBuilder {
	app.to = to
	return app
}

// ExpiresOn adds an expiration time to the builder
func (app *contentBuilder) ExpiresOn(expiresOn time.Time) ContentBuilder {
	app.expiresOn = &expiresOn
	return app
}

// CreatedOn adds a creation time to the builder
func (app *contentBuilder) CreatedOn(createdOn time.Time) ContentBuilder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new Content instance
func (app *contentBuilder) Now() (Content, error) {
	if app.request == nil {
		return nil, errors.New("the request is mandatory in order to build a trade Content instance")
	}

	if app.transfer == nil {
		return nil, errors.New("the transfer is mandatory in order to build a trade Content instance")
	}

	if app.expiresOn == nil {
		return nil, errors.New("the expiration time is mandatory in order to build a trade Content instance")
	}

	if app.to == nil {
		app.to = []hash.Hash{}
	}

	amount := len(app.to)
	if amount < int(app.minPubKeys) {
		str := fmt.Sprintf("there must be at least %d public key hashes (to) in order to build a trade Content instance, %d provided", app.minPubKeys, amount)
		return nil, errors.New(str)
	}

	if app.createdOn == nil {
		createdOn := time.Now().UTC()
		app.createdOn = &createdOn
	}

	if app.expiresOn.Before(*app.createdOn) {
		str := fmt.Sprintf("the expiration time (%s) cannot be before the creation time (%s)", app.expiresOn.String(), app.createdOn.String())
		return nil, errors.New(str)
	}

	// make sure the amounts matches:
	reqAmount := app.request.Content().Amount()
	trsfAmount := app.transfer.Content().Section().Amount()
	if reqAmount != trsfAmount {
		str := fmt.Sprintf("the requested amount (%d) does not match the transfer amount (%d)", reqAmount, trsfAmount)
		return nil, errors.New(str)
	}

	// make sure the pubkeys matches:
	reqPubKeys := app.request.Content().To()
	trsfNewOwnerPubKeys := app.transfer.Content().NewOwner()
	if !compareHashes(reqPubKeys, trsfNewOwnerPubKeys) {
		return nil, errors.New("the request to pubKey hashes do not matche the transfer new owner pubKey hashes")
	}

	data := [][]byte{
		app.request.Hash().Bytes(),
		app.transfer.Hash().Bytes(),
		[]byte(strconv.Itoa(app.expiresOn.Second())),
		[]byte(strconv.Itoa(app.createdOn.Second())),
	}

	for _, oneHash := range app.to {
		data = append(data, oneHash.Bytes())
	}

	hash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createContent(*hash, app.request, app.transfer, app.to, *app.expiresOn, *app.createdOn), nil
}
