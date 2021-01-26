package schemas

import (
	"errors"

	"github.com/steve-care-software/products/bobby/domain/resources"
	"github.com/steve-care-software/products/bobby/domain/structures/tables"
)

type builder struct {
	resource         resources.Accessible
	name             string
	table            tables.Table
	isUniqueElements bool
}

func createBuilder() Builder {
	out := builder{
		resource:         nil,
		name:             "",
		table:            nil,
		isUniqueElements: false,
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

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// WithTable adds a table to the builder
func (app *builder) WithTable(table tables.Table) Builder {
	app.table = table
	return app
}

// IsUniqueElements flags the schema as unique elements
func (app *builder) IsUniqueElements() Builder {
	app.isUniqueElements = true
	return app
}

// Now builds a new Schema instance
func (app *builder) Now() (Schema, error) {
	if app.resource == nil {
		return nil, errors.New("the resource is mandatory in order to build a Schema instance")
	}

	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Schema instance")
	}

	if app.table == nil {
		return nil, errors.New("the table is mandatory in order to build a Schema instance")
	}

	return createSchema(app.resource, app.name, app.table, app.isUniqueElements), nil
}
