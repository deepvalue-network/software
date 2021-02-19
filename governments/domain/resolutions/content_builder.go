package resolutions

import (
	"errors"
	"strconv"
	"time"

	"github.com/deepvalue-network/software/governments/domain/propositions"
	"github.com/deepvalue-network/software/governments/domain/propositions/votes"
	"github.com/deepvalue-network/software/libs/hash"
)

type contentBuilder struct {
	hashAdapter hash.Adapter
	prop        propositions.Proposition
	votes       []votes.Vote
	createdOn   *time.Time
}

func createContentBuilder(
	hashAdapter hash.Adapter,
) ContentBuilder {
	out := contentBuilder{
		hashAdapter: hashAdapter,
		prop:        nil,
		votes:       nil,
		createdOn:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *contentBuilder) Create() ContentBuilder {
	return createContentBuilder(app.hashAdapter)
}

// WithProposition adds a proposition to the builder
func (app *contentBuilder) WithProposition(propositon propositions.Proposition) ContentBuilder {
	app.prop = propositon
	return app
}

// WithVotes add votes to the builder
func (app *contentBuilder) WithVotes(votes []votes.Vote) ContentBuilder {
	app.votes = votes
	return app
}

// CreatedOn adds a creation time to the builder
func (app *contentBuilder) CreatedOn(createdOn time.Time) ContentBuilder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new Content instance
func (app *contentBuilder) Now() (Content, error) {
	if app.prop == nil {
		return nil, errors.New("the proposition is mandatory in order to build a resolution Content instance")
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

	return createContent(*hash, app.prop, validVotes, *app.createdOn), nil
}
