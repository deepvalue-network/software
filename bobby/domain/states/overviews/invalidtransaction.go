package overviews

import (
	domain_errors "github.com/deepvalue-network/software/bobby/domain/errors"
	"github.com/deepvalue-network/software/bobby/domain/transactions"
)

type invalidTransaction struct {
	trx   transactions.Transaction
	error domain_errors.Error
}

func createInvalidTransaction(
	trx transactions.Transaction,
	error domain_errors.Error,
) InvalidTransaction {
	out := invalidTransaction{
		trx:   trx,
		error: error,
	}

	return &out
}

// Transaction returns the transaction
func (obj *invalidTransaction) Transaction() transactions.Transaction {
	return obj.trx
}

// Error returns the error
func (obj *invalidTransaction) Error() domain_errors.Error {
	return obj.error
}
