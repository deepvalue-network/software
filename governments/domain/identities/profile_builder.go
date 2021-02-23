package identities

import "errors"

type profileBuilder struct {
	name string
	rank uint
}

func createProfileBuilder() ProfileBuilder {
	out := profileBuilder{
		name: "",
		rank: 0,
	}

	return &out
}

// Create initializes the builder
func (app *profileBuilder) Create() ProfileBuilder {
	return createProfileBuilder()
}

// WithName adds a name to the builder
func (app *profileBuilder) WithName(name string) ProfileBuilder {
	app.name = name
	return app
}

// WithRank adds a rank to the builder
func (app *profileBuilder) WithRank(rank uint) ProfileBuilder {
	app.rank = rank
	return app
}

// Now builds a new Profile instance
func (app *profileBuilder) Now() (Profile, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Profile instance")
	}

	if app.rank > 0 {
		return createProfileWithRank(app.name, app.rank), nil
	}

	return createProfile(app.name), nil
}
