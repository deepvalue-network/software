package selectors

import (
	"errors"

	"github.com/steve-care-software/products/bobby/domain/selectors/specifiers"
	"github.com/steve-care-software/products/libs/hash"
)

type graphbaseBuilder struct {
	hashAdapter hash.Adapter
	parent      specifiers.Specifier
	specifier   specifiers.Specifier
	metaData    Table
}

func createGraphbaseBuilder(
	hashAdapter hash.Adapter,
) GraphbaseBuilder {
	out := graphbaseBuilder{
		hashAdapter: hashAdapter,
		parent:      nil,
		specifier:   nil,
		metaData:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *graphbaseBuilder) Create() GraphbaseBuilder {
	return createGraphbaseBuilder(app.hashAdapter)
}

// WithParent adds a parent to the builder
func (app *graphbaseBuilder) WithParent(parent specifiers.Specifier) GraphbaseBuilder {
	app.parent = parent
	return app
}

// WithSpecifier adds a specifier to the builder
func (app *graphbaseBuilder) WithSpecifier(specifier specifiers.Specifier) GraphbaseBuilder {
	app.specifier = specifier
	return app
}

// WithMetaData adds a metaData to the builder
func (app *graphbaseBuilder) WithMetaData(metaData Table) GraphbaseBuilder {
	app.metaData = metaData
	return app
}

// Now builds a new Graphbase instance
func (app *graphbaseBuilder) Now() (Graphbase, error) {

	data := [][]byte{}

	var content GraphbaseContent
	if app.specifier != nil {
		data = append(data, app.specifier.Hash().Bytes())
		content = createGraphbaseContentWithSpecifier(app.specifier)
	}

	if app.metaData != nil {
		data = append(data, app.metaData.Hash().Bytes())
		content = createGraphbaseContentWithMetaData(app.metaData)
	}

	if content == nil {
		return nil, errors.New("the content (specifier, metaData) is mandatory in order to build a Graphbase instance")
	}

	if app.parent != nil {
		data = append(data, app.parent.Hash().Bytes())
	}

	hsh, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.parent != nil {
		return createGraphbaseWithParent(*hsh, content, app.parent), nil
	}

	return createGraphbase(*hsh, content), nil

}
