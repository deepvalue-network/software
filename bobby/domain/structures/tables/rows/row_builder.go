package rows

import (
	"errors"

	"github.com/steve-care-software/products/bobby/domain/resources"
	"github.com/steve-care-software/products/bobby/domain/structures/tables"
	"github.com/steve-care-software/products/bobby/domain/structures/tables/elements"
	"github.com/steve-care-software/products/libs/hash"
)

type rowBuilder struct {
	hashAdapter     hash.Adapter
	resourceBuilder resources.Builder
	elementsBuilder elements.Builder
	elements        []elements.Element
	table           tables.Table
}

func createRowBuilder(
	hashAdapter hash.Adapter,
	resourceBuilder resources.Builder,
	elementsBuilder elements.Builder,
) RowBuilder {
	out := rowBuilder{
		hashAdapter:     hashAdapter,
		resourceBuilder: resourceBuilder,
		elementsBuilder: elementsBuilder,
		elements:        nil,
		table:           nil,
	}

	return &out
}

// Create initializes the builder
func (app *rowBuilder) Create() RowBuilder {
	return createRowBuilder(
		app.hashAdapter,
		app.resourceBuilder,
		app.elementsBuilder,
	)
}

// WithElements add elements to the builder
func (app *rowBuilder) WithElements(elements []elements.Element) RowBuilder {
	app.elements = elements
	return app
}

// OnTable adds a table to the builder to the builder
func (app *rowBuilder) OnTable(table tables.Table) RowBuilder {
	app.table = table
	return app
}

// Now builds a new Row instance
func (app *rowBuilder) Now() (Row, error) {
	if app.elements == nil {
		return nil, errors.New("the elements are mandatory in order to build an Elements instance")
	}

	if app.table == nil {
		return nil, errors.New("the table is mandatory in order to build an Elements instance")
	}

	elements, err := app.elementsBuilder.Create().WithElements(app.elements).Now()
	if err != nil {
		return nil, err
	}

	properties := app.table.Schema().Properties()
	err = elements.Fits(properties)
	if err != nil {
		return nil, err
	}

	hash, err := app.hashAdapter.FromMultiBytes([][]byte{
		elements.Hash().Bytes(),
		app.table.Resource().Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	prpertiesResource := properties.Resource()
	resource, err := app.resourceBuilder.Create().WithHash(*hash).WithAccessible(prpertiesResource).Now()
	if err != nil {
		return nil, err
	}

	return createRow(resource, elements, app.table), nil
}
