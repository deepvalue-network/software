package signature

import (
	"errors"
	"fmt"

	kyber "go.dedis.ch/kyber/v3"
)

type privateKey struct {
	x kyber.Scalar
}

func createPrivateKey(x kyber.Scalar) PrivateKey {
	out := privateKey{
		x: x,
	}

	return &out
}

// PublicKey returns the public key
func (app *privateKey) PublicKey() PublicKey {
	g := curve.Point().Base()
	p := curve.Point().Mul(app.x, g)
	return createPublicKey(p)
}

// RingSign signs a ring signature on the given message, in the given ring pubKey
func (app *privateKey) RingSign(msg string, ringPubKeys []PublicKey) (RingSignature, error) {

	retrieveSignerIndexFn := func(ringPubKeys []PublicKey, pk PrivateKey) int {
		pubKey := pk.PublicKey()
		for index, oneRingPubKey := range ringPubKeys {
			if oneRingPubKey.Equals(pubKey) {
				return index
			}
		}

		return -1
	}

	// retrieve our signerIndex:
	signerIndex := retrieveSignerIndexFn(ringPubKeys, app)
	if signerIndex == -1 {
		return nil, errors.New("the signer PublicKey is not in the ring")
	}

	// generate k:
	k := genK(app.x, msg)

	// random base:
	g := curve.Point().Base()

	// length:
	r := len(ringPubKeys)

	// initialize:
	es := make([]kyber.Scalar, r)
	ss := make([]kyber.Scalar, r)
	beginIndex := (signerIndex + 1) % r

	// ei = H(m || k * G)
	es[beginIndex] = createHash(msg + curve.Point().Mul(k, g).String())

	// loop:
	for i := beginIndex; i != signerIndex; i = (i + 1) % r {
		// si = random value
		ss[i] = genK(app.x, fmt.Sprintf("%s%d", msg, i))

		//eiPlus1ModR = H(m || si * G + ei * Pi)
		sig := curve.Point().Mul(ss[i], g)
		eipi := curve.Point().Mul(es[i], ringPubKeys[i].Point())
		es[(i+1)%r] = createHash(msg + curve.Point().Add(sig, eipi).String())

	}

	// close the ring:
	esx := curve.Scalar().Mul(es[signerIndex], app.x)
	ss[signerIndex] = curve.Scalar().Sub(k, esx)
	out := createRingSignature(ringPubKeys, ss, es[0])
	return out, nil
}

// Sign signs a message
func (app *privateKey) Sign(msg string) (Signature, error) {
	// generate k:
	k := genK(app.x, msg)

	// random base:
	g := curve.Point().Base()

	// r = k * G (a.k.a the same operation as r = g^k)
	r := curve.Point().Mul(k, g)

	// hash(m || r)
	e := createHash(msg + r.String())

	// s = k - e * x
	s := curve.Scalar().Sub(k, curve.Scalar().Mul(e, app.x))

	// create signature:
	pubKey := createPublicKey(r)
	return createSignature(pubKey, s)
}

// String returns the string representation of the PrivateKey
func (app *privateKey) String() string {
	return app.x.String()
}
