package mined

import (
	"errors"
	"strconv"
	"time"

	"github.com/deepvalue-network/software/blockchain/domain/links"
	"github.com/deepvalue-network/software/libs/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	link        links.Link
	results     string
	createdOn   *time.Time
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		link:        nil,
		results:     "",
		createdOn:   nil,
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

// CreatedOn adds a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.createdOn = &createdOn
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

	if app.createdOn == nil {
		createdOn := time.Now().UTC()
		app.createdOn = &createdOn
	}

	hash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.link.Hash().Bytes(),
		[]byte(app.results),
		[]byte(strconv.Itoa(app.createdOn.Second())),
	})

	if err != nil {
		return nil, err
	}

	return createLink(*hash, app.link, app.results, *app.createdOn), nil
}
