package votes

import (
	"errors"
	"strconv"
	"time"

	"github.com/deepvalue-network/software/governments/domain/propositions"
	"github.com/deepvalue-network/software/libs/hash"
)

type contentBuilder struct {
	hashAdapter   hash.Adapter
	prop          propositions.Proposition
	createdOn     *time.Time
	isApproved    bool
	isCancel      bool
	isDisapproved bool
}

func createContentBuilder(
	hashAdapter hash.Adapter,
) ContentBuilder {
	out := contentBuilder{
		hashAdapter:   hashAdapter,
		prop:          nil,
		createdOn:     nil,
		isApproved:    false,
		isCancel:      false,
		isDisapproved: false,
	}

	return &out
}

// Create initializes the builder
func (app *contentBuilder) Create() ContentBuilder {
	return createContentBuilder(app.hashAdapter)
}

// WithProposition adds a proposition to the builder
func (app *contentBuilder) WithProposition(proposition propositions.Proposition) ContentBuilder {
	app.prop = proposition
	return app
}

// IsApproved flags the builder as approved
func (app *contentBuilder) IsApproved() ContentBuilder {
	app.isApproved = true
	return app
}

// IsCancel flags the builder as canceled
func (app *contentBuilder) IsCancel() ContentBuilder {
	app.isCancel = true
	return app
}

// IsDisapproved flags the builder as disapproved
func (app *contentBuilder) IsDisapproved() ContentBuilder {
	app.isDisapproved = true
	return app
}

// CreatedOn adds a creation time
func (app *contentBuilder) CreatedOn(createdOn time.Time) ContentBuilder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new Content instance
func (app *contentBuilder) Now() (Content, error) {
	if app.prop == nil {
		return nil, errors.New("the proposition is mandatory in order to build a vote Content instance")
	}

	if app.createdOn == nil {
		createdOn := time.Now().UTC()
		app.createdOn = &createdOn
	}

	voteAsStr := "is_approved"
	if app.isCancel {
		voteAsStr = "is_cancel"
	}

	if app.isDisapproved {
		voteAsStr = "is_disapproved"
	}

	hash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.prop.Hash().Bytes(),
		[]byte(voteAsStr),
		[]byte(strconv.Itoa(app.createdOn.Second())),
	})

	if err != nil {
		return nil, err
	}

	if app.isApproved {
		return createContentWithApproved(*hash, app.prop, *app.createdOn), nil
	}

	if app.isCancel {
		return createContentWithCancel(*hash, app.prop, *app.createdOn), nil
	}

	if app.isDisapproved {
		return createContentWithDisapproved(*hash, app.prop, *app.createdOn), nil
	}

	return nil, errors.New("the vote Content is invalid")
}
