package transactions

import (
	"github.com/deepvalue-network/software/governments/domain/identities/shareholders/transactions/payments"
	"github.com/deepvalue-network/software/governments/domain/identities/shareholders/transactions/swaps"
	"github.com/deepvalue-network/software/governments/domain/identities/shareholders/transactions/transfers"
)

// Builder represents a transactions builder
type Builder interface {
	Create() Builder
	WithSwaps(swaps swaps.Swaps) Builder
	WithPayments(payments payments.Payments) Builder
	WithTransfers(transfers transfers.Transfers) Builder
	Now() (Transactions, error)
}

// Transactions represents the transactions
type Transactions interface {
	Swaps() swaps.Swaps
	Payments() payments.Payments
	Transfers() transfers.Transfers
}
