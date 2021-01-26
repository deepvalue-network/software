package signature

import (
	"encoding/base64"
	"fmt"

	kyber "go.dedis.ch/kyber/v3"
)

type ringSignature struct {
	ring []PublicKey
	s    []kyber.Scalar
	e    kyber.Scalar
}

func createRingSignature(ring []PublicKey, s []kyber.Scalar, e kyber.Scalar) RingSignature {
	out := ringSignature{
		ring: ring,
		s:    s,
		e:    e,
	}

	return &out
}

// Ring returns ring pubKeys
func (app *ringSignature) Ring() []PublicKey {
	return app.ring
}

// Verify verifies if the message has been signed by at least 1 shared signature
func (app *ringSignature) Verify(msg string) bool {
	// random base:
	g := curve.Point().Base()

	// first e:
	e := app.e

	//e = H(m || s[i] * G + e * P[i]);
	amount := len(app.ring)
	for i := 0; i < amount; i++ {
		sg := curve.Point().Mul(app.s[i], g)
		ep := curve.Point().Mul(e, app.ring[i].Point())
		added := curve.Point().Add(sg, ep)
		e = createHash(msg + added.String())
	}

	return app.e.Equal(e)
}

// String returns the string representation of the ring signature
func (app *ringSignature) String() string {
	ringStr := ""
	for _, onePubKey := range app.ring {
		ringStr = fmt.Sprintf("%s%s%s", ringStr, onePubKey.String(), elementDelimiter)
	}

	sScalarStr := ""
	for _, oneScalar := range app.s {
		sScalarStr = fmt.Sprintf("%s%s%s", sScalarStr, oneScalar.String(), elementDelimiter)
	}

	str := fmt.Sprintf("%s%s%s%s%s", ringStr, delimiter, sScalarStr, delimiter, app.e.String())
	return base64.StdEncoding.EncodeToString([]byte(str))
}
