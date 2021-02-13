package colors

type builder struct {
	red   uint8
	green uint8
	blue  uint8
}

func createBuilder() Builder {
	out := builder{
		red:   0x00,
		green: 0x00,
		blue:  0x00,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithRed adds a red color to the builder
func (app *builder) WithRed(red uint8) Builder {
	app.red = red
	return app
}

// WithGreen adds a green color to the builder
func (app *builder) WithGreen(green uint8) Builder {
	app.green = green
	return app
}

// WithBlue adds a blue color to the builder
func (app *builder) WithBlue(blue uint8) Builder {
	app.blue = blue
	return app
}

// Now builds a new Color instance
func (app *builder) Now() Color {
	return createColor(app.red, app.green, app.blue)
}
