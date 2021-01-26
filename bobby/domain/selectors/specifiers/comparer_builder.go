package specifiers

import (
	"errors"

	"github.com/steve-care-software/products/libs/hash"
)

type comparerBuilder struct {
	hashAdapter hash.Adapter
	first       Identifier
	second      Identifier
	isAnd       bool
}

func createComparerBuilder(
	hashAdapter hash.Adapter,
) ComparerBuilder {
	out := comparerBuilder{
		hashAdapter: hashAdapter,
		first:       nil,
		second:      nil,
		isAnd:       false,
	}

	return &out
}

// Create initializes the builder
func (app *comparerBuilder) Create() ComparerBuilder {
	return createComparerBuilder(app.hashAdapter)
}

// WithFirst adds a first identifier to the builder
func (app *comparerBuilder) WithFirst(first Identifier) ComparerBuilder {
	app.first = first
	return app
}

// WithSecond adds a second identifier to the builder
func (app *comparerBuilder) WithSecond(second Identifier) ComparerBuilder {
	app.second = second
	return app
}

// IsAnd flags the builder as an and comparer
func (app *comparerBuilder) IsAnd() ComparerBuilder {
	app.isAnd = true
	return app
}

// Now builds a new Comparer instance
func (app *comparerBuilder) Now() (Comparer, error) {
	if app.first == nil {
		return nil, errors.New("the first Identifier is mandatory in order to build a Comparer instance")
	}

	if app.second == nil {
		return nil, errors.New("the second Identifier is mandatory in order to build a Comparer instance")
	}

	data := [][]byte{
		app.first.Hash().Bytes(),
		app.second.Hash().Bytes(),
	}

	if app.isAnd {
		data = append(data, []byte("and"))
		hsh, err := app.hashAdapter.FromMultiBytes(data)
		if err != nil {
			return nil, err
		}

		return createComparerWithAnd(*hsh, app.first, app.second), nil
	}

	data = append(data, []byte("or"))
	hsh, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createComparerWithOr(*hsh, app.first, app.second), nil
}
