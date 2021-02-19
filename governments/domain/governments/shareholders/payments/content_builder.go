package payments

import (
	"errors"
	"strconv"
	"time"

	"github.com/deepvalue-network/software/governments/domain/governments/shareholders"
	"github.com/deepvalue-network/software/libs/hash"
)

type contentBuilder struct {
	hashAdapter hash.Adapter
	holder      shareholders.ShareHolder
	amount      uint
	createdOn   *time.Time
}

func createContentBuilder(
	hashAdapter hash.Adapter,
) ContentBuilder {
	out := contentBuilder{
		hashAdapter: hashAdapter,
		holder:      nil,
		amount:      0,
		createdOn:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *contentBuilder) Create() ContentBuilder {
	return createContentBuilder(app.hashAdapter)
}

// WithShareHolder adds a shareHolder to the builder
func (app *contentBuilder) WithShareHolder(shareHolder shareholders.ShareHolder) ContentBuilder {
	app.holder = shareHolder
	return app
}

// WithAmount adds an amount to the builder
func (app *contentBuilder) WithAmount(amount uint) ContentBuilder {
	app.amount = amount
	return app
}

// CreatedOn adds a creation time to the builder
func (app *contentBuilder) CreatedOn(createdOn time.Time) ContentBuilder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new Content instance
func (app *contentBuilder) Now() (Content, error) {
	if app.holder == nil {
		return nil, errors.New("the shareHolder is mandatory in order to build a payment Content instance")
	}

	if app.amount <= 0 {
		return nil, errors.New("the amount must be greater than zero (0) in order to build a payment Content instance")
	}

	if app.createdOn == nil {
		createdOn := time.Now().UTC()
		app.createdOn = &createdOn
	}

	hash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.holder.Hash().Bytes(),
		[]byte(strconv.Itoa(int(app.amount))),
		[]byte(strconv.Itoa(app.createdOn.Second())),
	})

	if err != nil {
		return nil, err
	}

	return createContent(*hash, app.holder, app.amount, *app.createdOn), nil
}
