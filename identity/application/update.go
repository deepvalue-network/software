package application

type update struct {
	name     string
	password string
}

func createUpdateWithName(
	name string,
) Update {
	return createUpdateInternally(name, "")
}

func createUpdateWithPassword(
	password string,
) Update {
	return createUpdateInternally("", password)
}

func createUpdateWithNameAndPassword(
	name string,
	password string,
) Update {
	return createUpdateInternally(name, password)
}

func createUpdateInternally(
	name string,
	password string,
) Update {
	out := update{
		name:     name,
		password: password,
	}

	return &out
}

// HasName returns true if there is a name, false otherwise
func (obj *update) HasName() bool {
	return obj.name != ""
}

// Name returns the name, if any
func (obj *update) Name() string {
	return obj.name
}

// HasPassword returns true if there is a password, false otherwise
func (obj *update) HasPassword() bool {
	return obj.password != ""
}

// Password returns the password, if any
func (obj *update) Password() string {
	return obj.password
}
