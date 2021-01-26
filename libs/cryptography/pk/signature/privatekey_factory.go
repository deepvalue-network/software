package signature

type privateKeyFactory struct {
}

func createPrivateKeyFactory() PrivateKeyFactory {
	out := privateKeyFactory{}
	return &out
}

// Create creates a new PrivateKey instance
func (app *privateKeyFactory) Create() PrivateKey {
	x := curve.Scalar().Pick(curve.RandomStream())
	return createPrivateKey(x)
}
