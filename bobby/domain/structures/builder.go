package structures

import (
	"errors"
	"time"

	"github.com/deepvalue-network/software/bobby/domain/structures/graphbases"
	"github.com/deepvalue-network/software/bobby/domain/structures/identities"
	"github.com/deepvalue-network/software/bobby/domain/structures/sets"
	set_schemas "github.com/deepvalue-network/software/bobby/domain/structures/sets/schemas"
	"github.com/deepvalue-network/software/bobby/domain/structures/tables"
	"github.com/deepvalue-network/software/bobby/domain/structures/tables/elements"
	"github.com/deepvalue-network/software/bobby/domain/structures/tables/rows"
	table_schemas "github.com/deepvalue-network/software/bobby/domain/structures/tables/schemas"
	"github.com/deepvalue-network/software/bobby/domain/structures/tables/values"
)

type builder struct {
	graphbase             graphbases.Graphbase
	identity              identities.Identity
	setSchema             set_schemas.Schema
	set                   sets.Set
	tableSchemaValue      values.Value
	tableSchemaProperty   table_schemas.Property
	tableSchemaProperties table_schemas.Properties
	tableSchema           table_schemas.Schema
	tableElement          elements.Element
	tableRow              rows.Row
	table                 tables.Table
	isDeleted             bool
	executesOn            *time.Time
	expiresOn             *time.Time
}

func createBuilder() Builder {
	out := builder{
		graphbase:             nil,
		identity:              nil,
		setSchema:             nil,
		set:                   nil,
		tableSchemaValue:      nil,
		tableSchemaProperty:   nil,
		tableSchemaProperties: nil,
		tableSchema:           nil,
		tableElement:          nil,
		tableRow:              nil,
		table:                 nil,
		isDeleted:             false,
		executesOn:            nil,
		expiresOn:             nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithGraphbase adds a graphbase to the builder
func (app *builder) WithGraphbase(graphbase graphbases.Graphbase) Builder {
	app.graphbase = graphbase
	return app
}

// WithIdentity adds an identity to the builder
func (app *builder) WithIdentity(identity identities.Identity) Builder {
	app.identity = identity
	return app
}

// WithSetSchema adds a setSchema to the builder
func (app *builder) WithSetSchema(setSchema set_schemas.Schema) Builder {
	app.setSchema = setSchema
	return app
}

// WithSet adds a set to the builder
func (app *builder) WithSet(set sets.Set) Builder {
	app.set = set
	return app
}

// WithTableSchemaValue adds a tableSchemaValue to the builder
func (app *builder) WithTableSchemaValue(tableSchemaValue values.Value) Builder {
	app.tableSchemaValue = tableSchemaValue
	return app
}

// WithTableSchemaProperty adds a tableSchemaProperty to the builder
func (app *builder) WithTableSchemaProperty(tableSchemaProperty table_schemas.Property) Builder {
	app.tableSchemaProperty = tableSchemaProperty
	return app
}

// WithTableSchemaProperties adds a tableSchemaProperties to the builder
func (app *builder) WithTableSchemaProperties(tableSchemaProperties table_schemas.Properties) Builder {
	app.tableSchemaProperties = tableSchemaProperties
	return app
}

// WithTableSchema adds a tableSchema to the builder
func (app *builder) WithTableSchema(tableSchema table_schemas.Schema) Builder {
	app.tableSchema = tableSchema
	return app
}

// WithTableElement adds a tableElement to the builder
func (app *builder) WithTableElement(tableElement elements.Element) Builder {
	app.tableElement = tableElement
	return app
}

// WithTableRow adds a tableRow to the builder
func (app *builder) WithTableRow(tableRow rows.Row) Builder {
	app.tableRow = tableRow
	return app
}

// WithTable adds a table to the builder
func (app *builder) WithTable(table tables.Table) Builder {
	app.table = table
	return app
}

// ExecutesOn adds an execution time to the builder
func (app *builder) ExecutesOn(executesOn time.Time) Builder {
	app.executesOn = &executesOn
	return app
}

// ExpiresOn adds an expiration time to the builder
func (app *builder) ExpiresOn(expiresOn time.Time) Builder {
	app.expiresOn = &expiresOn
	return app
}

// IsDeleted flags the builder as deleted
func (app *builder) IsDeleted() Builder {
	app.isDeleted = true
	return app
}

// Now builds a new Structure instance
func (app *builder) Now() (Structure, error) {

	var tableSchema TableSchema
	if app.tableSchemaValue != nil || app.tableSchemaProperty != nil || app.tableSchemaProperties != nil || app.tableSchema != nil {
		if app.tableSchemaValue != nil {
			tableSchema = createTableSchemaWithValue(app.tableSchemaValue)
		}

		if app.tableSchemaProperty != nil {
			tableSchema = createTableSchemaWithProperty(app.tableSchemaProperty)
		}

		if app.tableSchemaProperties != nil {
			tableSchema = createTableSchemaWithProperties(app.tableSchemaProperties)
		}

		if app.tableSchema != nil {
			tableSchema = createTableSchemaWithTableSchema(app.tableSchema)
		}
	}

	var table Table
	if tableSchema != nil || app.tableElement != nil || app.tableRow != nil || app.table != nil {
		if tableSchema != nil {
			table = createTableWithSchema(tableSchema)
		}

		if app.tableElement != nil {
			table = createTableWithElement(app.tableElement)
		}

		if app.tableRow != nil {
			table = createTableWithRow(app.tableRow)
		}

		if app.table != nil {
			table = createTableWithTable(app.table)
		}
	}

	var set Set
	if app.setSchema != nil || app.set != nil {
		if app.setSchema != nil {
			set = createSetWithSchema(app.setSchema)
		}

		if app.set != nil {
			set = createSetWithSet(app.set)
		}
	}

	var content Content
	if table != nil || set != nil || app.graphbase != nil || app.identity != nil {
		if table != nil {
			content = createContentWithTable(table)
		}

		if set != nil {
			content = createContentWithSet(set)
		}

		if app.graphbase != nil {
			content = createContentWithGraph(app.graphbase)
		}

		if app.identity != nil {
			content = createContentWithIdentity(app.identity)
		}
	}

	if content == nil {
		return nil, errors.New("the structure is invalid, its content cannot be built")
	}

	if app.executesOn != nil && app.expiresOn != nil {
		return createStructureWithExecutesOnAndExpiresOn(content, app.isDeleted, app.executesOn, app.expiresOn), nil
	}

	if app.executesOn != nil {
		return createStructureWithExecutesOn(content, app.isDeleted, app.executesOn), nil
	}

	if app.expiresOn != nil {
		return createStructureWithExpiresOn(content, app.isDeleted, app.expiresOn), nil
	}

	return createStructure(content, app.isDeleted), nil
}
