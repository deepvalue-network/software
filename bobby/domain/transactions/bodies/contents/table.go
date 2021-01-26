package contents

import (
	"github.com/steve-care-software/products/bobby/domain/selectors"
	"github.com/steve-care-software/products/bobby/domain/structures/tables/rows"
)

type table struct {
	selector selectors.Selector
	rows     rows.Rows
}

func createTable(
	selector selectors.Selector,
	rows rows.Rows,
) Table {
	out := table{
		selector: selector,
		rows:     rows,
	}

	return &out
}

// Table returns the table selector
func (obj *table) Table() selectors.Selector {
	return obj.selector
}

// Rows returns the table rows
func (obj *table) Rows() rows.Rows {
	return obj.rows
}
