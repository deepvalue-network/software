package states

import (
	"errors"

	"github.com/steve-care-software/products/blockchain/application/services"
	"github.com/steve-care-software/products/bobby/domain/resources"
	"github.com/steve-care-software/products/bobby/domain/transactions"
	"github.com/steve-care-software/products/libs/hash"
)

type builder struct {
	hashAdapter      hash.Adapter
	immutableBuilder resources.ImmutableBuilder
	blockAppService  services.Block
	transList        []transactions.Transactions
	prev             *hash.Hash
}

func createBuilder(
	hashAdapter hash.Adapter,
	immutableBuilder resources.ImmutableBuilder,
	blockAppService services.Block,
) Builder {
	out := builder{
		hashAdapter:      hashAdapter,
		immutableBuilder: immutableBuilder,
		blockAppService:  blockAppService,
		transList:        nil,
		prev:             nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
		app.immutableBuilder,
		app.blockAppService,
	)
}

// WithTransactionsList adds a transactions list to the builder
func (app *builder) WithTransactionsList(transList []transactions.Transactions) Builder {
	app.transList = transList
	return app
}

// WithPrevious adds a previous hash to the builder
func (app *builder) WithPrevious(prev hash.Hash) Builder {
	app.prev = &prev
	return app
}

// Now builds a new State instance
func (app *builder) Now() (State, error) {

	if app.transList != nil && len(app.transList) <= 0 {
		app.transList = nil
	}

	if app.transList == nil {
		return nil, errors.New("the transactions list is mandatory in order to build a State instance")
	}

	hashes := []hash.Hash{}
	for _, oneTrans := range app.transList {
		hashes = append(hashes, oneTrans.Hash())
	}

	block, err := app.blockAppService.Create(hashes)
	if err != nil {
		return nil, err
	}

	data := [][]byte{
		block.Tree().Head().Bytes(),
	}

	if app.prev != nil {
		data = append(data, app.prev.Bytes())
	}

	hash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	resource, err := app.immutableBuilder.Create().WithHash(*hash).Now()
	if err != nil {
		return nil, err
	}

	if app.prev != nil {
		return createStateWithPrevious(resource, app.transList, block, app.prev), nil
	}

	return createState(resource, app.transList, block), nil
}
