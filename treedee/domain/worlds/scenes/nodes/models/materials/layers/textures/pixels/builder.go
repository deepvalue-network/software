package pixels

import (
	"errors"

	"github.com/deepvalue-network/software/treedee/domain/worlds/colors"
)

type builder struct {
	color colors.Color
	alpha uint8
}

func createBuilder() Builder {
	out := builder{
		color: nil,
		alpha: 0x00,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithColor adds a color to the builder
func (app *builder) WithColor(color colors.Color) Builder {
	app.color = color
	return app
}

// WithAlpha adds an alpha to the builder
func (app *builder) WithAlpha(alpha uint8) Builder {
	app.alpha = alpha
	return app
}

// Now builds a new Pixel instance
func (app *builder) Now() (Pixel, error) {
	if app.color == nil {
		return nil, errors.New("the color is mandatory in order to build a Pixel instance")
	}

	return createPixel(app.color, app.alpha), nil
}
