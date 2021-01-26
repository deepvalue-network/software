package elements

import (
	"errors"

	"github.com/steve-care-software/products/bobby/domain/resources"
	"github.com/steve-care-software/products/bobby/domain/structures/tables/schemas"
	"github.com/steve-care-software/products/bobby/domain/structures/tables/values"
	"github.com/steve-care-software/products/libs/hash"
)

type elementBuilder struct {
	hashAdapter     hash.Adapter
	resourceBuilder resources.Builder
	property        schemas.Property
	value           values.Value
}

func createElementBuilder(
	hashAdapter hash.Adapter,
	resourceBuilder resources.Builder,
) ElementBuilder {
	out := elementBuilder{
		hashAdapter:     hashAdapter,
		resourceBuilder: resourceBuilder,
		property:        nil,
		value:           nil,
	}

	return &out
}

// Create initializes the builder
func (app *elementBuilder) Create() ElementBuilder {
	return createElementBuilder(
		app.hashAdapter,
		app.resourceBuilder,
	)
}

// WithProperty adds a property to the builder
func (app *elementBuilder) WithProperty(property schemas.Property) ElementBuilder {
	app.property = property
	return app
}

// WithValue adds a value to the builder
func (app *elementBuilder) WithValue(value values.Value) ElementBuilder {
	app.value = value
	return app
}

// Now builds a new Element instance
func (app *elementBuilder) Now() (Element, error) {
	if app.property == nil {
		return nil, errors.New("the property is mandatory in order to build an Element instance")
	}

	if app.value == nil {
		return nil, errors.New("the value is mandatory in order to build an Element instance")
	}

	hash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.property.Resource().Hash().Bytes(),
		app.value.Resource().Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	propertyResource := app.property.Resource()
	entity, err := app.resourceBuilder.Create().WithHash(*hash).WithAccessible(propertyResource).Now()
	if err != nil {
		return nil, err
	}

	return createElement(entity, app.property, app.value), nil
}
