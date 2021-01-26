package bodies

import (
	"errors"
	"strconv"
	"time"

	"github.com/steve-care-software/products/bobby/domain/resources"
	"github.com/steve-care-software/products/bobby/domain/transactions/bodies/access"
	"github.com/steve-care-software/products/bobby/domain/transactions/bodies/containers"
	"github.com/steve-care-software/products/bobby/domain/transactions/bodies/contents"
	"github.com/steve-care-software/products/libs/hash"
)

type builder struct {
	hashAdapter     hash.Adapter
	resourceBuilder resources.ImmutableBuilder
	container       containers.Transaction
	content         contents.Transaction
	access          access.Transaction
	executesOn      *time.Time
}

func createBuilder(
	hashAdapter hash.Adapter,
	resourceBuilder resources.ImmutableBuilder,
) Builder {
	out := builder{
		hashAdapter:     hashAdapter,
		resourceBuilder: resourceBuilder,
		container:       nil,
		content:         nil,
		access:          nil,
		executesOn:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter, app.resourceBuilder)
}

// WithContainer adds a container to the builder
func (app *builder) WithContainer(container containers.Transaction) Builder {
	app.container = container
	return app
}

// WithContent adds a content to the builder
func (app *builder) WithContent(content contents.Transaction) Builder {
	app.content = content
	return app
}

// WithAccess adds an access to the builder
func (app *builder) WithAccess(access access.Transaction) Builder {
	app.access = access
	return app
}

// ExecutesOn adds an execution time to the builder
func (app *builder) ExecutesOn(executesOn time.Time) Builder {
	app.executesOn = &executesOn
	return app
}

// Now builds a new Body instance
func (app *builder) Now() (Body, error) {
	data := [][]byte{}
	var content Content
	if app.container != nil {
		data = append(data, app.container.Hash().Bytes())
		content = createContentWithContainer(app.container)
	}

	if app.content != nil {
		data = append(data, app.content.Hash().Bytes())
		content = createContentWithContent(app.content)
	}

	if app.access != nil {
		data = append(data, app.access.Hash().Bytes())
		content = createContentWithAccess(app.access)
	}

	if content == nil {
		return nil, errors.New("the content (container, content, access) is mandatory in order to build a transaction Body instance")
	}

	if app.executesOn != nil {
		data = append(data, []byte(strconv.Itoa(app.executesOn.Nanosecond())))
	}

	hsh, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	immutable, err := app.resourceBuilder.Create().WithHash(*hsh).Now()
	if err != nil {
		return nil, err
	}

	if app.executesOn != nil {
		return createBodyWithExecutesOn(immutable, content, app.executesOn), nil
	}

	return createBody(immutable, content), nil
}
