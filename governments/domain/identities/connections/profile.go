package connections

type profile struct {
	name string
	rank uint
}

func createProfile(
	name string,
) Profile {
	return createProfileInternally(name, 0)
}

func createProfileWithRank(
	name string,
	rank uint,
) Profile {
	return createProfileInternally(name, rank)
}

func createProfileInternally(
	name string,
	rank uint,
) Profile {
	out := profile{
		name: name,
		rank: rank,
	}

	return &out
}

// Name returns the name
func (obj *profile) Name() string {
	return obj.name
}

// HasRank returns true if there is a rank, false otherwise
func (obj *profile) HasRank() bool {
	return obj.rank > 0
}

// Rank returns the rank, if any
func (obj *profile) Rank() uint {
	return obj.rank
}
