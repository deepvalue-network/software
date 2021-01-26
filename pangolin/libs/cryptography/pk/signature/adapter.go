package signature

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
)

type adapter struct {
}

func createAdapter() Adapter {
	out := adapter{}
	return &out
}

// ToSignature converts a string to a Signature instance
func (app *adapter) ToSignature(sig string) (Signature, error) {
	decoded, err := base64.StdEncoding.DecodeString(sig)
	if err != nil {
		return nil, err
	}

	splitted := strings.Split(string(decoded), delimiter)
	if len(splitted) != 2 {
		str := fmt.Sprintf("the signature string was expected to have %d sections, %d found", 2, len(splitted))
		return nil, errors.New(str)
	}

	point, err := fromStringToPoint(splitted[0])
	if err != nil {
		return nil, err
	}

	scalar, err := fromStringToScalar(splitted[1])
	if err != nil {
		return nil, err
	}

	pubKey := createPublicKey(point)
	return createSignature(pubKey, scalar)
}
