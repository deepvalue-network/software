package mines

import (
	"errors"
	"strconv"
	"time"

	"github.com/deepvalue-network/software/blockchain/domain/chains"
	"github.com/deepvalue-network/software/libs/hash"
	"github.com/deepvalue-network/software/libs/hashtree"
)

type builder struct {
	hashAdapter hash.Adapter
	chain       chains.Chain
	diamonds    hashtree.HashTree
	createdOn   *time.Time
}

func createBuilder(hashAdapter hash.Adapter) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		chain:       nil,
		diamonds:    nil,
		createdOn:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithChain adds a chain to the builder
func (app *builder) WithChain(chain chains.Chain) Builder {
	app.chain = chain
	return app
}

// WithDiamonds adds diamonds to the builder
func (app *builder) WithDiamonds(diamonds hashtree.HashTree) Builder {
	app.diamonds = diamonds
	return app
}

// CreatedOn adds a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new Mine instance
func (app *builder) Now() (Mine, error) {
	if app.chain == nil {
		return nil, errors.New("the chain is mandatory in order to build a Mine instance")
	}

	if app.diamonds == nil {
		return nil, errors.New("the diamonds hashtree is mandatory in order to build a Mine instance")
	}

	if app.createdOn == nil {
		createdOn := time.Now().UTC()
		app.createdOn = &createdOn
	}

	hsh, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.chain.ID().Bytes(),
		app.diamonds.Head().Bytes(),
		[]byte(strconv.Itoa(int(app.createdOn.UnixNano()))),
	})

	if err != nil {
		return nil, err
	}

	return createMine(*hsh, app.chain, app.diamonds, *app.createdOn), nil
}
