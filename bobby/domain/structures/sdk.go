package structures

import (
	uuid "github.com/satori/go.uuid"
	"github.com/steve-care-software/products/bobby/domain/selectors"
	"github.com/steve-care-software/products/bobby/domain/structures/graphbases"
	"github.com/steve-care-software/products/bobby/domain/structures/identities"
	"github.com/steve-care-software/products/bobby/domain/structures/sets"
	set_schemas "github.com/steve-care-software/products/bobby/domain/structures/sets/schemas"
	"github.com/steve-care-software/products/bobby/domain/structures/tables"
	"github.com/steve-care-software/products/bobby/domain/structures/tables/elements"
	"github.com/steve-care-software/products/bobby/domain/structures/tables/rows"
	table_schemas "github.com/steve-care-software/products/bobby/domain/structures/tables/schemas"
	"github.com/steve-care-software/products/bobby/domain/structures/tables/values"
	"github.com/steve-care-software/products/libs/hash"
)

// Builder represents a structure builder
type Builder interface {
	Create() Builder
	WithGraphbase(graphbase graphbases.Graphbase) Builder
	WithIdentity(identity identities.Identity) Builder
	WithSetSchema(setSchame set_schemas.Schema) Builder
	WithSet(set sets.Set) Builder
	WithTableSchemaValue(tableSchemaValue values.Value) Builder
	WithTableSchemaProperty(tableSchemaProperty table_schemas.Property) Builder
	WithTableSchemaProperties(tableSchemaProperties table_schemas.Properties) Builder
	WithTableSchema(tableSchema table_schemas.Schema) Builder
	WithTableElement(tableElement elements.Element) Builder
	WithTableRow(tableRow rows.Row) Builder
	WithTable(table tables.Table) Builder
	Now() (Structure, error)
}

// Structure represents a structure
type Structure interface {
	IsGraphbase() bool
	Graphbase() graphbases.Graphbase
	IsIdentity() bool
	Identity() identities.Identity
	IsTable() bool
	Table() Table
	IsSet() bool
	Set() Set
}

// Set represents a set
type Set interface {
	IsSchema() bool
	Schema() set_schemas.Schema
	IsSet() bool
	Set() sets.Set
}

// Table represents a table
type Table interface {
	IsSchema() bool
	Schema() TableSchema
	IsElement() bool
	Element() elements.Element
	IsRow() bool
	Row() rows.Row
	IsTable() bool
	Table() tables.Table
}

// TableSchema represents a table schema
type TableSchema interface {
	IsValue() bool
	Value() values.Value
	IsProperty() bool
	Property() table_schemas.Property
	IsProperties() bool
	Properties() table_schemas.Properties
	IsSchema() bool
	Schema() table_schemas.Schema
}

// Repository represents a structure repository
type Repository interface {
	Retrieve(id *uuid.UUID) (Structure, error)
	RetrieveByHash(hash hash.Hash) (Structure, error)
	Search(selector selectors.Selector) ([]Structure, error)
}

// Service represents a structure service
type Service interface {
	Save(structure Structure) error
	SaveAll(list []Structure) error
}
