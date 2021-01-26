package signature

import (
	"bytes"
	"encoding/hex"

	kyber "go.dedis.ch/kyber/v3"
)

func createHash(msg string) kyber.Scalar {
	sha256 := curve.Hash()
	sha256.Reset()
	sha256.Write([]byte(msg))

	return curve.Scalar().SetBytes(sha256.Sum(nil))
}

func genK(x kyber.Scalar, msg string) kyber.Scalar {
	return createHash(msg + x.String())
}

func fromStringToScalar(str string) (kyber.Scalar, error) {
	decoded, decodedErr := hex.DecodeString(str)
	if decodedErr != nil {
		return nil, decodedErr
	}

	x := curve.Scalar()
	reader := bytes.NewReader(decoded)
	_, err := x.UnmarshalFrom(reader)
	if err != nil {
		return nil, err
	}

	return x, nil
}

func fromStringToPoint(str string) (kyber.Point, error) {
	decoded, decodedErr := hex.DecodeString(str)
	if decodedErr != nil {
		return nil, decodedErr
	}

	p := curve.Point()
	err := p.UnmarshalBinary(decoded)
	if err != nil {
		return nil, err
	}

	return p, nil
}
