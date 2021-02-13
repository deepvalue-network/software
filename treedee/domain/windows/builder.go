package windows

import "errors"

type builder struct {
	title        string
	width        uint
	height       uint
	isResizable  bool
	isFullscreen bool
}

func createBuilder() Builder {
	out := builder{
		title:        "",
		width:        0,
		height:       0,
		isResizable:  false,
		isFullscreen: false,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithTitle adds a title to the builder
func (app *builder) WithTitle(title string) Builder {
	app.title = title
	return app
}

// WithWidth adds a width to the builder
func (app *builder) WithWidth(width uint) Builder {
	app.width = width
	return app
}

// WithHeight adds an height to the builder
func (app *builder) WithHeight(height uint) Builder {
	app.height = height
	return app
}

// IsResizable flags the builder as resizable
func (app *builder) IsResizable() Builder {
	app.isResizable = true
	return app
}

// IsFullscreen flags the builder as fullscreen
func (app *builder) IsFullscreen() Builder {
	app.isFullscreen = true
	return app
}

// Now builds a new Window instance
func (app *builder) Now() (Window, error) {
	if app.width <= 0 {
		return nil, errors.New("the width must be greater than zero (0)")
	}

	if app.height <= 0 {
		return nil, errors.New("the height must be greater than zero (0)")
	}

	return createWindow(
		app.title,
		app.width,
		app.height,
		app.isResizable,
		app.isFullscreen,
	), nil
}
