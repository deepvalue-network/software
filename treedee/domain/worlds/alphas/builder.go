package alphas

import "errors"

type builder struct {
	alpha    uint8
	variable string
}

func createBuilder() Builder {
	out := builder{
		alpha:    0x00,
		variable: "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithAlpha adds an alpha to the builder
func (app *builder) WithAlpha(alpha uint8) Builder {
	app.alpha = alpha
	return app
}

// WithVariable adds a variable to the builder
func (app *builder) WithVariable(variable string) Builder {
	app.variable = variable
	return app
}

// Now builds a new Alpha instance
func (app *builder) Now() (Alpha, error) {
	if app.variable == "" {
		return nil, errors.New("the variable is mandatory in order to build an Alpha instance")
	}

	return createAlpha(app.alpha, app.variable), nil
}
