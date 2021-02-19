package transactions

import (
	"github.com/deepvalue-network/software/governments/domain/shareholders/transactions/payments"
    "github.com/deepvalue-network/software/governments/domain/shareholders/transactions/transfers"
)

// Builder represents a transactions builder
type Builder interface {
    Create() Builder
    WithPayments(payments []payments.Payment) Builder
    WithTransfers(transfers transfers.Transfers) Builder
    Now() (Transactions, error)
}

// Transactions represents the transactions
type Transactions interface {
	Payments() []payments.Payment
	Transfers() transfers.Transfers
}
