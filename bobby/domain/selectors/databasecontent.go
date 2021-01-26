package selectors

import "github.com/steve-care-software/products/bobby/domain/selectors/specifiers"

type databaseContent struct {
	specifier specifiers.Specifier
	name      string
	names     []string
}

func createDatabaseContentWithSpecifier(
	specifier specifiers.Specifier,
) DatabaseContent {
	return createDatabaseContentInternally(specifier, "", nil)
}

func createDatabaseContentWithName(
	name string,
) DatabaseContent {
	return createDatabaseContentInternally(nil, name, nil)
}

func createDatabaseContentWithNames(
	names []string,
) DatabaseContent {
	return createDatabaseContentInternally(nil, "", names)
}

func createDatabaseContentInternally(
	specifier specifiers.Specifier,
	name string,
	names []string,
) DatabaseContent {
	out := databaseContent{
		specifier: specifier,
		name:      name,
		names:     names,
	}

	return &out
}

// IsSpecifier returns true if there is a specifier, false otherwise
func (obj *databaseContent) IsSpecifier() bool {
	return obj.specifier != nil
}

// Specifier returns the specifier, if any
func (obj *databaseContent) Specifier() specifiers.Specifier {
	return obj.specifier
}

// IsName returns true if there is a name, false otherwise
func (obj *databaseContent) IsName() bool {
	return obj.name != ""
}

// Name returns the name, if any
func (obj *databaseContent) Name() string {
	return obj.name
}

// IsNames returns true if there is names, false otherwise
func (obj *databaseContent) IsNames() bool {
	return obj.names != nil
}

// Names returns the names, if any
func (obj *databaseContent) Names() []string {
	return obj.names
}
