package parsers

import "errors"

type fetchRegistryBuilder struct {
	to    string
	from  string
	index IntPointer
}

func createFetchRegistryBuilder() FetchRegistryBuilder {
	out := fetchRegistryBuilder{
		to:    "",
		from:  "",
		index: nil,
	}

	return &out
}

// Create initializes the builder
func (app *fetchRegistryBuilder) Create() FetchRegistryBuilder {
	return createFetchRegistryBuilder()
}

// From adds a from variable to the builder
func (app *fetchRegistryBuilder) From(from string) FetchRegistryBuilder {
	app.from = from
	return app
}

// To adds a to variable to the builder
func (app *fetchRegistryBuilder) To(to string) FetchRegistryBuilder {
	app.to = to
	return app
}

// WithIndex adds an index to the builder
func (app *fetchRegistryBuilder) WithIndex(index IntPointer) FetchRegistryBuilder {
	app.index = index
	return app
}

// Now builds a new FetchRegistry instance
func (app *fetchRegistryBuilder) Now() (FetchRegistry, error) {
	if app.from == "" {
		return nil, errors.New("the from variable is mandatory in order to build a FetchRegistry instance")
	}

	if app.to == "" {
		return nil, errors.New("the to variable is mandatory in order to build a FetchRegistry instance")
	}

	if app.index != nil {
		return createFetchRegisterWithIndex(app.to, app.from, app.index), nil
	}

	return createFetchRegister(app.to, app.from), nil
}
