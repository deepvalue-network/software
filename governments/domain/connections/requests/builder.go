package requests

import (
	"errors"
	"fmt"

	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
	"github.com/deepvalue-network/software/libs/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	content     Content
	sig         signature.Signature
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
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

// WithContent adds a content to the builder
func (app *builder) WithContent(content Content) Builder {
	app.content = content
	return app
}

// WithSignature adds a signature to the builder
func (app *builder) WithSignature(sig signature.Signature) Builder {
	app.sig = sig
	return app
}

// Now builds a new Request instance
func (app *builder) Now() (Request, error) {
	if app.content == nil {
		return nil, errors.New("the content is mandatory in order to build a Request instance")
	}

	if app.sig == nil {
		return nil, errors.New("the signature is mandatory in order to build a Request instance")
	}

	msg := app.content.Hash().String()
	pubKey := app.sig.PublicKey(msg)
	hashPubKey, err := app.hashAdapter.FromBytes([]byte(pubKey.String()))
	if err != nil {
		return nil, err
	}

	if hashPubKey.Compare(app.content.Requestor()) {
		str := fmt.Sprintf("the given signature (%s) cannot verify the content (hash: %s)", app.sig.String(), app.content.Hash().String())
		return nil, errors.New(str)
	}

	hash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.content.Hash().Bytes(),
		[]byte(app.sig.String()),
	})

	if err != nil {
		return nil, err
	}

	return createRequest(*hash, app.content, app.sig), nil
}
