package users

import "github.com/steve-care-software/products/identity/domain/accesses"

type user struct {
	name     string            `hydro:"Name, Name"`
	seed     string            `hydro:"Seed, Seed"`
	accesses accesses.Accesses `hydro:"Accesses, Accesses"`
}

func createUser(
	name string,
	seed string,
	accesses accesses.Accesses,
) User {
	out := user{
		name:     name,
		seed:     seed,
		accesses: accesses,
	}

	return &out
}

// Name returns the name
func (obj *user) Name() string {
	return obj.name
}

// Seed returns the seed
func (obj *user) Seed() string {
	return obj.seed
}

// Accesses returns the accesses
func (obj *user) Accesses() accesses.Accesses {
	return obj.accesses
}
