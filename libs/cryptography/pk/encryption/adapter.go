package encryption

import (
	"crypto/x509"
	"encoding/base64"
)

type adapter struct {
	builder Builder
}

func createAdapter(builder Builder) Adapter {
	out := adapter{
		builder: builder,
	}

	return &out
}

// FromBytes converts []byte to PrivateKey
func (app *adapter) FromBytes(bytes []byte) (PrivateKey, error) {
	pk, err := x509.ParsePKCS1PrivateKey(bytes)
	if err != nil {
		return nil, err
	}

	return app.builder.WithPK(*pk).Now()
}

// FromEncoded converts an encoded string to PrivateKey
func (app *adapter) FromEncoded(encoded string) (PrivateKey, error) {
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return nil, err
	}

	return app.FromBytes(decoded)
}

// ToBytes converts a PrivateKey to []byte
func (app *adapter) ToBytes(pk PrivateKey) []byte {
	key := pk.Key()
	return x509.MarshalPKCS1PrivateKey(&key)
}

// ToEncoded converts a PrivateKey to an encoded string
func (app *adapter) ToEncoded(pk PrivateKey) string {
	bytes := app.ToBytes(pk)
	return base64.StdEncoding.EncodeToString(bytes)
}
