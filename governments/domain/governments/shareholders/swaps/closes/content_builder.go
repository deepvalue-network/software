package closes

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/deepvalue-network/software/governments/domain/governments/shareholders/swaps/trades"
	"github.com/deepvalue-network/software/libs/hash"
)

type contentBuilder struct {
	hashAdapter hash.Adapter
	trade       trades.Trade
	createdOn   *time.Time
}

func createContentBuilder(
	hashAdapter hash.Adapter,
) ContentBuilder {
	out := contentBuilder{
		hashAdapter: hashAdapter,
		trade:       nil,
		createdOn:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *contentBuilder) Create() ContentBuilder {
	return createContentBuilder(app.hashAdapter)
}

// WithTrade adds a trade to the builder
func (app *contentBuilder) WithTrade(trade trades.Trade) ContentBuilder {
	app.trade = trade
	return app
}

// CreatedOn adds a creation time to the builder
func (app *contentBuilder) CreatedOn(createdOn time.Time) ContentBuilder {
	app.createdOn = &createdOn
	return app
}

// CreatedOn adds a creation time to the builder
func (app *contentBuilder) Now() (Content, error) {
	if app.trade == nil {
		return nil, errors.New("the trade is mandatory in order to build a close Content instance")
	}

	if app.createdOn == nil {
		createdOn := time.Now().UTC()
		app.createdOn = &createdOn
	}

	if app.createdOn.After(app.trade.Content().ExpiresOn()) {
		str := fmt.Sprintf("the trade cannot be closed because it is expired (expiresOn: %s, currentTime: %s)", app.trade.Content().ExpiresOn().String(), app.createdOn.String())
		return nil, errors.New(str)
	}

	hash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.trade.Hash().Bytes(),
		[]byte(strconv.Itoa(app.createdOn.Second())),
	})

	if err != nil {
		return nil, err
	}

	return createContent(*hash, app.trade, *app.createdOn), nil
}
