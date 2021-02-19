package views

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/deepvalue-network/software/governments/domain/governments/shareholders/transfers"
	"github.com/deepvalue-network/software/libs/hash"
)

type sectionBuilder struct {
	hashAdapter hash.Adapter
	transfer    transfers.Transfer
	seed        string
	amount      uint
}

func createSectionBuilder(
	hashAdapter hash.Adapter,
) SectionBuilder {
	out := sectionBuilder{
		hashAdapter: hashAdapter,
		transfer:    nil,
		seed:        "",
		amount:      0,
	}

	return &out
}

// Create initializes the builder
func (app *sectionBuilder) Create() SectionBuilder {
	return createSectionBuilder(app.hashAdapter)
}

// WithTransfer adds a transfer to the builder
func (app *sectionBuilder) WithTransfer(transfer transfers.Transfer) SectionBuilder {
	app.transfer = transfer
	return app
}

// WithSeed adds a seed to the builder
func (app *sectionBuilder) WithSeed(seed string) SectionBuilder {
	app.seed = seed
	return app
}

// WithAmount adds an amount to the builder
func (app *sectionBuilder) WithAmount(amount uint) SectionBuilder {
	app.amount = amount
	return app
}

// Now builds a new Section instance
func (app *sectionBuilder) Now() (Section, error) {
	if app.transfer == nil {
		return nil, errors.New("the transfer is mandatory in order to build a Section instance")
	}

	if app.seed == "" {
		return nil, errors.New("the seed is mandatory in order to build a Section instance")
	}

	if app.amount <= 0 {
		return nil, errors.New("the amount must be greater than zero (0) in order to build a Section instance")
	}

	amountHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.seed),
		[]byte(strconv.Itoa(int(app.amount))),
	})

	if err != nil {
		return nil, err
	}

	transferContent := app.transfer.Content()
	if !transferContent.Amount().Compare(*amountHash) {
		str := fmt.Sprintf("the amount (%d) and seed (%s) when hashed (%s) does not equal the transfer amount hash (%s)", app.amount, app.seed, amountHash.String(), transferContent.Amount().String())
		return nil, errors.New(str)
	}

	seedHash, err := app.hashAdapter.FromBytes([]byte(app.seed))
	if err != nil {
		return nil, err
	}

	if !transferContent.Seed().Compare(*seedHash) {
		str := fmt.Sprintf("the seed (%s) when hashed (%s) does not equal the transfer seed hash (%s)", app.seed, seedHash.String(), transferContent.Seed().String())
		return nil, errors.New(str)
	}

	hash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.transfer.Hash().Bytes(),
		[]byte(app.seed),
		[]byte(strconv.Itoa(int(app.amount))),
	})

	if err != nil {
		return nil, err
	}

	return createSection(*hash, app.transfer, app.seed, app.amount), nil
}
