package selectors

type content struct {
	graphbase Graphbase
	db        Database
	table     Table
	set       Set
}

func createContentWithGraphbase(
	graphbase Graphbase,
) Content {
	return createContentInternally(graphbase, nil, nil, nil)
}

func createContentWithDatabase(
	db Database,
) Content {
	return createContentInternally(nil, db, nil, nil)
}

func createContentWithTable(
	table Table,
) Content {
	return createContentInternally(nil, nil, table, nil)
}

func createContentWithSet(
	set Set,
) Content {
	return createContentInternally(nil, nil, nil, set)
}

func createContentInternally(
	graphbase Graphbase,
	db Database,
	table Table,
	set Set,
) Content {
	out := content{
		graphbase: graphbase,
		db:        db,
		table:     table,
		set:       set,
	}

	return &out
}

// IsGraphbase returns true if there is a graphbase, false otherwise
func (obj *content) IsGraphbase() bool {
	return obj.graphbase != nil
}

// Graphbase returns the graphbase, if any
func (obj *content) Graphbase() Graphbase {
	return obj.graphbase
}

// IsDatabase returns true if there is a database, false otherwise
func (obj *content) IsDatabase() bool {
	return obj.db != nil
}

// Database returns the database, if any
func (obj *content) Database() Database {
	return obj.db
}

// IsTable returns true if there is a table, false otherwise
func (obj *content) IsTable() bool {
	return obj.table != nil
}

// Table returns the table, if any
func (obj *content) Table() Table {
	return obj.table
}

// IsSet returns true if there is a set, false otherwise
func (obj *content) IsSet() bool {
	return obj.set != nil
}

// Set returns the set, if any
func (obj *content) Set() Set {
	return obj.set
}
