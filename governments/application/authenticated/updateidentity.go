package authenticated

type updateIdentity struct {
	name     string
	seed     string
	password string
}

func createUpdateIdentityWithName(name string) UpdateIdentity {
	return createUpdateIdentityInternally(name, "", "")
}

func createUpdateIdentityWithSeed(seed string) UpdateIdentity {
	return createUpdateIdentityInternally("", seed, "")
}

func createUpdateIdentityWithPassword(password string) UpdateIdentity {
	return createUpdateIdentityInternally("", "", password)
}

func createUpdateIdentityWithNameAndSeed(name string, seed string) UpdateIdentity {
	return createUpdateIdentityInternally(name, seed, "")
}

func createUpdateIdentityWithNameAndPassword(name string, password string) UpdateIdentity {
	return createUpdateIdentityInternally(name, "", password)
}

func createUpdateIdentityWithNameAndSeedAndPassword(name string, seed string, password string) UpdateIdentity {
	return createUpdateIdentityInternally(name, seed, password)
}

func createUpdateIdentityInternally(
	name string,
	seed string,
	password string,
) UpdateIdentity {
	out := updateIdentity{
		name:     name,
		seed:     seed,
		password: password,
	}

	return &out
}

// HasName returns true if there is a name, false otherwise
func (obj *updateIdentity) HasName() bool {
	return obj.name != ""
}

// Name returns the name, if any
func (obj *updateIdentity) Name() string {
	return obj.name
}

// HasSeed returns true if there is a seed, false otherwise
func (obj *updateIdentity) HasSeed() bool {
	return obj.seed != ""
}

// Seed returns the seed, if any
func (obj *updateIdentity) Seed() string {
	return obj.seed
}

// HasPassword returns true if there is a password, false otherwise
func (obj *updateIdentity) HasPassword() bool {
	return obj.password != ""
}

// Password returns the password, if any
func (obj *updateIdentity) Password() string {
	return obj.password
}
