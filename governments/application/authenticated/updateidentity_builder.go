package authenticated

import "errors"

type updateIdentityBuilder struct {
	name     string
	seed     string
	password string
}

func createUpdateIdentityBuilder() UpdateIdentityBuilder {
	out := updateIdentityBuilder{
		name:     "",
		seed:     "",
		password: "",
	}

	return &out
}

// Create initializes the builder
func (app *updateIdentityBuilder) Create() UpdateIdentityBuilder {
	return createUpdateIdentityBuilder()
}

// WithName adds a name to the builder
func (app *updateIdentityBuilder) WithName(name string) UpdateIdentityBuilder {
	app.name = name
	return app
}

// WithSeed adds a seed to the builder
func (app *updateIdentityBuilder) WithSeed(seed string) UpdateIdentityBuilder {
	app.seed = seed
	return app
}

// WithPassword adds a password to the builder
func (app *updateIdentityBuilder) WithPassword(password string) UpdateIdentityBuilder {
	app.password = password
	return app
}

// Now builds a new UpdateIdentity instance
func (app *updateIdentityBuilder) Now() (UpdateIdentity, error) {
	if app.name != "" && app.seed != "" && app.password != "" {
		return createUpdateIdentityWithNameAndSeedAndPassword(app.name, app.seed, app.password), nil
	}

	if app.name != "" && app.seed != "" {
		return createUpdateIdentityWithNameAndSeed(app.name, app.seed), nil
	}

	if app.name != "" && app.password != "" {
		return createUpdateIdentityWithNameAndPassword(app.name, app.password), nil
	}

	if app.name != "" {
		return createUpdateIdentityWithName(app.name), nil
	}

	if app.seed != "" {
		return createUpdateIdentityWithSeed(app.seed), nil
	}

	if app.password != "" {
		return createUpdateIdentityWithPassword(app.password), nil
	}

	return nil, errors.New("the UpdateIdentity instance is invalid")
}
