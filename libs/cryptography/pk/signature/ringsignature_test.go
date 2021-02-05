package signature

import (
	"testing"

	"github.com/deepvalue-network/software/libs/hash"
)

func TestRingSignature_Success(t *testing.T) {
	ringSignatureAdapter := NewRingSignatureAdapter()
	hashAdapter := hash.NewAdapter()

	// variables:
	msg := "this is a message to sign"
	pk := NewPrivateKeyFactory().Create()
	secondPK := NewPrivateKeyFactory().Create()
	ringPubKeys := []PublicKey{
		pk.PublicKey(),
		secondPK.PublicKey(),
	}

	ringPubKeyHashes := []hash.Hash{}
	for _, onePubKey := range ringPubKeys {
		hsh, _ := hashAdapter.FromString(onePubKey.String())
		ringPubKeyHashes = append(ringPubKeyHashes, *hsh)
	}

	firstRing, err := pk.RingSign(msg, ringPubKeys)
	if err != nil {
		t.Errorf("the returned error was expected to be nil, error returned; %s", err.Error())
		return
	}

	secondRing, err := secondPK.RingSign(msg, ringPubKeys)
	if err != nil {
		t.Errorf("the returned error was expected to be nil, error returned; %s", err.Error())
		return
	}

	if !firstRing.Verify(msg) {
		t.Errorf("the first ring was expected to be verified")
		return
	}

	firstRingVerified, err := ringSignatureAdapter.ToVerification(firstRing, msg, ringPubKeyHashes)
	if err != nil {
		t.Errorf("the returned error was expected to be nil, error returned; %s", err.Error())
		return
	}

	if !firstRingVerified {
		t.Errorf("the first ring was expected to be deep verified")
		return
	}

	if !secondRing.Verify(msg) {
		t.Errorf("the second ring was expected to be verified")
		return
	}

	secondRingVerified, err := ringSignatureAdapter.ToVerification(secondRing, msg, ringPubKeyHashes)
	if err != nil {
		t.Errorf("the returned error was expected to be nil, error returned; %s", err.Error())
		return
	}

	if !secondRingVerified {
		t.Errorf("the second ring was expected to be deep verified")
		return
	}

	// encode to string, back and forth:
	firstRingStr := firstRing.String()
	newRing, err := NewRingSignatureAdapter().ToSignature(firstRingStr)
	if err != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if firstRingStr != newRing.String() {
		t.Errorf("the rings were expected to be the same.  Expected: %s, Actual: %s", firstRingStr, newRing.String())
		return
	}
}

func TestRingSignature_PubKeyIsNotInTheRing_returnsError(t *testing.T) {
	// variables:
	msg := "this is a message to sign"
	pk := NewPrivateKeyFactory().Create()
	secondPK := NewPrivateKeyFactory().Create()
	invalidPK := NewPrivateKeyFactory().Create()
	ringPubKeys := []PublicKey{
		pk.PublicKey(),
		secondPK.PublicKey(),
	}

	_, err := invalidPK.RingSign(msg, ringPubKeys)
	if err == nil {
		t.Errorf("the returned error was expected to be valid, nil returned")
		return
	}
}
