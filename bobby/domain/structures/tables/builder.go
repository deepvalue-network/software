package tables

import (
	"errors"
	"fmt"

	"github.com/deepvalue-network/software/blockchain/domain/chains"
	"github.com/deepvalue-network/software/bobby/domain/resources"
	"github.com/deepvalue-network/software/bobby/domain/structures/graphbases"
	"github.com/deepvalue-network/software/bobby/domain/structures/tables/schemas"
	"github.com/deepvalue-network/software/libs/hash"
)

type builder struct {
	hashAdapter     hash.Adapter
	resourceBuilder resources.Builder
	schema          schemas.Schema
	graphbase       graphbases.Graphbase
	chain           chains.Chain
}

func createBuilder(
	hashAdapter hash.Adapter,
	resourceBuilder resources.Builder,
) Builder {
	out := builder{
		hashAdapter:     hashAdapter,
		resourceBuilder: resourceBuilder,
		schema:          nil,
		graphbase:       nil,
		chain:           nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter, app.resourceBuilder)
}

// WithSchema adds a schema to the builder
func (app *builder) WithSchema(schema schemas.Schema) Builder {
	app.schema = schema
	return app
}

// OnGraphbase adds a graphbase to the builder
func (app *builder) OnGraphbase(graphbase graphbases.Graphbase) Builder {
	app.graphbase = graphbase
	return app
}

// OnChain adds a chain to the builder
func (app *builder) OnChain(chain chains.Chain) Builder {
	app.chain = chain
	return app
}

// Now builds a new Table instance
func (app *builder) Now() (Table, error) {
	if app.schema == nil {
		return nil, errors.New("the schema is mandatory in order to build a Table instance")
	}

	if app.graphbase == nil {
		return nil, errors.New("the graphbase is mandatory in order to build a Table instance")
	}

	data := [][]byte{
		app.schema.Resource().Hash().Bytes(),
		app.graphbase.Resource().Hash().Bytes(),
	}

	if app.chain != nil {
		data = append(data, app.chain.ID().Bytes())
	}

	hash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	schemaResource := app.schema.Resource()
	resource, err := app.resourceBuilder.Create().WithHash(*hash).WithAccessible(schemaResource).Now()
	if err != nil {
		return nil, err
	}

	if !resource.IsCompatible(app.graphbase.Resource()) {
		str := fmt.Sprintf("the table schema (ID: %s) is not compatible with the given Graphbase (ID: %s)", schemaResource.ID().String(), app.graphbase.Resource().ID().String())
		return nil, errors.New(str)
	}

	if app.chain != nil {
		return createTableWithChain(resource, app.graphbase, app.schema, app.chain), nil
	}

	return createTable(resource, app.graphbase, app.schema), nil
}
