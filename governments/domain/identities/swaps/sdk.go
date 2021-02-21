package swaps

import (
	"github.com/deepvalue-network/software/governments/domain/governments/shareholders/swaps/closes"
	"github.com/deepvalue-network/software/governments/domain/governments/shareholders/swaps/trades"
	"github.com/deepvalue-network/software/governments/domain/identities/transfers"
	"github.com/deepvalue-network/software/libs/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewIncomingBuilder creates a new incoming builder instance
func NewIncomingBuilder() IncomingBuilder {
	return createIncomingBuilder()
}

// NewCompleteBuilder creates a new complete builder instance
func NewCompleteBuilder() CompleteBuilder {
	return createCompleteBuilder()
}

// Builder represents a swap builder
type Builder interface {
	Create() Builder
	WithIncoming(incoming Incoming) Builder
	WithOutgoing(outgoing Complete) Builder
	Now() (Swap, error)
}

// Swap represents a swap
type Swap interface {
	IsIncoming() bool
	Incoming() Incoming
	IsOutgoing() bool
	Outgoing() Complete
}

// IncomingBuilder represents an incoming builder
type IncomingBuilder interface {
	Create() IncomingBuilder
	WithComplete(complete Complete) IncomingBuilder
	WithIncoming(transfer transfers.Transfer) IncomingBuilder
	Now() (Incoming, error)
}

// Incoming represents an incoming swap
type Incoming interface {
	Complete() Complete
	Incoming() transfers.Transfer
}

// CompleteBuilder represents a complete builder
type CompleteBuilder interface {
	Create() CompleteBuilder
	WithTrade(trade trades.Trade) CompleteBuilder
	WithClose(close closes.Close) CompleteBuilder
	Now() (Complete, error)
}

// Complete represents a complete swap
type Complete interface {
	Trade() trades.Trade
	HasClose() bool
	Close() closes.Close
}

// Repository represents a swap repository
type Repository interface {
	Retrieve(requestHash hash.Hash) (Swap, error)
}
