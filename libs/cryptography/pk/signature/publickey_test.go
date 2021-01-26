package signature

import (
	"testing"
)

func TestPublicKey_Success(t *testing.T) {
	//variables:
	p := curve.Point().Base()

	// execute:
	pKey := createPublicKey(p)
	pubKeyStr := pKey.String()
	samePubKey, samePubKeyErr := NewPublicKeyAdapter().ToPublicKey(pubKeyStr)

	if samePubKeyErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", samePubKeyErr.Error())
		return
	}

	if !pKey.Equals(samePubKey) {
		t.Errorf("the public keys should be equal")
		return
	}
}
