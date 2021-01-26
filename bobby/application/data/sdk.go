package data

// Application represents the data application
type Application interface {
	Graphbase() Graphbase
	Database() Database
	Table() Table
	Set() Set
}

// Graphbase represents the graphbase data application
type Graphbase interface {
}

// Database represents the database data application
type Database interface {
}

// Table represents the table data application
type Table interface {
}

// Set represents the set data application
type Set interface {
}
