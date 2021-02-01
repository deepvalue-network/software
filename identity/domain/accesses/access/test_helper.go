package access

// CreateAccessForTests creates a new access instance for tests
func CreateAccessForTests(encPkBitrate int) Access {
	access, err := NewFactory(encPkBitrate).Create()
	if err != nil {
		panic(err)
	}

	return access
}
