package signature

type privateKeyAdapter struct {
}

func createPrivateKeyAdapter() PrivateKeyAdapter {
	out := privateKeyAdapter{}
	return &out
}

// ToPrivateKey converts a string to a privateKey instance
func (app *privateKeyAdapter) ToPrivateKey(pk string) (PrivateKey, error) {
	scalar, err := fromStringToScalar(pk)
	if err != nil {
		return nil, err
	}

	return createPrivateKey(scalar), nil
}
