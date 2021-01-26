package states

import (
	"github.com/steve-care-software/products/bobby/domain/states/overviews"
	"github.com/steve-care-software/products/bobby/domain/transactions"
)

// CallBackFn represents a callback, executed when the transaction is processed
type CallBackFn func(result Result)

// Factory represents a state factory
type Factory interface {
	Create() (State, error)
}

// State represents a transaction application state
type State interface {
	Queue() []Transaction
	Single(trx transactions.Transaction, callBack CallBackFn) error
	List(trans []transactions.Transaction, callBack CallBackFn) error
	Atomic(atomicTrx transactions.Transactions, callBack CallBackFn) error
}

// TransactionBuilder represents a transaction builder
type TransactionBuilder interface {
	Create() TransactionBuilder
	WithTransaction(trx transactions.Transaction) TransactionBuilder
	WithTransactions(trans []transactions.Transaction) TransactionBuilder
	WithCallBack(callBack CallBackFn) TransactionBuilder
	IsAtomic() TransactionBuilder
	Now() (Transaction, error)
}

// Transaction represents a state transaction
type Transaction interface {
	Transactions() []transactions.Transaction
	CallBack() CallBackFn
	IsAtomic() bool
}

// ResultBuilder represents a result builder
type ResultBuilder interface {
	Create() ResultBuilder
	WithValidTransaction(validTrx overviews.ValidTransaction) ResultBuilder
	WithInvalidTransaction(invalidTrx overviews.InvalidTransaction) ResultBuilder
	Now() (Result, error)
}

// Result represents the result of a transaction
type Result interface {
	IsValid() bool
	Valid() overviews.ValidTransaction
	IsInvalid() bool
	Invalid() overviews.InvalidTransaction
}
