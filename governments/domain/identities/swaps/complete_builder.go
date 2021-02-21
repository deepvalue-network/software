package swaps

import (
	"errors"

	"github.com/deepvalue-network/software/governments/domain/governments/shareholders/swaps/closes"
	"github.com/deepvalue-network/software/governments/domain/governments/shareholders/swaps/trades"
)

type completeBuilder struct {
	trade trades.Trade
	close closes.Close
}

func createCompleteBuilder() CompleteBuilder {
	out := completeBuilder{
		trade: nil,
		close: nil,
	}

	return &out
}

// Create initializes the builder
func (app *completeBuilder) Create() CompleteBuilder {
	return createCompleteBuilder()
}

// WithTrade adds a trade to the builder
func (app *completeBuilder) WithTrade(trade trades.Trade) CompleteBuilder {
	app.trade = trade
	return app
}

// WithClose adds a close to the builder
func (app *completeBuilder) WithClose(close closes.Close) CompleteBuilder {
	app.close = close
	return app
}

// Now builds a new Complete instance
func (app *completeBuilder) Now() (Complete, error) {
	if app.trade == nil {
		return nil, errors.New("the trade is mandatory in order to build a Complete instance")
	}

	if app.close != nil {
		return createCompleteWithClose(app.trade, app.close), nil
	}

	return createComplete(app.trade), nil
}
