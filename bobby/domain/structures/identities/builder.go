package identities

import (
	"errors"

	"github.com/deepvalue-network/software/bobby/domain/resources"
	"github.com/deepvalue-network/software/bobby/domain/structures/graphbases"
	"github.com/deepvalue-network/software/libs/hash"
)

type builder struct {
	resource  resources.Mutable
	graphbase graphbases.Graphbase
	key       *hash.Hash
	name      string
}

func createBuilder() Builder {
	out := builder{
		resource:  nil,
		graphbase: nil,
		key:       nil,
		name:      "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithResource adds a resource to the builder
func (app *builder) WithResource(res resources.Mutable) Builder {
	app.resource = res
	return app
}

// WithKey adds a key to the builder
func (app *builder) WithKey(key hash.Hash) Builder {
	app.key = &key
	return app
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// OnGraphbase adds a graphbase to the builder
func (app *builder) OnGraphbase(graphbase graphbases.Graphbase) Builder {
	app.graphbase = graphbase
	return app
}

// Now builds a new Identity instance
func (app *builder) Now() (Identity, error) {
	if app.resource == nil {
		return nil, errors.New("the resource is mandatory in order to build an Identity instance")
	}

	if app.graphbase == nil {
		return nil, errors.New("the graphbase is mandatory in order to build an Identity instance")
	}

	if app.key == nil {
		return nil, errors.New("the key is mandatory in order to build an Identity instance")
	}

	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build an Identity instance")
	}

	return createIdentity(app.resource, app.graphbase, *app.key, app.name), nil
}
