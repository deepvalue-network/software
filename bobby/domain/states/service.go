package states

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/steve-care-software/products/blockchain/domain/blocks"
	"github.com/steve-care-software/products/blockchain/domain/chains"
	derrors "github.com/steve-care-software/products/bobby/domain/errors"
	"github.com/steve-care-software/products/bobby/domain/states/overviews"
	"github.com/steve-care-software/products/bobby/domain/structures"
	"github.com/steve-care-software/products/bobby/domain/transactions"
	"github.com/steve-care-software/products/libs/hash"
	"github.com/steve-care-software/products/libs/hydro"
)

type service struct {
	errorBuilder      derrors.Builder
	validTrxBuilder   overviews.ValidTransactionBuilder
	invalidTrxBuilder overviews.InvalidTransactionBuilder
	overviewBuilder   overviews.Builder
	structureService  structures.Service
	trxProc           transactions.TransactionProcessor
	hydroAdapter      hydro.Adapter
	onChain           chains.Chain
	dirPath           string
	lastStateFileName string
	filePermission    os.FileMode
	mp                map[string]*stateOverviews
}

func createService(
	errorBuilder derrors.Builder,
	validTrxBuilder overviews.ValidTransactionBuilder,
	invalidTrxBuilder overviews.InvalidTransactionBuilder,
	overviewBuilder overviews.Builder,
	structureService structures.Service,
	trxProc transactions.TransactionProcessor,
	hydroAdapter hydro.Adapter,
	onChain chains.Chain,
	dirPath string,
	lastStateFileName string,
	filePermission os.FileMode,
) Service {
	out := service{
		errorBuilder:      errorBuilder,
		validTrxBuilder:   validTrxBuilder,
		invalidTrxBuilder: invalidTrxBuilder,
		overviewBuilder:   overviewBuilder,
		structureService:  structureService,
		trxProc:           trxProc,
		hydroAdapter:      hydroAdapter,
		onChain:           onChain,
		dirPath:           dirPath,
		lastStateFileName: lastStateFileName,
		filePermission:    filePermission,
		mp:                map[string]*stateOverviews{},
	}
	return &out
}

// Prepare prepares the state instance
func (app *service) Prepare(state State) ([]overviews.Overview, error) {
	block := state.Block()
	transList := state.Transactions()
	overviews, err := app.processTransList(block, transList)
	if err != nil {
		return nil, err
	}

	keyname := state.Resource().Hash().String()
	app.mp[keyname] = createStateOverviews(state, overviews)
	return overviews, nil
}

// Save saves the state instance
func (app *service) Save(hash hash.Hash) error {
	keyname := hash.String()
	defer func() {
		delete(app.mp, keyname)
	}()

	if stateOverviews, ok := app.mp[keyname]; ok {
		structures := []structures.Structure{}
		for _, oneOverview := range stateOverviews.overviews {
			if !oneOverview.CanBeSaved() {
				continue
			}

			validTrans := oneOverview.Valid()
			for _, oneValidTrx := range validTrans {
				structures = append(structures, oneValidTrx.Structures()...)
			}
		}

		if len(structures) <= 0 {
			return nil
		}

		err := app.structureService.SaveAll(structures)
		if err != nil {
			return err
		}

		dehydrated, err := app.hydroAdapter.Dehydrate(stateOverviews.state)
		if err != nil {
			return err
		}

		js, err := json.Marshal(dehydrated)
		if err != nil {
			return err
		}

		// save the data using the hash:
		stateHashStr := stateOverviews.state.Resource().Hash().String()
		stateFilename := fmt.Sprintf("%s.json", stateHashStr)
		path := filepath.Join(app.dirPath, stateFilename)
		err = ioutil.WriteFile(path, js, app.filePermission)
		if err != nil {
			return err
		}

		// save the current state hash:
		lastStatePath := filepath.Join(app.dirPath, app.lastStateFileName)
		return ioutil.WriteFile(lastStatePath, []byte(stateHashStr), app.filePermission)
	}

	str := fmt.Sprintf("the State (ID: %s) has never been prepared and therefore cannot be saved", keyname)
	return errors.New(str)
}

func (app *service) processTransList(
	block blocks.Block,
	transList []transactions.Transactions,
) ([]overviews.Overview, error) {
	out := []overviews.Overview{}
	for _, oneTrans := range transList {
		ins, err := app.processTrans(block, oneTrans)
		if err != nil {
			return nil, err
		}

		out = append(out, ins)
	}

	return out, nil
}

func (app *service) processTrans(
	block blocks.Block,
	trans transactions.Transactions,
) (overviews.Overview, error) {
	list := trans.All()
	validTrans := []overviews.ValidTransaction{}
	invalidTrans := []overviews.InvalidTransaction{}
	for _, oneTrx := range list {
		validTrx, invalidTrx, err := app.processTrx(block, oneTrx)
		if err != nil {
			return nil, err
		}

		if validTrx != nil {
			validTrans = append(validTrans, validTrx)
		}

		if invalidTrx != nil {
			invalidTrans = append(invalidTrans, invalidTrx)
		}
	}

	builder := app.overviewBuilder.Create()
	if len(validTrans) > 0 {
		builder.WithValid(validTrans)
	}

	hasInvalidTrans := len(invalidTrans) > 0
	if hasInvalidTrans {
		builder.WithInvalid(invalidTrans)
	}

	if !(trans.IsAtomic() && hasInvalidTrans) {
		builder.CanBeSaved()
	}

	return builder.Now()

}

func (app *service) processTrx(
	block blocks.Block,
	trx transactions.Transaction,
) (overviews.ValidTransaction, overviews.InvalidTransaction, error) {
	structureList, err := app.trxProc.Execute(trx)
	if err != nil {
		str := fmt.Sprintf("there was an error while processing the transaction (hash: %s), error: %s", trx.Hash().String(), err.Error())
		errIns, err := app.errorBuilder.Create().WithMessage(str).WithCode(derrors.CannotProcessTrx).Now()
		if err != nil {
			return nil, nil, err
		}

		invalidTrx, err := app.invalidTrxBuilder.Create().WithTransaction(trx).WithError(errIns).Now()
		if err != nil {
			return nil, nil, err
		}

		return nil, invalidTrx, nil
	}

	validTrx, err := app.validTrxBuilder.Create().WithTransaction(trx).WithStructures(structureList).WithChain(app.onChain).WithBlock(block).Now()
	if err != nil {
		return nil, nil, err
	}

	return validTrx, nil, nil
}
