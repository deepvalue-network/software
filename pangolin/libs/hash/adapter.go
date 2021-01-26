package hash

import (
	"crypto/sha512"
	"encoding/hex"
)

type adapter struct {
}

func createAdapter() Adapter {
	out := adapter{}
	return &out
}

// FromBytes converts []byte to an Hash
func (app *adapter) FromBytes(input []byte) (*Hash, error) {
	if len(input) != 64 {
		h := sha512.New()
		_, err := h.Write(input)
		if err != nil {
			return nil, err
		}

		input = h.Sum(nil)
	}

	out := Hash{}
	for index, oneByte := range input {
		out[index] = oneByte
	}

	return &out, nil
}

// FromMultiBytes converts multiple []byte to an Hash
func (app *adapter) FromMultiBytes(input [][]byte) (*Hash, error) {
	merged := []byte{}
	for _, oneRow := range input {
		merged = append(merged, oneRow...)
	}

	return app.FromBytes(merged)
}

// FromString converts a string to an Hash
func (app *adapter) FromString(input string) (*Hash, error) {
	bytes, err := hex.DecodeString(input)
	if err != nil {
		return nil, err
	}

	hash, err := app.FromBytes(bytes)
	if err != nil {
		return nil, err
	}

	return hash, nil
}
