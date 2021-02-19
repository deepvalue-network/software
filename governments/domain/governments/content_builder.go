package governments

import (
	"errors"
	"strconv"
	"time"

	"github.com/deepvalue-network/software/blockchain/domain/chains"
	"github.com/deepvalue-network/software/libs/hash"
)

type contentBuilder struct {
	hashAdapter              hash.Adapter
	chain                    chains.Chain
	minPowerToPassResolution uint
	minPowerToPropose        uint
	canCancelVote            bool
	sharesVelocity           uint
	sharesCap                uint
	createdOn                *time.Time
}

func createContentBuilder(
	hashAdapter hash.Adapter,
) ContentBuilder {
	out := contentBuilder{
		hashAdapter:              hashAdapter,
		chain:                    nil,
		minPowerToPassResolution: 0,
		minPowerToPropose:        0,
		canCancelVote:            false,
		sharesVelocity:           0,
		sharesCap:                0,
		createdOn:                nil,
	}

	return &out
}

// Create initializes the builder
func (app *contentBuilder) Create() ContentBuilder {
	return createContentBuilder(app.hashAdapter)
}

// WithChain adds a chain to the builder
func (app *contentBuilder) WithChain(chain chains.Chain) ContentBuilder {
	app.chain = chain
	return app
}

// WithMinPowerToPassResolution adds a minimum power to pass resolution
func (app *contentBuilder) WithMinPowerToPassResolution(minPowerToPassRes uint) ContentBuilder {
	app.minPowerToPassResolution = minPowerToPassRes
	return app
}

// WithMinPowerToPropose adds a minimum power to propose resolution
func (app *contentBuilder) WithMinPowerToPropose(minPowerToPropose uint) ContentBuilder {
	app.minPowerToPropose = minPowerToPropose
	return app
}

// WithSharesVelocity adds a shares velocity
func (app *contentBuilder) WithSharesVelocity(sharesVelocity uint) ContentBuilder {
	app.sharesVelocity = sharesVelocity
	return app
}

// WithSharesCap adds a shares cap
func (app *contentBuilder) WithSharesCap(sharesCap uint) ContentBuilder {
	app.sharesCap = sharesCap
	return app
}

// CanCancelVote flags the builder as canCancelVote
func (app *contentBuilder) CanCancelVote() ContentBuilder {
	app.canCancelVote = true
	return app
}

// CreatedOn adds the creation time
func (app *contentBuilder) CreatedOn(createdOn time.Time) ContentBuilder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new Content instance
func (app *contentBuilder) Now() (Content, error) {
	if app.chain == nil {
		return nil, errors.New("the chain is mandatory in order to build a Content instance")
	}

	if app.minPowerToPassResolution <= 0 {
		return nil, errors.New("the minPowerToPassResolution must be greater than zero (0) in order to build a Content instance")
	}

	if app.minPowerToPropose <= 0 {
		return nil, errors.New("the minPowerToPropose must be greater than zero (0) in order to build a Content instance")
	}

	if app.sharesVelocity <= 0 {
		return nil, errors.New("the sharesVelocity must be greater than zero (0) in order to build a Content instance")
	}

	if app.sharesCap <= 0 {
		return nil, errors.New("the sharesCap must be greater than zero (0) in order to build a Content instance")
	}

	if app.createdOn == nil {
		createdOn := time.Now().UTC()
		app.createdOn = &createdOn
	}

	canCancelStr := "false"
	if app.canCancelVote {
		canCancelStr = "true"
	}

	hash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.chain.ID().Bytes(),
		[]byte(strconv.Itoa(int(app.minPowerToPassResolution))),
		[]byte(strconv.Itoa(int(app.minPowerToPropose))),
		[]byte(strconv.Itoa(int(app.sharesVelocity))),
		[]byte(strconv.Itoa(int(app.sharesCap))),
		[]byte(canCancelStr),
		[]byte(strconv.Itoa(app.createdOn.Second())),
	})

	if err != nil {
		return nil, err
	}

	return createContent(
		*hash,
		app.chain,
		app.minPowerToPassResolution,
		app.minPowerToPropose,
		app.canCancelVote,
		app.sharesVelocity,
		app.sharesCap,
		*app.createdOn,
	), nil
}
