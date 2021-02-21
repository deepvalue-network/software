package swaps

import (
	"github.com/deepvalue-network/software/governments/domain/governments/shareholders/swaps/closes"
	"github.com/deepvalue-network/software/governments/domain/governments/shareholders/swaps/trades"
	"github.com/deepvalue-network/software/governments/domain/identities/shareholders/transactions/transfers/incomings"
	"github.com/deepvalue-network/software/libs/hash"
)

// Builder represents a builder
type Builder interface {
	Create() Builder
	WithSwaps(swaps []Swap) Builder
	Now() (Swaps, error)
}

// Swaps represents swaps
type Swaps interface {
	Hash() hash.Hash
	All() []Swap
}

// SwapBuilder represents a swap builder
type SwapBuilder interface {
	Create() SwapBuilder
	WithIncoming(incoming Incoming) SwapBuilder
	WithOutgoing(outgoing Complete) SwapBuilder
	Now() (Swap, error)
}

// Swap represents a swap
type Swap interface {
	Hash() hash.Hash
	IsIncoming() bool
	Incoming() Incoming
	IsOutgoing() bool
	Outgoing() Complete
}

// IncomingBuilder represents an incoming builder
type IncomingBuilder interface {
	Create() IncomingBuilder
	WithComplete(complete Complete) IncomingBuilder
	WithIncoming(incoming incomings.Incoming) IncomingBuilder
	Now() (Incoming, error)
}

// Incoming represents an incoming swap
type Incoming interface {
	Hash() hash.Hash
	Complete() Complete
	Incoming() incomings.Incoming
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
	Hash() hash.Hash
	Trade() trades.Trade
	HasClose() bool
	Close() closes.Close
}
