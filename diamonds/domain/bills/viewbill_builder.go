package bills

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/steve-care-software/products/libs/hash"
)

type viewBillBuilder struct {
	hashAdapter hash.Adapter
	bill        Bill
	seed        string
	amount      uint64
}

func createViewBillBuilder(
	hashAdapter hash.Adapter,
) ViewBillBuilder {
	out := viewBillBuilder{
		hashAdapter: hashAdapter,
		bill:        nil,
		seed:        "",
		amount:      0,
	}

	return &out
}

// Create initializes the builder
func (app *viewBillBuilder) Create() ViewBillBuilder {
	return createViewBillBuilder(app.hashAdapter)
}

// WithBill adds a bill to the builder
func (app *viewBillBuilder) WithBill(bill Bill) ViewBillBuilder {
	app.bill = bill
	return app
}

// WithSeed adds a seed to the builder
func (app *viewBillBuilder) WithSeed(seed string) ViewBillBuilder {
	app.seed = seed
	return app
}

// WithAmount adds an amount
func (app *viewBillBuilder) WithAmount(amount uint64) ViewBillBuilder {
	app.amount = amount
	return app
}

// Now builds a new ViewBill instance
func (app *viewBillBuilder) Now() (ViewBill, error) {
	if app.bill == nil {
		return nil, errors.New("the bill is mandatory in order to build a view Bill instance")
	}

	if app.seed == "" {
		return nil, errors.New("the seed is mandatory in order to build a view Bill instance")
	}

	// make sure the seed works with the given spent genesis:
	seedHash, err := app.hashAdapter.FromBytes([]byte(app.seed))
	if err != nil {
		return nil, err
	}

	content := app.bill.Content()
	if !content.Seed().Compare(*seedHash) {
		str := fmt.Sprintf(
			"the given seed (seed: %s), when hashed (hash: %s) does not match the hashed seed (hash: %s inside the given spent Bill (hash: %s)",
			app.seed,
			seedHash.String(),
			content.Seed().String(),
			app.bill.Hash().String(),
		)
		return nil, errors.New(str)
	}

	// make sure the hashed amount fits:
	hashedAmount, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.seed),
		[]byte(strconv.Itoa(int(app.amount))),
	})

	if err != nil {
		return nil, err
	}

	if !content.Amount().Compare(*hashedAmount) {
		str := fmt.Sprintf(
			"the given amount (amount: %d), when hashed by combing it with the seed (seed: %s, hash: %s) does not match the hashed amount (hash: %s inside the given spent Bill (hash: %s)",
			app.amount,
			app.seed,
			hashedAmount.String(),
			content.Amount().String(),
			app.bill.Hash().String(),
		)
		return nil, errors.New(str)
	}

	hsh, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.bill.Hash().Bytes(),
		[]byte(app.seed),
	})

	if err != nil {
		return nil, err
	}

	return createViewBill(*hsh, app.bill, app.seed, app.amount), nil
}
