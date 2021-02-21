package swaps

import (
	"github.com/deepvalue-network/software/governments/domain/governments/shareholders/swaps/closes"
	"github.com/deepvalue-network/software/governments/domain/governments/shareholders/swaps/trades"
)

type complete struct {
	trade trades.Trade
	close closes.Close
}

func createComplete(
	trade trades.Trade,
) Complete {
	return createCompleteInternally(trade, nil)
}

func createCompleteWithClose(
	trade trades.Trade,
	close closes.Close,
) Complete {
	return createCompleteInternally(trade, close)
}

func createCompleteInternally(
	trade trades.Trade,
	close closes.Close,
) Complete {
	out := complete{
		trade: trade,
		close: close,
	}

	return &out
}

// Trade returns the trade
func (obj *complete) Trade() trades.Trade {
	return obj.trade
}

// HasClose returns true if there is a close, false otherwise
func (obj *complete) HasClose() bool {
	return obj.close != nil
}

// Close returns the close, if any
func (obj *complete) Close() closes.Close {
	return obj.close
}
