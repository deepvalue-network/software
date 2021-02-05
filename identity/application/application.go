package application

import "github.com/deepvalue-network/software/identity/domain/users"

type application struct {
	repository users.Repository
	service    users.Service
	builder    users.Builder
}

func createApplication(
	repository users.Repository,
	service users.Service,
	builder users.Builder,
) Application {
	out := application{
		repository: repository,
		service:    service,
		builder:    builder,
	}

	return &out
}

// List lists the available user names
func (app *application) List() ([]string, error) {
	return app.repository.List()
}

// Retrieve retrieves a user
func (app *application) Retrieve(name string, seed string, password string) (users.User, error) {
	return app.repository.Retrieve(name, seed, password)
}

// Insert inserts a user
func (app *application) Insert(name string, seed string, password string) error {
	user, err := app.builder.Create().WithName(name).WithSeed(seed).Now()
	if err != nil {
		return err
	}

	return app.service.Insert(user, password)
}

// Update updates a user
func (app *application) Update(name string, seed string, password string, update Update) error {
	original, err := app.builder.Create().WithName(name).WithSeed(seed).Now()
	if err != nil {
		return err
	}

	userBuilder := app.builder.Create().WithName(name).WithSeed(seed)
	if update.HasName() {
		updatedName := update.Name()
		userBuilder.WithName(updatedName)
	}

	updatedUser, err := userBuilder.Now()
	if err != nil {
		return err
	}

	if update.HasPassword() {
		updatedPassword := update.Password()
		return app.service.UpdateWithPassword(original, updatedUser, password, updatedPassword)
	}

	return app.service.Update(original, updatedUser, password)
}

// Delete deletes a user
func (app *application) Delete(name string, seed string, password string) error {
	user, err := app.builder.Create().WithName(name).WithSeed(seed).Now()
	if err != nil {
		return err
	}

	return app.service.Delete(user, password)
}
