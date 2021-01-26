package signature

import (
	"github.com/steve-care-software/products/pangolin/libs/hash"
	kyber "go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/group/edwards25519"
)

/*
 * H'(m, s, e) = H(m || s * G + e * P)
 * P = x * G
 * e = H(m || k * G)
 * k = s + e * x
 * s = k – e * x
 * k = H(m || x) -> to generate a new k, since nobody but us knows x
 * where ...
 * 1. H is a hash function, for instance SHA256.
 * 2. s and e are 2 numbers forming the ring signature
 * 3. s and r are a pubKey and a number forming a signature
 * 4. m is the message we want to sign
 * 5. P is the public key.
 * 6. G is the random base
 * 7. k is a number chosen randomly.  A new one every time we sign must be generated
 * 8. x is the private key
 */

const delimiter = "#"
const elementDelimiter = "|"

var curve = edwards25519.NewBlakeSHA256Ed25519()

// NewPrivateKeyFactory creates a PrivateKey factory
func NewPrivateKeyFactory() PrivateKeyFactory {
	return createPrivateKeyFactory()
}

// NewPrivateKeyAdapter creates a new PrivateKey adapter
func NewPrivateKeyAdapter() PrivateKeyAdapter {
	return createPrivateKeyAdapter()
}

// NewPublicKeyAdapter creates a new publicKey adapter
func NewPublicKeyAdapter() PublicKeyAdapter {
	return createPublicKeyAdapter()
}

// NewAdapter creates a signature adapter
func NewAdapter() Adapter {
	return createAdapter()
}

// NewRingSignatureAdapter creates a ring signature adapter
func NewRingSignatureAdapter() RingSignatureAdapter {
	hashAdapter := hash.NewAdapter()
	return createRingSignatureAdapter(hashAdapter)
}

// PrivateKeyFactory represents a privateKey factory
type PrivateKeyFactory interface {
	Create() PrivateKey
}

// PrivateKeyAdapter represents a privateKey adapter
type PrivateKeyAdapter interface {
	ToPrivateKey(pk string) (PrivateKey, error)
}

// PrivateKey represents a private key
type PrivateKey interface {
	PublicKey() PublicKey
	Sign(msg string) (Signature, error)
	RingSign(msg string, ringPubKeys []PublicKey) (RingSignature, error)
	String() string
}

// PublicKeyAdapter represents a publicKey adapter
type PublicKeyAdapter interface {
	ToPublicKey(pubKey string) (PublicKey, error)
}

// PublicKey represents the public key
type PublicKey interface {
	Point() kyber.Point
	Equals(pubKey PublicKey) bool
	String() string
}

// Adapter represents a signature adapter
type Adapter interface {
	ToSignature(sig string) (Signature, error)
}

// Signature represents a signature
type Signature interface {
	PublicKey(msg string) PublicKey
	Verify() bool
	String() string
}

// RingSignatureAdapter represents a ring signature adapter
type RingSignatureAdapter interface {
	ToSignature(sig string) (RingSignature, error)
	ToVerification(sig RingSignature, msg string, pubKeyHashes []hash.Hash) (bool, error)
}

// RingSignature represents a RingSignature
type RingSignature interface {
	Ring() []PublicKey
	Verify(msg string) bool
	String() string
}
