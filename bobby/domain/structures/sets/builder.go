package sets

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/products/bobby/domain/resources"
	"github.com/steve-care-software/products/bobby/domain/structures/graphbases"
	"github.com/steve-care-software/products/bobby/domain/structures/sets/schemas"
	"github.com/steve-care-software/products/libs/hash"
)

type builder struct {
	hashAdapter     hash.Adapter
	resourceBuilder resources.Builder
	graphbase       graphbases.Graphbase
	schema          schemas.Schema
	elements        Elements
	name            string
}

func createBuilder(
	hashAdapter hash.Adapter,
	resourceBuilder resources.Builder,
) Builder {
	out := builder{
		hashAdapter:     hashAdapter,
		resourceBuilder: resourceBuilder,
		graphbase:       nil,
		schema:          nil,
		elements:        nil,
		name:            "",
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

// WithElements add elements to the builder
func (app *builder) WithElements(elements Elements) Builder {
	app.elements = elements
	return app
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// OnGraphbase adds a graphbase to the builder
func (app *builder) OnGraphbase(graphbase graphbases.Graphbase) Builder {
	app.graphbase = graphbase
	return app
}

// Now builds a new Set instance
func (app *builder) Now() (Set, error) {
	if app.graphbase == nil {
		return nil, errors.New("the graphbase is mandatory in order to build a Set instance")
	}

	if app.schema == nil {
		return nil, errors.New("the schema is mandatory in order to build a Set instance")
	}

	if app.elements == nil {
		return nil, errors.New("the elements are mandatory in order to build a Set instance")
	}

	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Set instance")
	}

	if app.schema.IsUniqueElements() {
		if !app.elements.IsUnique() {
			return nil, errors.New("the schema was expecting unique elements, but the given set elements are not unique")
		}
	}

	schemaResource := app.schema.Resource()
	hash, err := app.hashAdapter.FromMultiBytes([][]byte{
		schemaResource.Hash().Bytes(),
		app.elements.Hash().Bytes(),
		[]byte(app.name),
	})

	if err != nil {
		return nil, err
	}

	resource, err := app.resourceBuilder.Create().WithHash(*hash).WithAccessible(schemaResource).Now()
	if err != nil {
		return nil, err
	}

	if !resource.IsCompatible(app.graphbase.Resource()) {
		str := fmt.Sprintf("the set schema (ID: %s) is not compatible with the given Graphbase (ID: %s)", schemaResource.ID().String(), app.graphbase.Resource().ID().String())
		return nil, errors.New(str)
	}

	return createSet(resource, app.graphbase, app.schema, app.elements, app.name), nil
}
