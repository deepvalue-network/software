package parsers

import "errors"

type saveBuilder struct {
	to   string
	from string
}

func createSaveBuilder() SaveBuilder {
	out := saveBuilder{
		to:   "",
		from: "",
	}

	return &out
}

// Create initializes the builder
func (app *saveBuilder) Create() SaveBuilder {
	return createSaveBuilder()
}

// From adds a from to the builder
func (app *saveBuilder) From(from string) SaveBuilder {
	app.from = from
	return app
}

// To adds a to to the builder
func (app *saveBuilder) To(to string) SaveBuilder {
	app.to = to
	return app
}

// Now builds a new Save instance
func (app *saveBuilder) Now() (Save, error) {
	if app.to == "" {
		return nil, errors.New("the to is mandatory in order to build a Save instance")
	}

	if app.from != "" {
		return createSaveWithFrom(app.to, app.from), nil
	}

	return createSave(app.to), nil
}
