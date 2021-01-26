package values

import (
	"errors"

	uuid "github.com/satori/go.uuid"
	"github.com/steve-care-software/products/bobby/domain/resources"
)

type builder struct {
	resource   resources.Resource
	id         *uuid.UUID
	stringVal  *string
	intVal     *int
	float32Val *float32
	float64Val *float64
	data       []byte
}

func createBuilder() Builder {
	out := builder{
		resource:   nil,
		id:         nil,
		stringVal:  nil,
		intVal:     nil,
		float32Val: nil,
		float64Val: nil,
		data:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithResource adds a resource to the builder
func (app *builder) WithResource(res resources.Resource) Builder {
	app.resource = res
	return app
}

// WithID adds an id to the builder
func (app *builder) WithID(id *uuid.UUID) Builder {
	app.id = id
	return app
}

// WithString adds a string to the builder
func (app *builder) WithString(stringVal string) Builder {
	app.stringVal = &stringVal
	return app
}

// WithInt adds an int to the builder
func (app *builder) WithInt(intVal int) Builder {
	app.intVal = &intVal
	return app
}

// WithFloat32 adds a float32 to the builder
func (app *builder) WithFloat32(float32Val float32) Builder {
	app.float32Val = &float32Val
	return app
}

// WithFloat64 adds a float64 to the builder
func (app *builder) WithFloat64(float64Val float64) Builder {
	app.float64Val = &float64Val
	return app
}

// WithData adds data to the builder
func (app *builder) WithData(data []byte) Builder {
	app.data = data
	return app
}

// Now builds a new Value instance
func (app *builder) Now() (Value, error) {
	if app.resource == nil {
		return nil, errors.New("the resource is mandatory in order to build a Value instance")
	}

	var content ValueContent
	if app.id != nil {
		content = createValueContentWithID(app.id)
	}

	if app.stringVal != nil {
		content = createValueContentWithString(app.stringVal)
	}

	if app.intVal != nil {
		content = createValueContentWithInt(app.intVal)
	}

	if app.float32Val != nil {
		content = createValueContentWithFloat32(app.float32Val)
	}

	if app.float64Val != nil {
		content = createValueContentWithFloat64(app.float64Val)
	}

	if app.data != nil {
		content = createValueContentWithData(app.data)
	}

	if content == nil {
		return nil, errors.New("the content is mandatory in order to build a Value instance")
	}

	return createValue(app.resource, content), nil
}
