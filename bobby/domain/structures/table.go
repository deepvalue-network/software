package structures

import (
	"github.com/steve-care-software/products/bobby/domain/structures/tables"
	"github.com/steve-care-software/products/bobby/domain/structures/tables/elements"
	"github.com/steve-care-software/products/bobby/domain/structures/tables/rows"
)

type table struct {
	schema  TableSchema
	element elements.Element
	row     rows.Row
	tb      tables.Table
}

func createTableWithSchema(
	schema TableSchema,
) Table {
	return createTableInternally(schema, nil, nil, nil)
}

func createTableWithElement(
	element elements.Element,
) Table {
	return createTableInternally(nil, element, nil, nil)
}

func createTableWithRow(
	row rows.Row,
) Table {
	return createTableInternally(nil, nil, row, nil)
}

func createTableWithTable(
	tb tables.Table,
) Table {
	return createTableInternally(nil, nil, nil, tb)
}

func createTableInternally(
	schema TableSchema,
	element elements.Element,
	row rows.Row,
	tb tables.Table,
) Table {
	out := table{
		schema:  schema,
		element: element,
		row:     row,
		tb:      tb,
	}

	return &out
}

// IsSchema returns true if there is a schema, false otherwise
func (obj *table) IsSchema() bool {
	return obj.schema != nil
}

// Schema returns the schema, if any
func (obj *table) Schema() TableSchema {
	return obj.schema
}

// IsElement returns true if there is an element, false otherwise
func (obj *table) IsElement() bool {
	return obj.element != nil
}

// Element returns the element, if any
func (obj *table) Element() elements.Element {
	return obj.element
}

// IsRow returns true if there is a row, false otherwise
func (obj *table) IsRow() bool {
	return obj.row != nil
}

// Row returns the row, if any
func (obj *table) Row() rows.Row {
	return obj.row
}

// IsTable returns true if there is a table, false otherwise
func (obj *table) IsTable() bool {
	return obj.tb != nil
}

// Table returns the table, if any
func (obj *table) Table() tables.Table {
	return obj.tb
}
