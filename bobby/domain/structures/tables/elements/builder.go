package elements

import (
	"github.com/steve-care-software/products/libs/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	list        []Element
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		list:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithElements add elements to the builder
func (app *builder) WithElements(elements []Element) Builder {
	app.list = elements
	return app
}

// Now builds a new Elements instance
func (app *builder) Now() (Elements, error) {
	if app.list == nil {
		app.list = []Element{}
	}

	data := [][]byte{
		[]byte("initial"),
	}

	for _, oneElement := range app.list {
		data = append(data, oneElement.Resource().Hash().Bytes())
	}

	hash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	mpByPropertyHash := map[string]Element{}
	for _, oneElement := range app.list {
		keyname := oneElement.Property().Resource().Hash().String()
		mpByPropertyHash[keyname] = oneElement
	}

	return createElements(*hash, app.list, mpByPropertyHash), nil
}
