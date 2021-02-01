package accesses

import "github.com/steve-care-software/products/identity/domain/accesses/access"

// CreateAccessesForTests creates a new accesses instance for tests
func CreateAccessesForTests(encPkBitrate int) Accesses {
	first := access.CreateAccessForTests(encPkBitrate)
	second := access.CreateAccessForTests(encPkBitrate)
	third := access.CreateAccessForTests(encPkBitrate)

	mp := map[string]access.Access{
		"first":  first,
		"second": second,
		"third":  third,
	}

	ins, err := NewBuilder(encPkBitrate).Create().WithMap(mp).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
