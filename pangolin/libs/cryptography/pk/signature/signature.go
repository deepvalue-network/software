package signature

import (
	"encoding/base64"
	"errors"
	"fmt"
	"math/rand"
	"strconv"

	kyber "go.dedis.ch/kyber/v3"
)

type signature struct {
	r PublicKey
	s kyber.Scalar
}

func createSignature(r PublicKey, s kyber.Scalar) (Signature, error) {
	out := signature{r: r, s: s}
	if !out.Verify() {
		return nil, errors.New("the signature could not verify")
	}

	return &out, nil
}

// PublicKey returns the public key of the signature
func (app *signature) PublicKey(msg string) PublicKey {
	// Create a generator.
	g := curve.Point().Base()

	// e = Hash(m || r)
	e := createHash(msg + app.r.String())

	// y = (r - s * G) * (1 / e)
	y := curve.Point().Sub(app.r.Point(), curve.Point().Mul(app.s, g))
	y = curve.Point().Mul(curve.Scalar().Div(curve.Scalar().One(), e), y)

	return createPublicKey(y)
}

// Verify verifies if the signature has been made by the given public key
func (app *signature) Verify() bool {

	// generate a message:
	msg := strconv.Itoa(rand.Int())

	// retrieve pubKey:
	p := app.PublicKey(msg)

	// Create a generator.
	g := curve.Point().Base()

	// e = Hash(m || r)
	e := createHash(msg + app.r.String())

	// Attempt to reconstruct 's * G' with a provided signature; s * G = r - e * p
	sGv := curve.Point().Sub(app.r.Point(), curve.Point().Mul(e, p.Point()))

	// Construct the actual 's * G'
	sG := curve.Point().Mul(app.s, g)

	// Equality check; ensure signature and public key outputs to s * G.
	return sG.Equal(sGv)
}

// String returns the string representation of the signature
func (app *signature) String() string {
	str := fmt.Sprintf("%s%s%s", app.r.String(), delimiter, app.s.String())
	return base64.StdEncoding.EncodeToString([]byte(str))
}
