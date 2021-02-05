package bills

import (
	"errors"
	"fmt"

	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
	"github.com/deepvalue-network/software/libs/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	content     Content
	sig         signature.RingSignature
}

func createBuilder(hashAdapter hash.Adapter) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		content:     nil,
		sig:         nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithContent adds content to the builder
func (app *builder) WithContent(content Content) Builder {
	app.content = content
	return app
}

// WithSignature adds a signature to the builder
func (app *builder) WithSignature(signature signature.RingSignature) Builder {
	app.sig = signature
	return app
}

// Now builds a new Bill instance
func (app *builder) Now() (Bill, error) {
	if app.content == nil {
		return nil, errors.New("the content is mandatory in order to build a Bill instance")
	}

	if app.sig == nil {
		return nil, errors.New("the signature is mandatory in order to build a Bill instance")
	}

	msg := app.content.Hash().String()
	if !app.sig.Verify(msg) {
		str := fmt.Sprintf("the signature could not be verified against the content hash: %s", msg)
		return nil, errors.New(str)
	}

	// make sure all the public keys are the same:
	pubKeys := app.sig.Ring()
	pubKeyHashes := app.content.Owner()
	if len(pubKeys) != len(pubKeyHashes) {
		str := fmt.Sprintf("the ring signature contains %d public keys, but %d hashed public keys were provided", len(pubKeys), len(pubKeyHashes))
		return nil, errors.New(str)
	}

	for _, onePubKey := range pubKeys {
		pubKeyHash, err := app.hashAdapter.FromBytes([]byte(onePubKey.String()))
		if err != nil {
			return nil, err
		}

		exists := false
		for _, oneHash := range pubKeyHashes {
			if pubKeyHash.Compare(oneHash) {
				exists = true
				break
			}
		}

		if !exists {
			str := fmt.Sprintf("the hashed public key (%s) was expected in the hashed pub key owners because it was provided in the ring signature", pubKeyHash.String())
			return nil, errors.New(str)
		}
	}

	hsh, err := app.hashAdapter.FromMultiBytes([][]byte{})
	if err != nil {
		return nil, err
	}

	return createBill(*hsh, app.content, app.sig), nil
}
