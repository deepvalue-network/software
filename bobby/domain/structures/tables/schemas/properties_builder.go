package schemas

import (
	"errors"

	"github.com/deepvalue-network/software/bobby/domain/resources"
)

type propertiesBuilder struct {
	resource resources.Accessible
	list     []Property
}

func createPropertiesBuilder() PropertiesBuilder {
	out := propertiesBuilder{
		resource: nil,
		list:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *propertiesBuilder) Create() PropertiesBuilder {
	return createPropertiesBuilder()
}

// WithResource adds a resource to the builder
func (app *propertiesBuilder) WithResource(res resources.Accessible) PropertiesBuilder {
	app.resource = res
	return app
}

// WithProperties add properties to the builder
func (app *propertiesBuilder) WithProperties(properties []Property) PropertiesBuilder {
	app.list = properties
	return app
}

// Now builds a new Properties instance
func (app *propertiesBuilder) Now() (Properties, error) {
	if app.resource == nil {
		return nil, errors.New("the resource is mandatory in order to build a Properties instance")
	}

	if app.list == nil {
		app.list = []Property{}
	}

	return createProperties(app.resource, app.list), nil
}
