package signature

import (
	kyber "go.dedis.ch/kyber/v3"
)

type publicKey struct {
	p kyber.Point
}

func createPublicKey(p kyber.Point) PublicKey {
	out := publicKey{
		p: p,
	}

	return &out
}

// Point returns the point
func (obj *publicKey) Point() kyber.Point {
	return obj.p
}

// Equals returns true if the given PublicKey equals the current one
func (obj *publicKey) Equals(pubKey PublicKey) bool {
	return obj.p.Equal(pubKey.Point())
}

// String returns the string representation of the public key
func (obj *publicKey) String() string {
	return obj.p.String()
}
