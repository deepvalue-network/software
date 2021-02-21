package requests

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/deepvalue-network/software/governments/domain/governments"
	"github.com/deepvalue-network/software/governments/domain/governments/shareholders/transfers/views"
	"github.com/deepvalue-network/software/libs/hash"
)

type contentBuilder struct {
	hashAdapter hash.Adapter
	minPubKeys  uint
	from        governments.Government
	stake       views.Section
	forGov      governments.Government
	to          []hash.Hash
	amount      uint
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
		from:        nil,
		stake:       nil,
		forGov:      nil,
		to:          nil,
		amount:      0,
		expiresOn:   nil,
		createdOn:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *contentBuilder) Create() ContentBuilder {
	return createContentBuilder(app.hashAdapter, app.minPubKeys)
}

// From adds a from government to the builder
func (app *contentBuilder) From(from governments.Government) ContentBuilder {
	app.from = from
	return app
}

// WithStake adds a stake to the builder
func (app *contentBuilder) WithStake(stake views.Section) ContentBuilder {
	app.stake = stake
	return app
}

// For adds a for government to the builder
func (app *contentBuilder) For(forGov governments.Government) ContentBuilder {
	app.forGov = forGov
	return app
}

// To adds a to public key hashes to the builder
func (app *contentBuilder) To(to []hash.Hash) ContentBuilder {
	app.to = to
	return app
}

// WithAmount adds an amount to the builder
func (app *contentBuilder) WithAmount(amount uint) ContentBuilder {
	app.amount = amount
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
	if app.from == nil {
		return nil, errors.New("the from government is mandatory in order to build a request Content instance")
	}

	if app.stake == nil {
		return nil, errors.New("the stake is mandatory in order to build a request Content instance")
	}

	if app.forGov == nil {
		return nil, errors.New("the for government is mandatory in order to build a request Content instance")
	}

	if app.expiresOn == nil {
		return nil, errors.New("the expiration time is mandatory in order to build a request Content instance")
	}

	if app.amount <= 0 {
		return nil, errors.New("the amount is mandatory in order to build a request Content instance")
	}

	if app.to == nil {
		app.to = []hash.Hash{}
	}

	amount := len(app.to)
	if amount < int(app.minPubKeys) {
		str := fmt.Sprintf("there must be at least %d public key hashes (to) in order to build a request Content instance, %d provided", app.minPubKeys, amount)
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

	data := [][]byte{
		app.from.Hash().Bytes(),
		app.stake.Hash().Bytes(),
		app.forGov.Hash().Bytes(),
		[]byte(strconv.Itoa(int(app.amount))),
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

	return createContent(*hash, app.from, app.stake, app.forGov, app.to, app.amount, *app.expiresOn, *app.createdOn), nil
}
