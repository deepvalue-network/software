package graphbases

import (
	"errors"

	"github.com/steve-care-software/products/blockchain/domain/chains"
	"github.com/steve-care-software/products/bobby/domain/resources"
)

type builder struct {
	resource resources.Accessible
	metaData string
	chain    chains.Chain
	parent   resources.Accessible
}

func createBuilder() Builder {
	out := builder{
		resource: nil,
		metaData: "",
		chain:    nil,
		parent:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithResource adds a resource to the builder
func (app *builder) WithResource(res resources.Accessible) Builder {
	app.resource = res
	return app
}

// WithMetaData adds a metaData to the builder
func (app *builder) WithMetaData(metaData string) Builder {
	app.metaData = metaData
	return app
}

// WithParent adds a parent to the builder
func (app *builder) WithParent(parent resources.Accessible) Builder {
	app.parent = parent
	return app
}

// OnChain adds a chain to the builder
func (app *builder) OnChain(chain chains.Chain) Builder {
	app.chain = chain
	return app
}

// Now builds a new Graphbase instance
func (app *builder) Now() (Graphbase, error) {
	if app.resource == nil {
		return nil, errors.New("the resource is mandatory in order to build a Graphbase instance")
	}

	if app.metaData == "" {
		return nil, errors.New("the metaData is mandatory in order to build a Graphbase instance")
	}

	if app.chain == nil {
		return nil, errors.New("the chain is mandatory in order to build a Graphbase instance")
	}

	if app.parent != nil {
		return createGraphbaseWithParent(app.resource, app.metaData, app.chain, app.parent), nil
	}

	return createGraphbase(app.resource, app.metaData, app.chain), nil
}
