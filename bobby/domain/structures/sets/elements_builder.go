package sets

import (
	"errors"
	"strconv"

	"github.com/deepvalue-network/software/bobby/domain/resources"
	"github.com/deepvalue-network/software/libs/hash"
)

type elementsBuilder struct {
	hashAdapter hash.Adapter
	ranked      map[uint]resources.Immutable
	unranked    []resources.Immutable
}

func createElementsBuilder(
	hashAdapter hash.Adapter,
) ElementsBuilder {
	out := elementsBuilder{
		hashAdapter: hashAdapter,
		ranked:      nil,
		unranked:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *elementsBuilder) Create() ElementsBuilder {
	return createElementsBuilder(app.hashAdapter)
}

// WithRanked adds a ranked elements to the builder
func (app *elementsBuilder) WithRanked(ranked map[uint]resources.Immutable) ElementsBuilder {
	app.ranked = ranked
	return app
}

// WithUnranked adds a unranked elements to the builder
func (app *elementsBuilder) WithUnranked(unranked []resources.Immutable) ElementsBuilder {
	app.unranked = unranked
	return app
}

// Now builds a new Elements instance
func (app *elementsBuilder) Now() (Elements, error) {
	if app.ranked != nil {
		data := [][]byte{}
		amount := uint(len(app.ranked))
		for i := uint(0); i < amount; i++ {
			if el, ok := app.ranked[i]; ok {
				data = append(data, []byte(strconv.Itoa(int(i))))
				data = append(data, el.Hash().Bytes())
			}
		}

		hsh, err := app.hashAdapter.FromMultiBytes(data)
		if err != nil {
			return nil, err
		}

		ranked := createRankedElements(*hsh, app.ranked)
		return createElementsWithRanked(ranked), nil
	}

	if app.unranked != nil {
		data := [][]byte{}
		for _, el := range app.unranked {
			data = append(data, el.Hash().Bytes())
		}

		hsh, err := app.hashAdapter.FromMultiBytes(data)
		if err != nil {
			return nil, err
		}

		unranked := createUnrankedElements(*hsh, app.unranked)
		return createElementsWithUnranked(unranked), nil
	}

	return nil, errors.New("the Elements instance is invalid")
}
