package identities

type identity struct {
	name    string
	seed    string
	holders ShareHolders
}

func createIdentity(
	name string,
	seed string,
	holders ShareHolders,
) Identity {
	out := identity{
		name:    name,
		seed:    seed,
		holders: holders,
	}

	return &out
}

// Name returns the name
func (obj *identity) Name() string {
	return obj.name
}

// Seed returns the seed
func (obj *identity) Seed() string {
	return obj.seed
}

// ShareHolders returns the shareHolders
func (obj *identity) ShareHolders() ShareHolders {
	return obj.holders
}
