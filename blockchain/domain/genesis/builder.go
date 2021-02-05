package genesis

import (
	"errors"
	"strconv"

	"github.com/deepvalue-network/software/libs/hash"
)

type builder struct {
	hashAdapter     hash.Adapter
	miningValue     uint8
	blockBaseDiff   uint
	incrPerHashDiff float64
	linkDiff        uint
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter:     hashAdapter,
		miningValue:     DefaultMiningValue,
		blockBaseDiff:   0,
		incrPerHashDiff: 0.0,
		linkDiff:        0,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithMiningValue adds a mining value to the builder
func (app *builder) WithMiningValue(miningValue uint8) Builder {
	app.miningValue = miningValue
	return app
}

// WithBlockBaseDifficulty adds a block base difficulty to the builder
func (app *builder) WithBlockBaseDifficulty(blockBaseDiff uint) Builder {
	app.blockBaseDiff = blockBaseDiff
	return app
}

// WithBlockIncreasePerHashDifficulty adds a block increase per hash difficulty to the builder
func (app *builder) WithBlockIncreasePerHashDifficulty(incrPerHashDiff float64) Builder {
	app.incrPerHashDiff = incrPerHashDiff
	return app
}

// WithLinkDifficulty adds a link difficulty to the builder
func (app *builder) WithLinkDifficulty(linkDiff uint) Builder {
	app.linkDiff = linkDiff
	return app
}

// Now builds a new Genesis instance
func (app *builder) Now() (Genesis, error) {
	if app.blockBaseDiff == 0 {
		return nil, errors.New("the block base difficulty must be greater than zero (0)")
	}

	if app.incrPerHashDiff == 0.0 {
		return nil, errors.New("the block increase per hash difficulty must be greater than zero (0.0)")
	}

	if app.linkDiff == 0 {
		return nil, errors.New("the link difficulty must be greater than zero (0)")
	}

	if app.miningValue > 9 {
		return nil, errors.New("the mining value must be a number between 0 and 9")
	}

	hash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(strconv.Itoa(int(app.blockBaseDiff))),
		[]byte(strconv.FormatFloat(float64(app.incrPerHashDiff), 'f', -1, 64)),
		[]byte(strconv.Itoa(int(app.linkDiff))),
		[]byte(strconv.Itoa(int(app.miningValue))),
	})

	if err != nil {
		return nil, err
	}

	block := createBlock(app.blockBaseDiff, app.incrPerHashDiff)
	diff := createDifficulty(block, app.linkDiff)
	return createGenesis(*hash, app.miningValue, diff), nil

}
