package registry

type registry struct {
	fetch      Fetch
	register   Register
	unregister string
}

func createRegistryWithFetch(
	fetch Fetch,
) Registry {
	return createRegistryInternally(fetch, nil, "")
}

func createRegistryWithRegister(
	register Register,
) Registry {
	return createRegistryInternally(nil, register, "")
}

func createRegistryWithUnregister(
	unregister string,
) Registry {
	return createRegistryInternally(nil, nil, unregister)
}

func createRegistryInternally(
	fetch Fetch,
	register Register,
	unregister string,
) Registry {
	out := registry{
		fetch:      fetch,
		register:   register,
		unregister: unregister,
	}

	return &out
}

// IsFetch returns true if there is a fetch, false otherwise
func (obj *registry) IsFetch() bool {
	return obj.fetch != nil
}

// Fetch returns the fetch, if any
func (obj *registry) Fetch() Fetch {
	return obj.fetch
}

// IsRegister returns true if there is a register, false otherwise
func (obj *registry) IsRegister() bool {
	return obj.register != nil
}

// Register returns the register, if any
func (obj *registry) Register() Register {
	return obj.register
}

// IsUnregister returns true if there is an unregister, false otherwise
func (obj *registry) IsUnregister() bool {
	return obj.unregister != ""
}

// Unregister returns the unregister, if any
func (obj *registry) Unregister() string {
	return obj.unregister
}
