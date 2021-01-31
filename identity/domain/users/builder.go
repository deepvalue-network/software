package users

import (
	"errors"

	"github.com/steve-care-software/products/identity/domain/accesses"
)

type builder struct {
	accessesFactory accesses.Factory
	name            string
	seed            string
}

func createBuilder(accessesFactory accesses.Factory) Builder {
	out := builder{
		accessesFactory: accessesFactory,
		name:            "",
		seed:            "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.accessesFactory)
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// WithSeed adds a seed to the builder
func (app *builder) WithSeed(seed string) Builder {
	app.seed = seed
	return app
}

// Now builds a new User instance
func (app *builder) Now() (User, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a User instance")
	}

	if app.seed == "" {
		return nil, errors.New("the seed is mandatory in order to build a User instance")
	}

	accesses := app.accessesFactory.Create()
	return createUser(app.name, app.seed, accesses), nil
}
