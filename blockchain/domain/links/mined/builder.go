package mined

import (
	"errors"
	"strconv"
	"time"

	"github.com/steve-care-software/products/blockchain/domain/links"
	"github.com/steve-care-software/products/libs/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	link        links.Link
	results     string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		link:        nil,
		results:     "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// Create adds a link to the builder
func (app *builder) WithLink(link links.Link) Builder {
	app.link = link
	return app
}

// WithResults add results to the builder
func (app *builder) WithResults(results string) Builder {
	app.results = results
	return app
}

// Now builds a new Link instance
func (app *builder) Now() (Link, error) {
	if app.link == nil {
		return nil, errors.New("the link is mandatory in order to build a mined Link instance")
	}

	if app.results == "" {
		return nil, errors.New("the results are mandatory in order to build a mined Link instance")
	}

	createdOn := time.Now().UTC()
	hash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.link.Hash().Bytes(),
		[]byte(app.results),
		[]byte(strconv.Itoa(createdOn.Nanosecond())),
	})

	if err != nil {
		return nil, err
	}

	return createLink(*hash, app.link, app.results, createdOn), nil
}
