package stackframe

import "errors"

type builder struct {
	isPush bool
	isPop  bool
	index  string
	skip   Skip
	save   Save
	swtch  string
}

func createBuilder() Builder {
	out := builder{
		isPush: false,
		isPop:  false,
		index:  "",
		skip:   nil,
		save:   nil,
		swtch:  "",
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

// WithSkip adds a skip to the builder
func (app *builder) WithSkip(skip Skip) Builder {
	app.skip = skip
	return app
}

// WithIndex adds an index to the builder
func (app *builder) WithIndex(indexVariable string) Builder {
	app.index = indexVariable
	return app
}

// WithSave adds a save to the builder
func (app *builder) WithSave(save Save) Builder {
	app.save = save
	return app
}

// WithSwitch adds a switch to the builder
func (app *builder) WithSwitch(swtch string) Builder {
	app.swtch = swtch
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

	if app.skip != nil {
		return createStackframeWithSkip(app.skip), nil
	}

	if app.index != "" {
		return createStackframeWithIndex(app.index), nil
	}

	if app.save != nil {
		return createStackframeWithSave(app.save), nil
	}

	if app.swtch != "" {
		return createStackframeWithSwitch(app.swtch), nil
	}

	return nil, errors.New("the Stackframe is invalid")
}
