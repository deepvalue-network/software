package resolutions

import (
	"errors"
	"strconv"
	"time"

	"github.com/deepvalue-network/software/governments/domain/propositions"
	"github.com/deepvalue-network/software/governments/domain/propositions/votes"
	"github.com/deepvalue-network/software/libs/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	prop        propositions.Proposition
	votes       []votes.Vote
	createdOn   *time.Time
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		prop:        nil,
		votes:       nil,
		createdOn:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithProposition adds a proposition to the builder
func (app *builder) WithProposition(propositon propositions.Proposition) Builder {
	app.prop = propositon
	return app
}

// WithVotes add votes to the builder
func (app *builder) WithVotes(votes []votes.Vote) Builder {
	app.votes = votes
	return app
}

// CreatedOn adds a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new Resolution instance
func (app *builder) Now() (Resolution, error) {
	if app.prop == nil {
		return nil, errors.New("the proposition is mandatory in order to build a Resolution instance")
	}

	if app.votes == nil {
		app.votes = []votes.Vote{}
	}

	if app.createdOn == nil {
		createdOn := time.Now().UTC()
		app.createdOn = &createdOn
	}

	validVotes := []votes.Vote{}
	holders := app.prop.Content().Government().ShareHolders()
	for _, oneVote := range app.votes {
		sig := oneVote.Signature()
		if !holders.Validate(sig) {
			continue
		}

		validVotes = append(validVotes, oneVote)
	}

	data := [][]byte{
		app.prop.Hash().Bytes(),
		[]byte(strconv.Itoa(app.createdOn.Second())),
	}

	for _, oneValidVote := range validVotes {
		data = append(data, oneValidVote.Hash().Bytes())
	}

	hash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createResolution(*hash, app.prop, validVotes, *app.createdOn), nil
}
