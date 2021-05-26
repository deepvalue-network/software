package registry

import "errors"

type fetchBuilder struct {
	from  string
	to    string
	index Index
}

func createFetchBuilder() FetchBuilder {
	out := fetchBuilder{
		from:  "",
		to:    "",
		index: nil,
	}

	return &out
}

// Create initializes the builder
func (app *fetchBuilder) Create() FetchBuilder {
	return createFetchBuilder()
}

// From adds a from variable to the builder
func (app *fetchBuilder) From(from string) FetchBuilder {
	app.from = from
	return app
}

// To adds a to variable to the builder
func (app *fetchBuilder) To(to string) FetchBuilder {
	app.to = to
	return app
}

// To adds a to variable to the builder
func (app *fetchBuilder) WithIndex(index Index) FetchBuilder {
	app.index = index
	return app
}

// Now builds a new Fetch instance
func (app *fetchBuilder) Now() (Fetch, error) {
	if app.from == "" {
		return nil, errors.New("the from variable is mandatory in order to build a Fetch instance")
	}

	if app.to == "" {
		return nil, errors.New("the to variable is mandatory in order to build a Fetch instance")
	}

	if app.index != nil {
		return createFetchWithIndex(app.from, app.to, app.index), nil
	}

	return createFetch(app.from, app.to), nil
}
