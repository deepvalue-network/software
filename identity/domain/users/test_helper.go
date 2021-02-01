package users

import "github.com/steve-care-software/products/identity/domain/accesses"

// CreateUserForTests creates a user instance for tests
func CreateUserForTests(encPkBitrate int) User {
	name := "my_user"
	seed := "this is my seed"
	accesses := accesses.CreateAccessesForTests(encPkBitrate)
	ins, err := NewBuilder(encPkBitrate).Create().WithName(name).WithSeed(seed).WithAccesses(accesses).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
