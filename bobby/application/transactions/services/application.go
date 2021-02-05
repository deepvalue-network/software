package services

import (
	"errors"

	states_app "github.com/deepvalue-network/software/bobby/application/transactions/services/states"
	"github.com/deepvalue-network/software/bobby/domain/states"
	"github.com/deepvalue-network/software/bobby/domain/states/overviews"
	"github.com/deepvalue-network/software/bobby/domain/transactions"
	"github.com/deepvalue-network/software/libs/hash"
)

type application struct {
	transactionsBuilder transactions.Builder
	stateBuilder        states.Builder
	stateService        states.Service
	stateAppFactory     states_app.Factory
	resultBuilder       states_app.ResultBuilder
	stateHash           *hash.Hash
	queue               []states_app.State
}

func createApplication(
	transactionsBuilder transactions.Builder,
	stateBuilder states.Builder,
	stateService states.Service,
	stateAppFactory states_app.Factory,
	resultBuilder states_app.ResultBuilder,
) Application {
	return createApplicationInternally(
		transactionsBuilder,
		stateBuilder,
		stateService,
		stateAppFactory,
		resultBuilder,
		nil,
	)
}

func createApplicationWithStateHash(
	transactionsBuilder transactions.Builder,
	stateBuilder states.Builder,
	stateService states.Service,
	stateAppFactory states_app.Factory,
	resultBuilder states_app.ResultBuilder,
	stateHash *hash.Hash,
) Application {
	return createApplicationInternally(
		transactionsBuilder,
		stateBuilder,
		stateService,
		stateAppFactory,
		resultBuilder,
		stateHash,
	)
}

func createApplicationInternally(
	transactionsBuilder transactions.Builder,
	stateBuilder states.Builder,
	stateService states.Service,
	stateAppFactory states_app.Factory,
	resultBuilder states_app.ResultBuilder,
	stateHash *hash.Hash,
) Application {
	out := application{
		transactionsBuilder: transactionsBuilder,
		stateBuilder:        stateBuilder,
		stateService:        stateService,
		stateAppFactory:     stateAppFactory,
		resultBuilder:       resultBuilder,
		stateHash:           stateHash,
		queue:               []states_app.State{},
	}

	return &out
}

// Begin opens a state
func (app *application) Begin() (states_app.State, error) {
	return app.stateAppFactory.Create()
}

// Commit commits a state
func (app *application) Commit(state states_app.State) {
	app.queue = append(app.queue, state)
}

// Rollback rollbacks the last state
func (app *application) Rollback() error {
	if len(app.queue) <= 0 {
		return errors.New("there is no state in the queue and therefore cannot Rollback")
	}

	app.queue = app.queue[:len(app.queue)-1]
	return nil
}

// Push pushes the state
func (app *application) Push() error {
	trxCallbacks := map[string]states_app.CallBackFn{}
	transactions := []transactions.Transactions{}
	for _, oneState := range app.queue {
		queue := oneState.Queue()
		for _, oneStateTrx := range queue {
			trans := oneStateTrx.Transactions()

			builder := app.transactionsBuilder.Create().WithTransactions(trans)
			if oneStateTrx.IsAtomic() {
				builder.IsAtomic()
			}

			ins, err := builder.Now()
			if err != nil {
				return err
			}

			// add the transactions in the list:
			transactions = append(transactions, ins)

			// add the trx -> callback mapping:\
			callback := oneStateTrx.CallBack()
			for _, oneTrx := range trans {
				trxCallbacks[oneTrx.Hash().String()] = callback
			}
		}
	}

	st, overviews, err := app.prepare(transactions, false)
	if err != nil {
		return err
	}

	for _, oneOverview := range overviews {
		// execute the invalid transaction callbacks:
		invalidTrans := oneOverview.Invalid()
		err = app.executeInvalidTransactionsCallback(invalidTrans, trxCallbacks)
		if err != nil {
			return err
		}

		// execute the valid transaction callbacks:
		validTrans := oneOverview.Valid()
		err = app.executeValidTransactionsCallback(validTrans, trxCallbacks)
		if err != nil {
			return err
		}

		// if the state can't be saved:
		if !oneOverview.CanBeSaved() {
			return nil
		}
	}

	// save the state:
	hash := st.Resource().Hash()
	return app.stateService.Save(hash)
}

func (app *application) prepare(list []transactions.Transactions, isAtomic bool) (states.State, []overviews.Overview, error) {
	builder := app.stateBuilder.Create().WithTransactionsList(list)
	if app.stateHash != nil {
		builder.WithPrevious(*app.stateHash)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	// prepare the state:
	overview, err := app.stateService.Prepare(ins)
	if err != nil {
		return nil, nil, err
	}

	return ins, overview, nil
}

func (app *application) executeValidTransactionsCallback(validTrans []overviews.ValidTransaction, trxCallbacks map[string]states_app.CallBackFn) error {
	if validTrans == nil {
		return nil
	}

	for _, oneValidTrx := range validTrans {
		keyname := oneValidTrx.Transaction().Hash().String()
		if callbackFn, ok := trxCallbacks[keyname]; ok {
			err := app.executeValidTransactionCallback(oneValidTrx, callbackFn)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (app *application) executeValidTransactionCallback(validTrx overviews.ValidTransaction, callbackFn states_app.CallBackFn) error {
	res, err := app.resultBuilder.Create().WithValidTransaction(validTrx).Now()
	if err != nil {
		return err
	}

	// call the callback:
	callbackFn(res)
	return nil
}

func (app *application) executeInvalidTransactionsCallback(invalidTrans []overviews.InvalidTransaction, trxCallbacks map[string]states_app.CallBackFn) error {
	if invalidTrans == nil {
		return nil
	}

	for _, oneInvalidTrx := range invalidTrans {
		keyname := oneInvalidTrx.Transaction().Hash().String()
		if callbackFn, ok := trxCallbacks[keyname]; ok {
			err := app.executeInvalidTransactionCallback(oneInvalidTrx, callbackFn)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (app *application) executeInvalidTransactionCallback(invalidTrx overviews.InvalidTransaction, callbackFn states_app.CallBackFn) error {
	res, err := app.resultBuilder.Create().WithInvalidTransaction(invalidTrx).Now()
	if err != nil {
		return err
	}

	// call the callback:
	callbackFn(res)
	return nil
}
