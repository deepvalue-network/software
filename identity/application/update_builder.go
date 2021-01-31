package application

import "errors"

type updateBuilder struct {
	name     string
	password string
}

func createUpdateBuilder() UpdateBuilder {
	out := updateBuilder{
		name:     "",
		password: "",
	}

	return &out
}

// Create initializes the builder
func (app *updateBuilder) Create() UpdateBuilder {
	return createUpdateBuilder()
}

// WithName adds a name to the builder
func (app *updateBuilder) WithName(name string) UpdateBuilder {
	app.name = name
	return app
}

// WithPassword adds a password to the builder
func (app *updateBuilder) WithPassword(pass string) UpdateBuilder {
	app.password = pass
	return app
}

// Now builds a new Update instance
func (app *updateBuilder) Now() (Update, error) {
	if app.name != "" && app.password != "" {
		return createUpdateWithNameAndPassword(app.name, app.password), nil
	}

	if app.name != "" {
		return createUpdateWithName(app.name), nil
	}

	if app.password != "" {
		return createUpdateWithPassword(app.password), nil
	}

	return nil, errors.New("the update instance is invalid")
}
