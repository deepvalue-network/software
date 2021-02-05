package schemas

import (
	"errors"

	"github.com/deepvalue-network/software/bobby/domain/resources"
)

type builder struct {
	propertiesBuilder PropertiesBuilder
	resource          resources.Accessible
	name              string
	properties        []Property
}

func createBuilder(
	propertiesBuilder PropertiesBuilder,
) Builder {
	out := builder{
		propertiesBuilder: propertiesBuilder,
		resource:          nil,
		name:              "",
		properties:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.propertiesBuilder)
}

// WithResource adds a resource to the builder
func (app *builder) WithResource(res resources.Accessible) Builder {
	app.resource = res
	return app
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// WithProperties add properties to the builder
func (app *builder) WithProperties(properties []Property) Builder {
	app.properties = properties
	return app
}

// Now builds a new Schema instance
func (app *builder) Now() (Schema, error) {
	if app.resource == nil {
		return nil, errors.New("the resource is mandatory in order to build a Schema instance")
	}

	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Schema instance")
	}

	properties, err := app.propertiesBuilder.Create().WithProperties(app.properties).Now()
	if err != nil {
		return nil, err
	}

	return createSchema(app.resource, app.name, properties), nil
}
