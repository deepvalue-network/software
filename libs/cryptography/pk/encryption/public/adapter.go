package public

import (
	"crypto/x509"
	"encoding/base64"

	"github.com/steve-care-software/products/libs/hash"
)

type adapter struct {
	hashAdapter hash.Adapter
	builder     Builder
}

func createAdapter(hashAdapter hash.Adapter, builder Builder) Adapter {
	out := adapter{
		hashAdapter: hashAdapter,
		builder:     builder,
	}

	return &out
}

// FromBytes converts []byte to Key
func (app *adapter) FromBytes(input []byte) (Key, error) {
	pubKey, err := x509.ParsePKCS1PublicKey(input)
	if err != nil {
		return nil, err
	}

	return app.builder.Create().WithKey(*pubKey).Now()
}

// FromEncoded converts an encoded string to Key
func (app *adapter) FromEncoded(encoded string) (Key, error) {
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return nil, err
	}

	return app.FromBytes(decoded)
}

// ToBytes converts Key to []byte
func (app *adapter) ToBytes(key Key) []byte {
	pubKey := key.Key()
	return x509.MarshalPKCS1PublicKey(&pubKey)
}

// ToEncoded converts a Key to an encoded string
func (app *adapter) ToEncoded(key Key) string {
	bytes := app.ToBytes(key)
	return base64.StdEncoding.EncodeToString(bytes)
}

// ToHash converts a Key to an hash
func (app *adapter) ToHash(key Key) (*hash.Hash, error) {
	bytes := app.ToBytes(key)
	return app.hashAdapter.FromBytes(bytes)
}
