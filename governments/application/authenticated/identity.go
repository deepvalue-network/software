package authenticated

import "github.com/deepvalue-network/software/governments/domain/identities"

type identity struct {
	identityRepository identities.Repository
	identityService    identities.Service
	identityBuilder    identities.Builder
	name               string
	seed               string
	password           string
}

func createIdentity(
	identityRepository identities.Repository,
	identityService identities.Service,
	identityBuilder identities.Builder,
	name string,
	seed string,
	password string,
) Identity {
	out := identity{
		identityRepository: identityRepository,
		identityService:    identityService,
		identityBuilder:    identityBuilder,
		name:               name,
		seed:               seed,
		password:           password,
	}

	return &out
}

// Retrieve retrieves the identity
func (app *identity) Retrieve() (identities.Identity, error) {
	return app.identityRepository.Retrieve(app.name, app.seed, app.password)
}

// Update updates the identity
func (app *identity) Update(update UpdateIdentity) error {
	origin, err := app.Retrieve()
	if err != nil {
		return err
	}

	builder := app.identityBuilder.Create().WithName(app.name).WithSeed(app.seed)
	if update.HasName() {
		name := update.Name()
		builder.WithName(name)
	}

	if update.HasSeed() {
		seed := update.Seed()
		builder.WithSeed(seed)
	}

	updated, err := builder.Now()
	if err != nil {
		return err
	}

	if update.HasPassword() {
		newPass := update.Password()
		return app.identityService.UpdateWithPassword(origin, updated, app.password, newPass)
	}

	return app.identityService.Update(origin, updated, app.password)
}

// Delete deletes the identity
func (app *identity) Delete() error {
	ins, err := app.Retrieve()
	if err != nil {
		return err
	}

	return app.identityService.Delete(ins, app.password)
}
