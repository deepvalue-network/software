package signature

type publicKeyAdapter struct {
}

func createPublicKeyAdapter() PublicKeyAdapter {
	out := publicKeyAdapter{}
	return &out
}

// ToPublicKey converts a string to a publicKey
func (app *publicKeyAdapter) ToPublicKey(pubKey string) (PublicKey, error) {
	point, err := fromStringToPoint(pubKey)
	if err != nil {
		return nil, err
	}

	return createPublicKey(point), err
}
