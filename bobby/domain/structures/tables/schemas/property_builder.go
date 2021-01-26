package schemas

import (
	"errors"

	"github.com/steve-care-software/products/bobby/domain/resources"
)

type propertyBuilder struct {
	resource     resources.Accessible
	name         string
	isPrimaryKey bool
	foreignKey   Schema
	typ          Type
}

func createPropertyBuilder() PropertyBuilder {
	out := propertyBuilder{
		resource:     nil,
		name:         "",
		isPrimaryKey: false,
		foreignKey:   nil,
		typ:          nil,
	}

	return &out
}

// Create initializes the builder
func (app *propertyBuilder) Create() PropertyBuilder {
	return createPropertyBuilder()
}

// WithResource adds a resource to the builder
func (app *propertyBuilder) WithResource(res resources.Accessible) PropertyBuilder {
	app.resource = res
	return app
}

// WithName adds a name to the builder
func (app *propertyBuilder) WithName(name string) PropertyBuilder {
	app.name = name
	return app
}

// IsPrimaryKey flags the builder as a primary key
func (app *propertyBuilder) IsPrimaryKey() PropertyBuilder {
	app.isPrimaryKey = true
	return app
}

// WithForeignKey adds a foreignKey to the builder
func (app *propertyBuilder) WithForeignKey(foreignKey Schema) PropertyBuilder {
	app.foreignKey = foreignKey
	return app
}

// WithType adds a type to the builder
func (app *propertyBuilder) WithType(typ Type) PropertyBuilder {
	app.typ = typ
	return app
}

// Now builds a new Property instance
func (app *propertyBuilder) Now() (Property, error) {
	if app.resource == nil {
		return nil, errors.New("the resource is mandatory in order to build a Property instance")
	}

	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Property instance")
	}

	var content PropertyContent
	if app.isPrimaryKey {
		content = createPropertyContentWithPrimaryKey()
	}

	if app.foreignKey != nil {
		content = createPropertyContentWithForeignKey(app.foreignKey)
	}

	if app.typ != nil {
		content = createPropertyContentWithType(app.typ)
	}

	if content == nil {
		return nil, errors.New("the property content (isPrimaryKey, ForeignKey, Type) is mandatory in order to build a Property instance")
	}

	return createProperty(app.resource, content, app.name), nil
}
