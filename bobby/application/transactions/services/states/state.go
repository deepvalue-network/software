package states

import (
	"github.com/deepvalue-network/software/bobby/domain/transactions"
)

type state struct {
	trxBuilder TransactionBuilder
	queue      []Transaction
}

func createState(
	trxBuilder TransactionBuilder,
) State {
	out := state{
		trxBuilder: trxBuilder,
		queue:      []Transaction{},
	}

	return &out
}

// Queue returns the transactions queue
func (app *state) Queue() []Transaction {
	return app.queue
}

// Single adds a single transaction to the state
func (app *state) Single(trx transactions.Transaction, callBack CallBackFn) error {
	ins, err := app.trxBuilder.Create().WithTransaction(trx).WithCallBack(callBack).Now()
	if err != nil {
		return err
	}

	app.queue = append(app.queue, ins)
	return nil
}

// List adds a list of transactions to the state
func (app *state) List(trans []transactions.Transaction, callBack CallBackFn) error {
	ins, err := app.trxBuilder.Create().WithTransactions(trans).WithCallBack(callBack).Now()
	if err != nil {
		return err
	}

	app.queue = append(app.queue, ins)
	return nil
}

// Atomic adds atomic transactions to the state
func (app *state) Atomic(atomicTrx transactions.Transactions, callBack CallBackFn) error {
	trans := atomicTrx.All()
	ins, err := app.trxBuilder.Create().WithTransactions(trans).WithCallBack(callBack).IsAtomic().Now()
	if err != nil {
		return err
	}

	app.queue = append(app.queue, ins)
	return nil
}
