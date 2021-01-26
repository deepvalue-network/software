package stackframe

import "errors"

type builder struct {
	isPush bool
	isPop  bool
}

func createBuilder() Builder {
	out := builder{
		isPush: false,
		isPop:  false,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// IsPush flags the builder as push
func (app *builder) IsPush() Builder {
	app.isPush = true
	return app
}

// IsPop flags the builder as pop
func (app *builder) IsPop() Builder {
	app.isPop = true
	return app
}

// Now builds a new Stackframe instance
func (app *builder) Now() (Stackframe, error) {
	if app.isPop {
		return createStackframeWithPop(), nil
	}

	if app.isPush {
		return createStackframeWithPush(), nil
	}

	return nil, errors.New("the Stackframe is invalid")
}
