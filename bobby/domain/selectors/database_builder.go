package selectors

import (
	"errors"

	"github.com/deepvalue-network/software/bobby/domain/selectors/specifiers"
	"github.com/deepvalue-network/software/libs/hash"
)

type databaseBuilder struct {
	hashAdapter hash.Adapter
	graphbase   specifiers.Specifier
	specifier   specifiers.Specifier
	name        string
	names       []string
}

func createDatabaseBuilder(
	hashAdapter hash.Adapter,
) DatabaseBuilder {
	out := databaseBuilder{
		hashAdapter: hashAdapter,
		graphbase:   nil,
		specifier:   nil,
		name:        "",
		names:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *databaseBuilder) Create() DatabaseBuilder {
	return createDatabaseBuilder(app.hashAdapter)
}

// WithGraphbase adds a graphbase to the builder
func (app *databaseBuilder) WithGraphbase(graphbase specifiers.Specifier) DatabaseBuilder {
	app.graphbase = graphbase
	return app
}

// WithSpecifier adds a specifier to the builder
func (app *databaseBuilder) WithSpecifier(specifier specifiers.Specifier) DatabaseBuilder {
	app.specifier = specifier
	return app
}

// WithName adds a name to the builder
func (app *databaseBuilder) WithName(name string) DatabaseBuilder {
	app.name = name
	return app
}

// WithNames adds names to the builder
func (app *databaseBuilder) WithNames(names []string) DatabaseBuilder {
	app.names = names
	return app
}

// Now builds a new Database instance
func (app *databaseBuilder) Now() (Database, error) {
	if app.graphbase == nil {
		return nil, errors.New("the graphbase is mandatory in order to build a Database instance")
	}

	data := [][]byte{
		app.graphbase.Hash().Bytes(),
	}

	var content DatabaseContent
	if app.specifier != nil {
		data = append(data, app.specifier.Hash().Bytes())
		content = createDatabaseContentWithSpecifier(app.specifier)
	}

	if app.name != "" {
		data = append(data, []byte(app.name))
		content = createDatabaseContentWithName(app.name)
	}

	if app.names != nil {
		names := [][]byte{}
		for _, oneName := range app.names {
			names = append(names, []byte(oneName))
		}

		data = append(data, names...)
		content = createDatabaseContentWithNames(app.names)
	}

	if content == nil {
		return nil, errors.New("the content (specifier, name, names) is mandatory in order to build a Database instance")
	}

	hsh, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createDatabase(*hsh, app.graphbase, content), nil
}
