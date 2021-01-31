package structures

import (
	table_schemas "github.com/steve-care-software/products/bobby/domain/structures/tables/schemas"
	"github.com/steve-care-software/products/bobby/domain/structures/tables/values"
)

type tableSchema struct {
	val        values.Value
	property   table_schemas.Property
	properties table_schemas.Properties
	schema     table_schemas.Schema
}

func createTableSchemaWithValue(
	val values.Value,
) TableSchema {
	return createTableSchemaInternally(val, nil, nil, nil)
}

func createTableSchemaWithProperty(
	property table_schemas.Property,
) TableSchema {
	return createTableSchemaInternally(nil, property, nil, nil)
}

func createTableSchemaWithProperties(
	properties table_schemas.Properties,
) TableSchema {
	return createTableSchemaInternally(nil, nil, properties, nil)
}

func createTableSchemaWithTableSchema(
	schema table_schemas.Schema,
) TableSchema {
	return createTableSchemaInternally(nil, nil, nil, schema)
}

func createTableSchemaInternally(
	val values.Value,
	property table_schemas.Property,
	properties table_schemas.Properties,
	schema table_schemas.Schema,
) TableSchema {
	out := tableSchema{
		val:        val,
		property:   property,
		properties: properties,
		schema:     schema,
	}

	return &out
}

// IsValue returns true if there is a value, false otherwise
func (obj *tableSchema) IsValue() bool {
	return obj.val != nil
}

// Value returns the value, if any
func (obj *tableSchema) Value() values.Value {
	return obj.val
}

// IsProperty returns true if there is a property, false otherwise
func (obj *tableSchema) IsProperty() bool {
	return obj.property != nil
}

// Property returns the property, if any
func (obj *tableSchema) Property() table_schemas.Property {
	return obj.property
}

// IsProperties returns true if there is a properties, false otherwise
func (obj *tableSchema) IsProperties() bool {
	return obj.properties != nil
}

// Properties returns the properties, if any
func (obj *tableSchema) Properties() table_schemas.Properties {
	return obj.properties
}

// IsSchema returns true if there is a schema, false otherwise
func (obj *tableSchema) IsSchema() bool {
	return obj.schema != nil
}

// Schema returns the schema, if any
func (obj *tableSchema) Schema() table_schemas.Schema {
	return obj.schema
}
