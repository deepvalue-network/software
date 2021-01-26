package signature

import (
	"testing"
)

func TestPrivateKey_Success(t *testing.T) {
	// variables:
	pk := NewPrivateKeyFactory().Create()
	pkStr := pk.String()
	bckPK, err := NewPrivateKeyAdapter().ToPrivateKey(pkStr)
	if err != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if pk.String() != bckPK.String() {
		t.Errorf("the PrivateKeys were expected to be the same.  Expected: %s, Returned: %s", pk.String(), bckPK.String())
		return
	}
}
